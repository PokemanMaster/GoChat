package resp

import (
	"encoding/json"
	"errors"
	"fmt"
	conf "github.com/PokemanMaster/GoChat/server/pkg/utils"
	"github.com/go-playground/validator/v10"
)

// ErrorResponse 标准化的错误响应
func ErrorResponse(err error) Response {
	// 校验错误
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return Response{
				Status: 40001,
				Msg:    fmt.Sprintf("%s%s", field, tag),
			}
		}
	}
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return Response{
			Status: 40001,
			Msg:    "参数类型不匹配",
		}
	}
	return Response{
		Status: 40001,
		Msg:    "错误类型无法识别",
	}
}
