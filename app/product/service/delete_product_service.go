package service

import (
	"github.com/PokemanMaster/GoChat/app/product/model"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"github.com/PokemanMaster/GoChat/resp"
)

// DeleteProductService 删除商品的服务
type DeleteProductService struct {
}

// Delete 删除商品
func (service *DeleteProductService) Delete(id string) resp.Response {
	var product model.Product
	code := e.SUCCESS

	err := db.DB.First(&product, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = db.DB.Delete(&product).Error
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
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
