package mid

import (
	"encoding/base64"
	"errors"
	"github.com/PokemanMaster/GoChat/server/pkg/e"

	"golang.org/x/exp/rand"

	"regexp"
)

// ValidateSearchInput 防止sql注入
func ValidateSearchInput(input string) (string, int, error) {
	// 允许字母、数字、空格和中文字符
	re := regexp.MustCompile(`^[\p{L}\p{N} ]+$`)
	if !re.MatchString(input) {
		return "", e.ERROR_DATABASE, errors.New("invalid search input")
	}
	return input, e.SUCCESS, nil
}

func GenerateRandomKey() string {
	// 生成32字节的随机数据
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err) // 确保生成密钥时没有问题
	}
	// 返回 base64 编码的密钥
	return base64.StdEncoding.EncodeToString(key)
}
