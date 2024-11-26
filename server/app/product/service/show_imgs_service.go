package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/product/model"
	"github.com/PokemanMaster/GoChat/v1/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"
)

// ShowProductParamService 商品图片详情的服务
type ShowProductParamService struct {
}

// Show 商品参数
func (service *ShowProductParamService) Show(id string) resp.Response {
	var productsParam []model.ProductParam
	code := e.SUCCESS

	err := db.DB.Where("product_id=?", id).Find(&productsParam).Error
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
		Data: serializer.BuildProductParams(productsParam),
	}
}
