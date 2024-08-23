package handler

import (
	"context"
	"goods_srv/global"
	"goods_srv/model"
	"goods_srv/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (g *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	brandListResponse := proto.BrandListResponse{}

	var brands []model.Brands
	result := global.Db.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	var count int64
	global.Db.Model(&model.Brands{}).Count(&count)
	brandListResponse.Total = int32(count)
	var brandResponses []*proto.BrandInfoResponse
	for _, brand := range brands {
		brandResponses = append(brandResponses, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}

	brandListResponse.Data = brandResponses
	return &brandListResponse, nil
}

func (g *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	//新建品牌 需要先查询品牌是否已存在
	result := global.Db.Where("name=?", req.Name).First(&model.Brands{})

	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌已存在")
	}

	brand := &model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	global.Db.Save(brand)
	return &proto.BrandInfoResponse{Id: brand.ID}, nil
}

func (g *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	result := global.Db.Delete(&model.Brands{}, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	return &emptypb.Empty{}, nil
}

func (g *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	brand := &model.Brands{}
	result := global.Db.First(brand, req.Id)

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	if req.Name != "" {
		brand.Name = req.Name
	}

	if req.Logo != "" {
		brand.Logo = req.Logo
	}

	global.Db.Save(brand)
	return &emptypb.Empty{}, nil
}

//DeleteBrand(context.Context, *BrandRequest) (*empty.Empty, error)
//UpdateBrand(context.Context, *BrandRequest) (*empty.Empty, error)
