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
	mySQLGetMask = `SELECT ID, SHORTNUMBER, MASKPATTERN, MINLENGTH, MAXLENGTH, EXCLUDEPATTERN, DIRECTION, APPLICATION_ID FROM MASK;`
)

func GetMask(ctx context.Context) ([]modeldb.Mask, error) {

	// Establece el contexto actual
	ctx = ins_log.SetPackageNameInContext(ctx, "database")

	//crear el modelo para poder almacenar los resultados
	var masks []modeldb.Mask

	ins_log.Tracef(ctx, "starting to get the mask patters")
	startTime := time.Now()

	ins_log.Tracef(ctx, "this is the QUERY: %s ", mySQLGetMask)
	var err error
	var rows *sql.Rows
	//realizamos la consula
	for i := 0; i < config.Config.QueryRetryCount; i++ {

		queryCtx, cancel := context.WithCancel(ctx)

		// Realizar la consulta
		db := GetDBMessage()
		rows, err = db.QueryContext(queryCtx, mySQLGetMask)
		defer cancel()
		if err == nil {
			// Consulta exitosa, salir del bucle
			break
		}

		// Si hay un error, registrar el intento
		ins_log.Tracef(ctx, "getUserDomain error on attempt %d: %v", i+1, err)

	}
	if err != nil {
		ins_log.Errorf(ctx, "query error %v", err)
		return masks, err
	}
	defer rows.Close()

	// Procesar los resultados de la consulta
	for rows.Next() {
		var maskDb modeldb.MaskDb
		err := rows.Scan(&maskDb.ID, &maskDb.ShortNumber, &maskDb.MaskPattern, &maskDb.MinLength, &maskDb.MaxLength, &maskDb.ExcludePattern, &maskDb.Direction, &maskDb.ApplicationID)
		if err != nil {
			ins_log.Errorf(ctx, "error scanning row: %v", err)
			continue // Continuar con la siguiente fila
		}
		mask := maskDb.ConvertMask()
		masks = append(masks, mask)
	}
	duration := time.Since(startTime)
	ins_log.Infof(ctx, "the query in the database tooks: %v", duration)
	return masks, nil
}
