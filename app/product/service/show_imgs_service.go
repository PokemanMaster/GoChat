package service

import (
	"IMProject/app/product/model"
	"IMProject/app/product/serializer"
	"IMProject/common/db"
	"IMProject/pkg/e"
	"IMProject/pkg/logging"
	"IMProject/resp"
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
		logging.Info(err)
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
