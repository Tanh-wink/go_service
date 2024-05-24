package dto

type ProductRecResponse struct {
	Cart  []string `json:"cart"`
	Wish  []string `json:"wish"`
	Trade []string `json:"trade"`
	Rec   []string `json:"rec"`
	Hot   []string `json:"hot"`
}

type Response struct {
	Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
