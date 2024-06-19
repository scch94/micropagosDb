package response

import modeldb "github.com/scch94/micropagosDb/models/db"

type GetUsersInfoResponse struct {
	Response
	Users []modeldb.UserInfo `json:"users"`
}
