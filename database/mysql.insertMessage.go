package database

import (
	"context"
	"time"

	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/models/request"
)

const (
	mySQLInsertMessage = `INSERT INTO message ` +
		`(type, content, mobile_number, mobile_countryisocode, shortnumber, telco, created, routingtype, ` +
		`matchedpattern, serviceid, telcoid, sessionaction, sessionparameters_map, sessiontimeoutseconds, ` +
		`priority, clientid, url, accesstimeoutseconds, request_id, defaultaction_id, application_id, ` +
		`session_id, processed, millissincerequest, sessionapplicationname, sendafter, sendbefore, sent, ` +
		`status, accesstimeouthandlerqueuename, useunsupportedmobilesregistry, originname) ` +
		`VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
)

func InsertMessage(m *request.InsertMessageModel, ctx context.Context) error {

	// Establece el contexto actual
	ctx = ins_log.SetPackageNameInContext(ctx, "database")

	startTime := time.Now() // Captura el tiempo de inicio de la operación
	ins_log.Tracef(ctx, "attempting to insert the message into the database")
	ins_log.Tracef(ctx, "this is the query we will attempt to insert: %s", mySQLInsertMessage)

	//realizamos la consula
	db := GetDBMessage()
	result, err := db.Exec(mySQLInsertMessage,
		(m.Type),
		(m.Content),
		StringToNull(m.MobileNumber),
		StringToNull(m.MobileCountryISOCode),
		StringToNull(m.ShortNumber),
		StringToNull(m.Telco),
		TimeToNull(m.Created),
		StringToNull(m.RoutingType),
		StringToNull(m.MatchedPattern),
		StringToNull(m.ServiceID),
		StringToNull(m.TelcoID),
		StringToNull(m.SessionAction),
		(m.SessionParametersMap),
		Uint64ToNull(m.SessionTimeoutSeconds),
		Uint64ToNull(m.Priority),
		StringToNull(m.ClientID),
		StringToNull(m.URL),
		Uint64ToNull(m.AccessTimeoutSeconds),
		Uint64ToNull(m.RequestID),
		Uint64ToNull(m.DefaultActionID),
		Uint64ToNull(m.ApplicationID),
		Uint64ToNull(m.SessionID),
		TimeToNull(m.Processed),
		Uint64ToNull(m.MillisSinceRequest),
		StringToNull(m.SessionApplicationName),
		StringToNull(m.Sendafter),
		StringToNull(m.Sendbefore),
		TimeToNull(m.Sent),
		StringToNull(m.Status),
		StringToNull(m.AccessTimeoutHandlerQueuename),
		Uint64ToNull(m.UseUnsupportedMobilesRegistry),
		StringToNull(m.OriginName),
	)
	if err != nil {
		return err
	}
	duration := time.Since(startTime) // Calcula la duración de la operación
	ins_log.Infof(ctx, "Inserting the message into the database took: %v", duration)

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.Id = uint64(id)
	ins_log.Infof(ctx, "Message inserted successfully. The message was saved with the id %d ", m.Id)
	return nil
}
