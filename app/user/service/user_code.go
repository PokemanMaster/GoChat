package service

import (
	"github.com/PokemanMaster/GoChat/server/app/user/serializer"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/pkg/utils"
	"github.com/PokemanMaster/GoChat/server/resp"
	"go.uber.org/zap"
)

// UserCategoryService 前端请求过来的数据
type UserCategoryService struct{}

// UserCategoryImages 给用户返回base64码的图片
func (service *UserCategoryService) UserCategoryImages() resp.Response {
	codeId, base64, err := utils.CreateCode()
	code := e2.SUCCESS
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e2.ERROR
		return resp.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return resp.Response{
		Status: code,
		Msg:    e2.GetMsg(code),
		Data:   serializer.BuildBase64(codeId, base64),
	}
}
