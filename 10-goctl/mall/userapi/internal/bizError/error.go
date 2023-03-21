package bizError

// ParamError 业务逻辑错误对象
// 错误码:1101101 其中110是前缀,表示出现错误的微服务是userapi
// 1101是有效载荷,表示具体错误
var ParamError = NewBizError(1101101, "参数错误", make(map[string]interface{}))

type BizError struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewBizError(code int, msg string, data map[string]interface{}) *BizError {
	return &BizError{
		Code:    code,
		Message: msg,
		// 默认业务错误的JSON中不会返回数据
		Data: data,
	}
}

func (b *BizError) Error() string {
	return b.Message
}

func (b *BizError) GenResp() *BizErrResponse {
	return &BizErrResponse{
		Code:    b.Code,
		Message: b.Message,
		Data:    b.Data,
	}
}

type BizErrResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
