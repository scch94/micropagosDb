package response

type UpdateLastLoginResponse struct {
	Response     `json:"response"`
	RowsAffected int64 `json:"rowsaffected"`
}
