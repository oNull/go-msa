package handler

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"goods_srv/global"
	"goods_srv/model"
	"goods_srv/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// https://github.com/grpc/grpc-go/blob/master/cmd/protoc-gen-go-grpc/README.md
type GoodsServer struct {
	*proto.UnimplementedGoodsServer //解决"missing mustEmbedUnimplementedBlogServiceServer method"的问题
}

func (g *GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {

	goodsListResponse := &proto.GoodsListResponse{}
	localDb := global.Db.Model(model.Goods{})
	if req.KeyWords != "" {
		//模糊查询
		localDb = localDb.Where("name like ?", "%"+req.KeyWords+"%")
	}

	if req.PriceMin != 0 {
		localDb = localDb.Where("shop_price >= ?", req.PriceMin)
	}

	if req.PriceMax != 0 {
		localDb = localDb.Where("shop_price <= ?", req.PriceMax)
	}

	var subQuery string
	if req.TopCategory > 0 {
		var category model.Category
		if result := global.Db.First(&category, req.TopCategory); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}
		if category.Level == 1 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id in (select id from category WHERE parent_category_id=%d)", req.TopCategory)
		} else if category.Level == 2 {
			subQuery = fmt.Sprintf("select id from category WHERE parent_category_id=%d", req.TopCategory)
		} else if category.Level == 3 {
			subQuery = fmt.Sprintf("select id from category WHERE id=%d", req.TopCategory)
		}

		localDb = localDb.Where(fmt.Sprintf("category_id in (%s)", subQuery))
	}

	var total int64
	localDb.Count(&total)
	goodsListResponse.Total = int32(total)

	var goods []model.Goods
	result := localDb.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&goods)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}

	return goodsListResponse, nil
}

// BatchGetGoods 现在用户提交订单有多个商品，你得批量查询商品的信息吧
func (g *GoodsServer) BatchGetGoods(ctx context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	goodsListResponse := &proto.GoodsListResponse{}
	var goods []model.Goods

	//调用where并不会真正执行sql 只是用来生成sql的 当调用find， first才会去执行sql，
	result := global.Db.Where(req.Id).Find(&goods)
	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}
	goodsListResponse.Total = int32(result.RowsAffected)
	return goodsListResponse, nil
}

func (g *GoodsServer) GetGoodsDetail(ctx context.Context, req *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
	var goods model.Goods

	if result := global.Db.Preload("Category").Preload("Brands").First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	goodsInfoResponse := ModelToResponse(goods)
	return &goodsInfoResponse, nil
}

func (g *GoodsServer) CreateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	// 查询分类
	var category model.Category
	if result := global.Db.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	// 查询品牌
	var brand model.Brands
	if result := global.Db.First(&brand, req.Brand); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品品牌不存在")
	}

	//先检查redis中是否有这个token
	//防止同一个token的数据重复插入到数据库中，如果redis中没有这个token则放入redis
	//这里没有看到图片文件是如何上传， 在微服务中 普通的文件上传已经不再使用
	goods := model.Goods{
		Brands:          brand,
		BrandsID:        brand.ID,
		Category:        category,
		CategoryID:      category.ID,
		Name:            req.Name,
		GoodsSn:         req.GoodsSn,
		MarketPrice:     req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodsBrief,
		ShipFree:        req.ShipFree,
		Images:          req.Images,
		DescImages:      req.DescImages,
		GoodsFrontImage: req.GoodsFrontImage,
		IsNew:           req.IsNew,
		IsHot:           req.IsHot,
		OnSale:          req.OnSale,
	}

	// 增加事务处理
	tx := global.Db.Begin()
	result := tx.Save(&goods) // Save的时候会自动调用钩子函数，所以在这里添加事务处理
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return &proto.GoodsInfoResponse{
		Id: goods.ID,
	}, nil
}

func (g *GoodsServer) DeleteGoods(ctx context.Context, req *proto.DeleteGoodsInfo) (*empty.Empty, error) {
	var goods model.Goods
	if result := global.Db.Delete(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	return &empty.Empty{}, nil
}

func (g *GoodsServer) UpdateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*empty.Empty, error) {
	var goods model.Goods

	if result := global.Db.First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	var category model.Category
	if result := global.Db.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.Db.First(&brand, req.Brand); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	goods.Brands = brand
	goods.BrandsID = brand.ID
	goods.Category = category
	goods.CategoryID = category.ID
	goods.Name = req.Name
	goods.GoodsSn = req.GoodsSn
	goods.MarketPrice = req.MarketPrice
	goods.ShopPrice = req.ShopPrice
	goods.GoodsBrief = req.GoodsBrief
	goods.ShipFree = req.ShipFree
	goods.Images = req.Images
	goods.DescImages = req.DescImages
	goods.GoodsFrontImage = req.GoodsFrontImage
	goods.IsNew = req.IsNew
	goods.IsHot = req.IsHot
	goods.OnSale = req.OnSale

	tx := global.Db.Begin()
	result := tx.Save(&goods)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return &empty.Empty{}, nil
}
