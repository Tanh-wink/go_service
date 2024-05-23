package utils

type ErrorBaseResponse struct {
	Code    int    `json:"code"`
	Msg string `json:"msg"`
}

type ErrorResponse struct {
	Error_Param		ErrorBaseResponse
	Error_Server	ErrorBaseResponse
	Error_Load		ErrorBaseResponse
	Error_Infer		ErrorBaseResponse
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{
		Error_Param: ErrorBaseResponse{Code:300, Msg:"请求参数错误"},
		Error_Server: ErrorBaseResponse{Code:500, Msg:"服务器错误"},
		Error_Load: ErrorBaseResponse{Code:1001, Msg:"加载时异常"},
		Error_Infer: ErrorBaseResponse{Code:1002, Msg:"推理时异常"},
	}
}

type ParamError struct {
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


