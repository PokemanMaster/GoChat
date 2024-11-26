package api

import (
	service2 "github.com/PokemanMaster/GoChat/server/app/cart/service"
	"github.com/PokemanMaster/GoChat/server/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateCart 添加商品到购物车接口
func CreateCart(ctx *gin.Context) {
	services := service2.CreateCartService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.cart.api", "cart.go"))
	} else {
		res := services.Create(ctx)
		ctx.JSON(200, res)
	}
}

// ShowCart 展示购物车接口
func ShowCart(ctx *gin.Context) {
	services := service2.ShowCartService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.cart.api", "cart.go"))
	} else {
		res := services.Show(ctx, ctx.Param("id"))
		ctx.JSON(200, res)
	}
}

// UpdateCart 修改购物车信息
func UpdateCart(ctx *gin.Context) {
	services := service2.UpdateCartService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.cart.api", "cart.go"))
	} else {
		res := services.Update(ctx)
		ctx.JSON(200, res)
	}
}

// DeleteCart 移除购物车接口
func DeleteCart(ctx *gin.Context) {
	services := service2.DeleteCartService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.cart.api", "cart.go"))
	} else {
		res := services.Delete(ctx)
		ctx.JSON(200, res)
	}
}
