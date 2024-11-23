package utils

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

// 用于签发和解析 JWT 的密钥
var secretKey = []byte("your_secret_key")

// 生成JWT
func generateToken(sessionID string) (string, error) {

	// 创建一个JWT声明
	claims := jwt.MapClaims{
		"session_id": sessionID,
		"exp":        time.Now().Add(time.Hour * 1).Unix(), // 设置过期时间为1小时
		"iat":        time.Now().Unix(),                    // 创建时间
	}

	// 创建 JWT 对象并签署
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签署 JWT 并返回生成的字符串
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析JWT
func parseToken(tokenString string) (*jwt.Token, error) {
	// 解析 token 时，需要提供一个验证函数来验证签名
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保 token 使用的是我们期望的签名算法
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥用于验证签名
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// 验证 JWT 的中间件
func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 获取 Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// 提取 Bearer Token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 解析 token
		token, err := parseToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// 获取 session_id
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			sessionID := claims["session_id"].(string)
			// 将 session_id 保存到上下文中，以便后续使用
			ctx := context.WithValue(r.Context(), "session_id", sessionID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
	})
}
