package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/config"
	modeldb "github.com/scch94/micropagosDb/models/db"
)

const (
	mySQLGetUsersInfo = `SELECT u.id, u.username, u.password, u.lastlogin, d.name FROM user u JOIN domain d ON d.id = u.domain_id JOIN user_role AS ur ON ur.user_id = u.id JOIN role AS r ON r.id = ur.role_id WHERE r.name = 'webservice' AND r.system = 'raven'`
)

func GetUsersInfo(ctx context.Context) ([]modeldb.UserInfo, error) {

	//traemos el contexto y le setiamos el contexto actual
	ctx = ins_log.SetPackageNameInContext(ctx, "database")

	//creamos la variable que guardara la respuesta de la db
	var users []modeldb.UserInfo

	ins_log.Tracef(ctx, "starting to get the users information")
	startTime := time.Now()

	ins_log.Tracef(ctx, "this is the QUERY: %s", mySQLGetUsersInfo)
	var err error
	var rows *sql.Rows
	for i := 0; i < config.Config.QueryRetryCount; i++ {

		queryCtx, cancel := context.WithCancel(ctx)

		// Realizar la consulta
		db := GetDBUsers()
		rows, err = db.QueryContext(queryCtx, mySQLGetUsersInfo)
		defer cancel()
		if err == nil {
			// Consulta exitosa, salir del bucle
			break
		}

		// Si hay un error, registrar el intento
		ins_log.Tracef(ctx, "GetUsersInfo error on attempt %d: %v", i+1, err)

	}
	if err != nil {
		ins_log.Errorf(ctx, "query error %v", err)
		return users, err
	}
	defer rows.Close()
	// Procesar los resultados de la consulta
	for rows.Next() {
		var userInfoDb modeldb.UserInfoDb
		err := rows.Scan(&userInfoDb.UserId, &userInfoDb.Username, &userInfoDb.UserPassword, &userInfoDb.UserLastLogin, &userInfoDb.DomainName)
		if err != nil {
			ins_log.Errorf(ctx, "error scanning row: %v", err)
			continue // Continuar con la siguiente fila
		}
		user := userInfoDb.ConvertUser()
		users = append(users, user)
	}
	duration := time.Since(startTime)
	ins_log.Infof(ctx, "the query in the database tooks: %v", duration)

	return users, nil
}
