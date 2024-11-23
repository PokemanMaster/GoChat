package router

import (
	ChatApi "IMProject/app/chat/api"
	FriendApi "IMProject/app/friend/api"
	GroupApi "IMProject/app/group/api"
	ProductApi "IMProject/app/product/api"
	UserApi "IMProject/app/user/api"
	"IMProject/pkg/docs"
	"IMProject/pkg/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(utils.CORS())
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	store, _ := redis.NewStore(10, "tcp", "47.113.104.184:6379", "123456", []byte("alkdnlakwdlawfhnolaqwfnlawm"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("/api/v1")
	{
		// 用户
		v1.POST("/user/register", UserApi.UserRegister)     // 用户注册
		v1.POST("/user/login", UserApi.UserLogin)           // 用户登录
		v1.PUT("/user/update", UserApi.UserUpdate)          // 更新用户
		v1.GET("/user/captcha/image", UserApi.CaptchaImage) // 验证码图片
		v1.GET("/user/lists", UserApi.UserLists)            // 用户列表
		v1.GET("/user/:id", UserApi.UserInfo)               // 用户信息
		v1.POST("user/logout", UserApi.UserLogout)          // 用户登出

		// 好友
		v1.POST("/friend/create", FriendApi.CreateFriend)  // 添加好友
		v1.GET("/friend/lists/:id", FriendApi.FriendLists) // 好友列表
		v1.POST("/friend/search", FriendApi.SearchFriend)  // 搜索好友
		v1.POST("/friend/delete", FriendApi.DeleteFriend)  // 删除好友

		// 群聊
		v1.POST("/group/create", GroupApi.CreateGroup)            // 创建群聊
		v1.GET("/group/lists/:id", GroupApi.GroupLists)           // 群聊列表
		v1.POST("/group/join", GroupApi.JoinGroup)                // 加入群聊
		v1.POST("/group/friend/lists", GroupApi.GroupFriendLists) // 群聊好友

		// 聊天
		v1.GET("/chat/send", ChatApi.SendMessage)    // 发送且接收消息
		v1.POST("/chat/message", ChatApi.GetMessage) // 获取聊天历史消息
		v1.POST("/attach/upload", ChatApi.Upload)    // 上传文件

		v1.GET("/carousels", ProductApi.ListCarousels)             // 获取所有轮播图
		v1.GET("/products/categories", ProductApi.ListCategories)  // 获取所有商品分类
		v1.GET("/products", ProductApi.ListProducts)               // 获取所有商品详情
		v1.GET("/products/:id", ProductApi.ShowProduct)            // 获取某个商品详情
		v1.GET("/products/param", ProductApi.ListProductsParams)   // 获取所有商品参数
		v1.GET("/products/:id/param", ProductApi.ShowProductParam) // 获取某个商品参数
		v1.GET("/products/brand", ProductApi.ListProducts)         // 获取所有商品品牌
		v1.POST("/searches", ProductApi.SearchProducts)            // 搜索商品
	}
	return r
}
