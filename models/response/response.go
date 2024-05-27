package response

type Response struct {
	Result  uint8  `json:"result"`
	Message string `json:"message"`
}

func NewResponse(result uint8, message string) *Response {
	return &Response{
		Result:  result,
		Message: message,
	}
}
