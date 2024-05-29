package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/scch94/ins_log"
	modeldb "github.com/scch94/micropagosDb/models/db"
	"github.com/scch94/micropagosDb/models/request"
)

const (
	mySQLGetDomain = `SELECT d.name, u.username, u.password FROM user u JOIN domain d ON u.domain_id = d.id where u.username=?`
)

func GetUserDomain(r request.GetUserDomain, ctx context.Context) (*modeldb.UserDomain, error) {

	// Establece el contexto actual
	ctx = ins_log.SetPackageNameInContext(ctx, "database")

	// Crear el modelo de dominio para almacenar los resultados
	var domainModel modeldb.UserDomain

	ins_log.Tracef(ctx, "starting to get the domain for the user :%s", r.UserName)
	startTime := time.Now()

	ins_log.Tracef(ctx, "this is the QUERY: %s and the params: Username=%s,", mySQLGetDomain, r.UserName)

	var err error
	//realizamos la consula
	for i := 0; i < 3; i++ {

		queryCtx, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
		defer cancel()

		// Realizar la consulta
		db := GetDBUsers()
		err = db.QueryRowContext(queryCtx, mySQLGetDomain, r.UserName).Scan(&domainModel.Domainname, &domainModel.Username, &domainModel.Password)

		if err == nil {
			// Consulta exitosa, salir del bucle
			break
		}

		// Si hay un error, registrar el intento
		ins_log.Tracef(ctx, "getUserDomain error on attempt %d: %v", i+1, err)
	}

	switch {
	case err == sql.ErrNoRows:
		ins_log.Infof(ctx, "didnt find domain for the user %s", r.UserName)
		domainModel.Domainname = ""
		domainModel.Result = err.Error()
		err = nil
	case err != nil:
		ins_log.Errorf(ctx, "query error %v", err)
	default:
		domainModel.Result = "the domain name is: " + domainModel.Domainname
	}

	// Calcula la duración de la operacion
	duration := time.Since(startTime)
	ins_log.Infof(ctx, "the query in the database tooks: %v", duration)
	ins_log.Infof(ctx, "and this is the domain: %v", domainModel.Domainname)

	return &domainModel, err

}
