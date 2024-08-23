package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"goods_srv/global"
	"goods_srv/model"
	"goods_srv/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (g *GoodsServer) GetAllCategorysList(ctx context.Context, req *emptypb.Empty) (*proto.CategoryListResponse, error) {
	//categoryListResponse := &proto.CategoryListResponse{}
	//var categories []model.Category
	//
	//result := global.Db.Where("parent_category_id = ?", 0).Find(&categories)
	//if result.Error != nil {
	//	return nil, result.Error
	//}
	//
	//categoryListResponse.Total = int32(result.RowsAffected)
	//var categoryData []*proto.CategoryInfoResponse
	//for _, category := range categories {
	//	categoryData = append(categoryData, &proto.CategoryInfoResponse{
	//		Id:             category.ID,
	//		Name:           category.Name,
	//		Level:          category.Level,
	//	})
	//}
	//categoryListResponse.Data = categoryData
	//return categoryListResponse, nil
	var categorys []model.Category
	// SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	// 配置指明了外键后，可以使用Preload预加载，来把品牌的子分类也取出来
	global.Db.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	b, _ := json.Marshal(&categorys)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

func (g *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	categoryListResponse := &proto.SubCategoryListResponse{}
	var category model.Category
	if result := global.Db.First(&category, req.Id); result.Error != nil {
		return nil, result.Error
	}

	categoryListResponse.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	var subCategoryResponse []*proto.CategoryInfoResponse
	var subCategorys []model.Category

	global.Db.Where(&model.Category{ParentCategoryID: req.Id}).Find(&subCategorys)
	for _, subCategory := range subCategorys {
		subCategoryResponse = append(subCategoryResponse, &proto.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			ParentCategory: subCategory.ParentCategoryID,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
		})
	}
	categoryListResponse.SubCategorys = subCategoryResponse
	return categoryListResponse, nil
}

func (g *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := model.Category{}
	cMap := map[string]interface{}{}
	cMap["name"] = req.Name
	cMap["level"] = req.Level
	cMap["is_tab"] = req.IsTab
	if req.Level != 1 {
		//去查询父类目是否存在
		cMap["parent_category_id"] = req.ParentCategory
	}
	tx := global.Db.Model(&model.Category{}).Create(cMap)
	fmt.Println(tx)
	return &proto.CategoryInfoResponse{Id: int32(category.ID)}, nil
}

func (g *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := global.Db.Delete(&model.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	return &emptypb.Empty{}, nil
}

func (g *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	var category model.Category
	if result := global.Db.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}
	global.Db.Save(&category)
	return &emptypb.Empty{}, nil
}
