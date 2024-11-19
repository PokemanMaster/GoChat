package router

import (
	ChatApi "IMProject/app/chat/api"
	FriendApi "IMProject/app/friend/api"
	GroupApi "IMProject/app/group/api"
	UserApi "IMProject/app/user/api"
	"IMProject/pkg/docs"
	"IMProject/pkg/utils"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(utils.CORS())
	docs.SwaggerInfo.BasePath = "" // 设置 Swagger 文档的基础路径。
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v1")
	{
		// 用户
		v1.POST("/user/register", UserApi.UserRegister) // 用户注册
		v1.POST("/user/login", UserApi.UserLogin)       // 用户登录
		v1.PUT("/user/update", UserApi.UserUpdate)      // 更新用户
		v1.GET("/user/lists", UserApi.UserLists)        // 用户列表
		v1.GET("/user/:id", UserApi.UserInfo)           // 用户信息

		// 群聊
		v1.POST("/group/create", GroupApi.CreateGroup)            // 创建群聊
		v1.POST("/group/lists", GroupApi.GroupLists)              // 群聊列表
		v1.POST("/group/join", GroupApi.JoinGroup)                // 加入群聊
		v1.POST("/group/friend/lists", GroupApi.GroupFriendLists) // 群聊好友

		// 好友
		v1.POST("/friend/add", FriendApi.CreateFriend)    // 添加好友
		v1.POST("/friend/lists", FriendApi.FriendLists)   // 好友列表
		v1.POST("/friend/search", FriendApi.SearchFriend) // 搜索好友

		// 聊天
		v1.GET("/chat/send", ChatApi.SendMessage)    // 发送且接收消息
		v1.POST("/chat/message", ChatApi.GetMessage) // 获取聊天历史消息
		v1.POST("/attach/upload", ChatApi.Upload)    // 上传文件
	}
	return r
}
