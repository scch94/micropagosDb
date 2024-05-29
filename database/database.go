package database

import (
	"context"
	"database/sql"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/config"
)

var (
	DBusers       *sql.DB
	DBmessage     *sql.DB
	dbMessageOnce sync.Once
	dbUsersOnce   sync.Once
)

func InitDb(ctx context.Context) {
	NewMysqlMessageDb(ctx)
	NewMysqlUserDb(ctx)
}

func NewMysqlUserDb(ctx context.Context) {
	dbUsersOnce.Do(func() {
		ctx = ins_log.SetPackageNameInContext(ctx, "database")
		var err error
		DBusers, err = sql.Open("mysql", config.Config.MySQLConnection.Weaver.ConnectionString)
		if err != nil {
			ins_log.Fatalf(ctx, "cant open myssql database with string connection :%s , with error :%s", config.Config.MySQLConnection.Weaver.ConnectionString, err)
		}

		DBusers.SetMaxOpenConns(config.Config.MySQLConnection.Weaver.MaxOpenConns)
		DBusers.SetMaxIdleConns(config.Config.MySQLConnection.Weaver.MaxIdleConns)
		//DBusers.SetConnMaxLifetime(time.Duration(config.Config.MySQLConnection.Weaver.ConnMaxLifeTime) * time.Second)

		if err = DBusers.Ping(); err != nil {
			ins_log.Fatalf(ctx, "cant do ping to the database error :%s", err)
		}
		ins_log.Info(ctx, "conected to the mysql user database")

	})
}

func NewMysqlMessageDb(ctx context.Context) {
	dbMessageOnce.Do(func() {
		ctx = ins_log.SetPackageNameInContext(ctx, "database")
		var err error
		DBmessage, err = sql.Open("mysql", config.Config.MySQLConnection.Raven.ConnectionString)
		if err != nil {
			ins_log.Fatalf(ctx, "cant open myssql database with string connection :%s , with error :%s", config.Config.MySQLConnection.Raven.ConnectionString, err)
		}

		DBmessage.SetMaxOpenConns(config.Config.MySQLConnection.Raven.MaxOpenConns)
		DBmessage.SetMaxIdleConns(config.Config.MySQLConnection.Raven.MaxIdleConns)
		//DBmessage.SetConnMaxLifetime(time.Duration(config.Config.MySQLConnection.Raven.ConnMaxLifeTime) * time.Second)

		if err = DBmessage.Ping(); err != nil {
			ins_log.Fatalf(ctx, "cant do ping to the database error :%s", err)
		}
		ins_log.Info(ctx, "conected to the mysql message database")

	})
}

func GetDBUsers() *sql.DB {
	return DBusers
}

// GetDBMessage returns the MySQL connection pool for message database
func GetDBMessage() *sql.DB {
	return DBmessage
}

func StringToNull(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func TimeToNull(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{Time: t, Valid: false}
	}
	return sql.NullTime{Time: t, Valid: true}
}

func Uint64ToNull(i uint64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: int64(i), Valid: true}
}
