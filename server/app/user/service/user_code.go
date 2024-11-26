package service

import (
	"github.com/PokemanMaster/GoChat/server/server/app/user/serializer"
	"github.com/PokemanMaster/GoChat/server/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/server/pkg/utils"
	"github.com/PokemanMaster/GoChat/server/server/resp"
	"go.uber.org/zap"
)

// UserCategoryService 前端请求过来的数据
type UserCategoryService struct{}

// UserCategoryImages 给用户返回base64码的图片
func (service *UserCategoryService) UserCategoryImages() resp.Response {
	codeId, base64, err := utils.CreateCode()
	code := e.SUCCESS
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e.ERROR
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildBase64(codeId, base64),
	}
}
