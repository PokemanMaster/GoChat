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
		zap.L().Error("请求参数错误", zap.String("app.favorite.api.favorite.go", ""))
		ctx.JSON(400, resp.ErrorResponse(err))
	} else {
		res := services.Create(ctx)
		ctx.JSON(200, res)
	}
}

// ShowFavorites 展示收藏夹接口
func ShowFavorites(ctx *gin.Context) {
	services := service.ShowFavoritesService{}
	if err := ctx.ShouldBind(&services); err != nil {
		zap.L().Error("请求参数错误", zap.String("app.favorite.api.favorite.go", ""))
		ctx.JSON(00, resp.ErrorResponse(err))
	} else {
		res := services.Show(ctx, ctx.Param("id"))
		ctx.JSON(200, res)

	}
}

// DeleteFavorite 删除收藏夹的接口
func DeleteFavorite(ctx *gin.Context) {
	services := service.DeleteFavoriteService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		zap.L().Error("请求参数错误", zap.String("app.favorite.api.favorite.go", ""))
		ctx.JSON(400, resp.ErrorResponse(err))
	} else {
		res := services.Delete(ctx)
		ctx.JSON(200, res)
	}
}
