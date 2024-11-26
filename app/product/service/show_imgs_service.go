package service

import (
	"github.com/PokemanMaster/GoChat/server/app/product/model"
	"github.com/PokemanMaster/GoChat/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"
	"go.uber.org/zap"
)

// ShowProductParamService 商品图片详情的服务
type ShowProductParamService struct {
}

// Show 商品参数
func (service *ShowProductParamService) Show(id string) resp.Response {
	var productsParam []model.ProductParam
	code := e2.SUCCESS

	err := db.DB.Where("product_id=?", id).Find(&productsParam).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e2.ERROR_DATABASE
		return resp.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return resp.Response{
		Data: serializer.BuildProductParams(productsParam),
	}
}
