package response

type IsFilterResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}
