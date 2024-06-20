package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/models/request"
)

// Función que genera la sentencia SQL
func generateUpdateSQL(usersData *request.UsersToUpdate) string {
	// Inicializa el principio de la consulta
	query := "UPDATE user SET lastlogin = CASE"

	// Construye las condiciones CASE
	var userNames []string
	for _, user := range usersData.Users {
		query += fmt.Sprintf(" WHEN username = '%s' THEN '%s'", user.UserName, user.LoginTime)
		userNames = append(userNames, fmt.Sprintf("'%s'", user.UserName))
	}

	// Finaliza la consulta con la cláusula WHERE
	query += fmt.Sprintf(" ELSE lastlogin END WHERE username IN (%s);", strings.Join(userNames, ", "))

	return query
}

func UpdateUsersLastLogin(usersData *request.UsersToUpdate, ctx context.Context) (int64, error) {
	// Establece el nombre del paquete en el contexto para el logging
	ctx = ins_log.SetPackageNameInContext(ctx, "database")

	startTime := time.Now() // Captura el tiempo de inicio de la operación
	ins_log.Tracef(ctx, "attempting to update the lastlogin of the users into the database")

	//generamos la consulta
	mySQLUpdateUsersLastLogin := generateUpdateSQL(usersData)
	ins_log.Tracef(ctx, "this is the query we will attempt to insert: %s", mySQLUpdateUsersLastLogin)

	//realizamos la consula
	db := GetDBUsers()
	results, err := db.Exec(mySQLUpdateUsersLastLogin)
	if err != nil {
		return 0, err
	}
	duration := time.Since(startTime) // Calcula la duración de la operación
	ins_log.Infof(ctx, "Updating the users last login, in the database took: %v", duration)

	// comprobamos el número de filas afectadas:
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		ins_log.Errorf(ctx, "error getting rows affected: %v", err)
		return 0, err
	}
	ins_log.Tracef(ctx, "Number of rows affected: %d", rowsAffected)

	return rowsAffected, nil
}
