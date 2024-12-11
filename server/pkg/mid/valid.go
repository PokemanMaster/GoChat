package mid

import (
	"encoding/base64"
	"errors"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"

	"golang.org/x/exp/rand"

	"regexp"
)

// ValidateSearchInput 防止sql注入
func ValidateSearchInput(input string) (string, error) {
	// 允许字母、数字、空格和中文字符
	re := regexp.MustCompile(`^[\p{L}\p{N} ]+$`)
	if !re.MatchString(input) {
		return "", errors.New("invalid search input")
	}
	return input, nil
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

// TelephoneNumberIsTure 判断手机号真假
func TelephoneNumberIsTure(number string) (code int) {
	phoneNumber := `^1[0-9]{10}$`
	match, _ := regexp.MatchString(phoneNumber, number)
	if match {
		return e.SUCCESS
	} else {
		return e.ERROR
	}
}
