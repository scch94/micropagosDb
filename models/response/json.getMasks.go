package response

import modeldb "github.com/scch94/micropagosDb/models/db"

type MaskResponse struct {
	Response
	Masks []modeldb.Mask `json:"masks"`
}
