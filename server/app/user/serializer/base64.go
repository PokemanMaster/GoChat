package serializer

// Base64 验证码图片
type Base64 struct {
	CodeId string `json:"code_id"`
	Base64 string `json:"base64"`
}

func BuildBase64(codeId string, base64 string) Base64 {
	return Base64{
		CodeId: codeId,
		Base64: base64,
	}
}
