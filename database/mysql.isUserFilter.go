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
	mySQLIsFilter = `SELECT id, added, comment, mobile_countryisocode, mobile_number, shortnumber ` +
		`FROM filter ` +
		`WHERE mobile_number = ? AND ` +
		`shortnumber = ? AND ` +
		`userremoved IS NULL AND ` +
		`(direction IS NULL OR direction = 'OUT')`
)

func IsUserFilter(r request.IsUserFilter, ctx context.Context) (*modeldb.IsUserFilter, error) {

	//traemos el contexto y le setiamos el contexto actual
	ctx = ins_log.SetPackageNameInContext(ctx, "database")

	//creamos la variable que guardara la respuesta de la db
	var filterModel modeldb.IsUserFilter

	//creamos variable que Captura el tiempo de inicio de la operación
	startTime := time.Now()
	ins_log.Tracef(ctx, "Checking if the combination of destination and origin mobile numbers is not filtered.")

	ins_log.Tracef(ctx, "this is the QUERY: %s and the params: mobile_number=%s, shortnumber=%s", mySQLIsFilter, r.Mobile, r.ShortNumber)

	// Ejecutamos la consulta directamente
	err := DBmessage.QueryRow(mySQLIsFilter, r.Mobile, r.ShortNumber).Scan(
		&filterModel.Id,
		&filterModel.Added,
		&filterModel.Comment,
		&filterModel.MobileCountryisocode,
		&filterModel.MobileNumber,
		&filterModel.ShortNumber,
	)

	//controlamos los errores, si no hay mas filas no es un error ! es lo que siempre esperamos
	switch {
	case err == sql.ErrNoRows:
		ins_log.Infof(ctx, "the combination of destination and origin mobile numbers is not filter")
		filterModel.Result = false
		filterModel.Comment = ""
		err = nil
	case err != nil:
		ins_log.Fatalf(ctx, "query error %v", err)
	default:
		ins_log.Infof(ctx, "the combination of destination and origin mobile numbers is filter")
		filterModel.Result = true
	}
	// Calcula la duración de la operacion
	duration := time.Since(startTime)
	ins_log.Infof(ctx, "the query in the database tooks: %v", duration)
	ins_log.Tracef(ctx, "and this is the information of the database: %v", filterModel)

	return &filterModel, err
}
