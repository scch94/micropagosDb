package modeldb

import "database/sql"

type MaskDb struct {
	ID             sql.NullString `json:"id"`
	ShortNumber    sql.NullString `json:"short_number"`
	MaskPattern    sql.NullString `json:"mask_pattern"`
	MinLength      sql.NullString `json:"min_length"`
	MaxLength      sql.NullString `json:"max_length"`
	ExcludePattern sql.NullString `json:"exclude_pattern"`
	Direction      sql.NullString `json:"direction"`
	ApplicationID  sql.NullString `json:"application_id"`
}

type Mask struct {
	ID             string `json:"id"`
	ShortNumber    string `json:"short_number"`
	MaskPattern    string `json:"mask_pattern"`
	MinLength      string `json:"min_length"`
	MaxLength      string `json:"max_length"`
	ExcludePattern string `json:"exclude_pattern"`
	Direction      string `json:"direction"`
	ApplicationID  string `json:"application_id"`
}

func (mask *MaskDb) ConvertMask() Mask {
	return Mask{
		ID:             mask.ID.String,
		ShortNumber:    mask.ShortNumber.String,
		MaskPattern:    mask.MaskPattern.String,
		MinLength:      mask.MinLength.String,
		MaxLength:      mask.MaxLength.String,
		ExcludePattern: mask.ExcludePattern.String,
		Direction:      mask.Direction.String,
		ApplicationID:  mask.ApplicationID.String,
	}
}
