package utils

type BaseResponse struct {
	Code    int    `json:"code"`
	Msg string `json:"msg"`
}

type ErrorResponse struct {
	Error_Server BaseResponse
	Error_Load BaseResponse
	Error_Infer BaseResponse
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{
		Error_Server: BaseResponse{Code:500, Msg:"服务器错误"},
		Error_Load: BaseResponse{Code:1001, Msg:"加载时异常"},
		Error_Infer: BaseResponse{Code:1002, Msg:"推理时异常"},
	}
}

type TraceBackError struct {
	Code    int    `json:"code"`
	Msg string `json:"msg"`
}

type LoadError struct {
	Code    int    `json:"code"`
	Msg string `json:"msg"`
}

type InferError struct {
	Code    int    `json:"code"`
	Msg string `json:"msg"`
}

type ServerError struct {
	Code    int    `json:"code"`
	Msg string `json:"msg"`
}
