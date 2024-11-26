package router

import (
	"github.com/PokemanMaster/GoChat/server/server/app/carousel/api"
	CartApi "github.com/PokemanMaster/GoChat/server/server/app/cart/api"
	api2 "github.com/PokemanMaster/GoChat/server/server/app/category/api"
	ChatApi "github.com/PokemanMaster/GoChat/server/server/app/chat/api"
	FavoriteApi "github.com/PokemanMaster/GoChat/server/server/app/favorite/api"
	FriendApi "github.com/PokemanMaster/GoChat/server/server/app/friend/api"
	GroupApi "github.com/PokemanMaster/GoChat/server/server/app/group/api"
	OrderApi "github.com/PokemanMaster/GoChat/server/server/app/order/api"
	PaymentApi "github.com/PokemanMaster/GoChat/server/server/app/payment/api"
	api4 "github.com/PokemanMaster/GoChat/server/server/app/product/api"
	UserApi "github.com/PokemanMaster/GoChat/server/server/app/user/api"
	"github.com/PokemanMaster/GoChat/server/server/pkg/mid"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(mid.CORS())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	store, _ := redis.NewStore(10, "tcp", "47.113.104.184:6379", "123456", []byte("alkdnlakwdlawfhnolaqwfnlawm"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("/api/v1")
	{
		// 用户
		v1.GET("/captcha/image", UserApi.CaptchaImage)  // 验证码图片
		v1.POST("/user/register", UserApi.UserRegister) // 用户注册
		v1.POST("/user/login", UserApi.UserLogin)       // 用户登录
		v1.POST("user/logout", UserApi.UserLogout)      // 用户登出
		v1.PUT("/user", UserApi.UserUpdate)             // 更新用户
		v1.GET("/user/:id", UserApi.UserInfo)           // 用户信息

		// 好友
		v1.POST("/friend", FriendApi.CreateFriend)         // 添加好友
		v1.GET("/friends/:id", FriendApi.FriendLists)      // 好友列表
		v1.POST("/friends/search", FriendApi.SearchFriend) // 搜索好友
		v1.DELETE("/friend", FriendApi.DeleteFriend)       // 删除好友

		// 群聊
		v1.POST("/group/create", GroupApi.CreateGroup)            // 创建群聊
		v1.GET("/group/lists/:id", GroupApi.GroupLists)           // 群聊列表
		v1.POST("/group/join", GroupApi.JoinGroup)                // 加入群聊
		v1.POST("/group/friend/lists", GroupApi.GroupFriendLists) // 群聊好友

		// 聊天
		v1.GET("/chat/send", ChatApi.SendMessage)    // 发送且接收消息
		v1.POST("/chat/message", ChatApi.GetMessage) // 获取聊天历史消息
		v1.POST("/attach/upload", ChatApi.Upload)    // 上传文件

		// 商品
		v1.GET("/carousels", api.ListCarousels)              // 获取所有轮播图
		v1.GET("/products/categories", api2.ListCategories)  // 获取所有商品分类
		v1.GET("/products", api4.ListProducts)               // 获取所有商品详情
		v1.GET("/products/:id", api4.ShowProduct)            // 获取某个商品详情
		v1.GET("/products/param", api4.ListProductsParams)   // 获取所有商品参数
		v1.GET("/products/:id/param", api4.ShowProductParam) // 获取某个商品参数
		v1.GET("/products/brand", api4.ListProducts)         // 获取所有商品品牌
		v1.POST("/searches", api4.SearchProducts)            // 搜索商品

		// 商品排行榜
		v1.GET("rankings", api4.ListRanking)          // 排行榜/热门
		v1.GET("rankings/elec", api4.ListElecRanking) // 排行榜/家电排行
		v1.GET("rankings/acce", api4.ListAcceRanking) // 排行榜/配件排行

		// 购物车
		v1.POST("carts", CartApi.CreateCart)   // 创建购物车
		v1.GET("carts/:id", CartApi.ShowCart)  // 展示购物车
		v1.PUT("carts", CartApi.UpdateCart)    // 修改购物车
		v1.DELETE("carts", CartApi.DeleteCart) // 删除购物车

		// 收藏
		v1.POST("favorites", FavoriteApi.CreateFavorite)   // 创建收藏夹
		v1.GET("favorites/:id", FavoriteApi.ShowFavorites) // 展示收藏夹
		v1.DELETE("favorites", FavoriteApi.DeleteFavorite) // 删除收藏夹

		// 订单
		v1.POST("orders", OrderApi.CreateOrder)        // 创建订单
		v1.GET("orders/:num", OrderApi.ShowOrder)      // 获取订单
		v1.GET("user/:id/orders", OrderApi.ListOrders) // 获取某个用户所有订单

		// 支付
		v1.POST("pay", PaymentApi.CreatePay) // 支付
	}
	return r
}
