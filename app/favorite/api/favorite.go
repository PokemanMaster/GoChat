package api

import (
	"github.com/PokemanMaster/GoChat/app/favorite/service"
	"github.com/PokemanMaster/GoChat/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateFavorite 创建收藏接口
func CreateFavorite(ctx *gin.Context) {
	services := service.CreateFavoriteService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.favorite.api", "favorite.go"))
	} else {
		res := services.Create(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// ShowFavorites 展示收藏夹接口
func ShowFavorites(ctx *gin.Context) {
	services := service.ShowFavoritesService{}
	if err := ctx.ShouldBind(&services); err == nil {
		res := services.Show(ctx, ctx.Param("id"))
		ctx.JSON(200, res)
	} else {
		ctx.JSON(200, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.favorite.api", "favorite.go"))
	}
}

// DeleteFavorite 删除收藏夹的接口
func DeleteFavorite(ctx *gin.Context) {
	services := service.DeleteFavoriteService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, resp.ErrorResponse(err))
		zap.L().Error("请求参数错误", zap.String("app.favorite.api", "favorite.go"))
	} else {
		res := services.Delete(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}
