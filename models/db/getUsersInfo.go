package modeldb

import "database/sql"

type UserInfoDb struct {
	UserId        sql.NullString
	Username      sql.NullString
	UserPassword  sql.NullString
	UserLastLogin sql.NullString
	DomainName    sql.NullString
}

type UserInfo struct {
	UserId        string `json:"user_id"`
	Username      string `json:"user_username"`
	UserPassword  string `json:"user_password"`
	UserLastLogin string `json:"user_last_login"`
	DomainName    string `json:"domain_name"`
}

func (user *UserInfoDb) ConvertUser() UserInfo {
	return UserInfo{
		UserId:        user.UserId.String,
		Username:      user.UserId.String,
		UserPassword:  user.UserPassword.String,
		UserLastLogin: user.UserLastLogin.String,
		DomainName:    user.DomainName.String,
	}
}
