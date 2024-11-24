package service

import (
	"github.com/PokemanMaster/GoChat/app/product/model"
	"github.com/PokemanMaster/GoChat/app/product/serializer"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"github.com/PokemanMaster/GoChat/resp"
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
