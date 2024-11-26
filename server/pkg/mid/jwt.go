package mid

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/user/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义JWT加密的密钥：
// jwtKey 是一个用于加密和解密JWT的密钥。这里使用了一个简单的字符串作为密钥
var jwtKey = []byte("a_secret_create")

// Claims 定义Claims结构体：
// Claims 是一个自定义的结构体，用于表示JWT中的声明（claims）。
// 它包含了一个 UserId 字段，表示用户的ID，以及继承了 jwt.StandardClaims 结构体，
// 该结构体是JWT库中预定义的一些标准声明（如过期时间、发放时间、发放者等）。
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	//token的有效时间
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 时间
			IssuedAt:  time.Now().Unix(),     //token发放的时间
			Issuer:    "jkdev.cn",            //是谁发放的token
			Subject:   "user token",          //主题
		},
	}
	//用jwt这个密钥来生成我们的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		zap.L().Error("权限不足", zap.String("app.user.pkg.mid.jwt", err.Error()))
		return "", err
	}
	return tokenString, nil
}

// Token 验证 HTTP 请求的身份认证。
func Token() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 authorization header 中的 token
		tokenString := ctx.GetHeader("Authorization")
		// 然后检查它的格式是否正确。
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			// 用于中止当前 HTTP 请求链中的所有处理程序和中间件。
			ctx.Abort()
			return
		}
		// 因为Bearer字符占了7位，所以从7开始截取字符
		tokenString = tokenString[7:]
		// 解析传入的 JWT（JSON Web Token）字符串并返回 token 对象、token 中的声明（claims）和任何错误。
		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 402,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		//token通过验证, 获取claims中的UserID
		userId := claims.UserId
		// 连接数据库
		//DB := common.GetDB()
		var user model.User
		// 查询数据库中符合条件的第一个记录，并将结果存储在给定的变量 user 中。
		db.DB.First(&user, userId)
		// 验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 403,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		// 如果用户存在，则将用户信息写入 Gin 上下文中，并将请求传递给后续的处理程序。
		// 在 Gin 的上下文（Context）中设置一个键值对。
		// 可以在后续的处理程序中使用 ctx.Get("user") 获取存储在上下文中的 user 对象，以便在需要时使用它。
		ctx.Set("user", user)
		// 当前中间件已处理完请求，并将请求传递给下一个中间件或路由处理程序。
		// 如果没有调用 ctx.Next()，则当前中间件将停止执行，不会将请求传递给后续的处理程序或中间件。
		ctx.Next()
	}
}

// ParseToken 解析token
// ParseToken 函数用于解析JWT。它接收一个JWT字符串作为参数，并返回解析后的 Token 对象、Claims 对象和可能出现的错误。
// 函数中的主要步骤包括：
// 创建一个空的 Claims 对象用于存储解析后的声明。
// 调用 jwt.ParseWithClaims 方法解析JWT，并将解析后的声明存储到 claims 对象中。
// 在解析过程中，还需要提供一个回调函数，用于验证JWT的签名。这里的回调函数简单地返回预定义的 jwtKey 作为密钥，表示对JWT的签名进行验证。
// 最后，函数返回解析后的 Token 对象、Claims 对象和可能的错误信息。
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
