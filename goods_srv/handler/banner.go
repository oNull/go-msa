package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"goods_srv/global"
	"goods_srv/model"
	"goods_srv/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BannerList BannerList(context.Context, *empty.Empty) (*BannerListResponse, error)
func (g *GoodsServer) BannerList(ctx context.Context, req *empty.Empty) (*proto.BannerListResponse, error) {
	bannerListResponse := &proto.BannerListResponse{}
	var banners []model.Banner
	result := global.Db.Where("is_deleted != 1").Find(&banners)
	if result.Error != nil {
		return nil, result.Error
	}

	bannerListResponse.Total = int32(result.RowsAffected)
	var bannerResponse []*proto.BannerResponse
	for _, banner := range banners {
		bannerResponse = append(bannerResponse, &proto.BannerResponse{
			Id:    banner.ID,
			Image: banner.Image,
			Url:   banner.Url,
			Index: banner.Index,
		})
	}

	bannerListResponse.Data = bannerResponse
	return bannerListResponse, nil
}

// CreateBanner CreateBanner(context.Context, *BannerRequest) (*BannerResponse, error)
func (g *GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {
	bannerResponse := &proto.BannerResponse{}
	banner := model.Banner{
		Image: req.Image,
		Url:   req.Url,
		Index: req.Index,
	}
	global.Db.Save(&banner)
	bannerResponse.Id = banner.ID
	return bannerResponse, nil
}

// DeleteBanner DeleteBanner(context.Context, *BannerRequest) (*empty.Empty, error)
func (g *GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerRequest) (*empty.Empty, error) {
	if result := global.Db.Delete(&model.Banner{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "未找到该Banner")
	}
	return &empty.Empty{}, nil
}

// UpdateBanner UpdateBanner(context.Context, *BannerRequest) (*empty.Empty, error)
func (g *GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerRequest) (*empty.Empty, error) {
	var banner model.Banner
	if result := global.Db.First(&banner, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "未找到该Banner")
	}
	if req.Url != "" {
		banner.Url = req.Url
	}
	if req.Image != "" {
		banner.Image = req.Image
	}
	if req.Index != 0 {
		banner.Index = req.Index
	}
	result := global.Db.Save(&banner)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "更新失败")
	}
	return &empty.Empty{}, nil
}
