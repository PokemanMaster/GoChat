package service

import (
	"github.com/PokemanMaster/GoChat/app/product/model"
	"github.com/PokemanMaster/GoChat/app/product/serializer"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
	"go.uber.org/zap"
)

// CreateCarouselService 轮播图创建的服务
type CreateCarouselService struct {
	ImgPath string `form:"img_path" json:"img_path"`
}

// Create 创建商品
func (service *CreateCarouselService) Create() resp.Response {
	carousel := model.Carousel{
		ImgPath: service.ImgPath,
	}
	code := e.SUCCESS

	err := db.DB.Create(&carousel).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e.ERROR_DATABASE
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarousel(carousel),
	}
}
