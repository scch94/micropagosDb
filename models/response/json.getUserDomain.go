package response

type DomainResponse struct {
	Response
	DomainName string `json:"domain"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}
