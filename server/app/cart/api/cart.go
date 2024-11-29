package api

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/cart/service"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateCart 添加商品到购物车接口
func CreateCart(ctx *gin.Context) {
	services := service.CreateCartService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.cart.api.cart", err.Error()))
	} else {
		res := services.Create(ctx)
		ctx.JSON(200, res)
	}
}

// ShowCart 展示购物车接口
func ShowCart(ctx *gin.Context) {
	services := service.ShowCartService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.cart.api.cart", err.Error()))
	} else {
		res := services.Show(ctx, ctx.Param("id"))
		ctx.JSON(200, res)
	}
}

// UpdateCart 修改购物车信息
func UpdateCart(ctx *gin.Context) {
	services := service.UpdateCartService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.cart.api.cart", err.Error()))
	} else {
		res := services.Update(ctx)
		ctx.JSON(200, res)
	}
}

// DeleteCart 移除购物车接口
func DeleteCart(ctx *gin.Context) {
	services := service.DeleteCartService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.cart.api.cart", err.Error()))
	} else {
		res := services.Delete(ctx)
		ctx.JSON(200, res)
	}
}
