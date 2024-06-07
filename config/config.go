package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/scch94/Gconfiguration"
	"github.com/scch94/ins_log"
)

var Config MicropagosDatabaseConfiguration

type MicropagosDatabaseConfiguration struct {
	ServerPort           string         `json:"serverPort"`
	LogLevel             string         `json:"log_level"`
	Log_name             string         `json:"log_name"`
	QueryRetryCount      int            `json:"queryRetryCount"`
	DbQueryTimeout       int            `json:"dbQueryTimeout"`
	MySQLConnection      MySQLConfig    `json:"mysqlConnection"`
	MyPostgresConnection PostgresConfig `json:"myPostgresConnection"`
}

type MySQLConfig struct {
	Raven  ConnectionConfig `json:"raven"`
	Weaver ConnectionConfig `json:"weaver"`
}

type PostgresConfig struct {
	Raven  ConnectionConfig `json:"raven"`
	Weaver ConnectionConfig `json:"weaver"`
}

type ConnectionConfig struct {
	ConnectionString string `json:"connectionString"`
	MaxOpenConns     int    `json:"maxOpenConns"`
	MaxIdleConns     int    `json:"maxIdleConns"`
	ConnMaxLifeTime  int    `json:"conMaxLifeTime"`
	ConnMaxIdleTime  int    `json:"conMaxIdleTime"`
}

func Upconfig(ctx context.Context) error {
	//traemos el contexto y le setiamos el contexto actual
	ctx = ins_log.SetPackageNameInContext(ctx, "config")

	ins_log.Info(ctx, "starting to get the config struct ")
	err := Gconfiguration.GetConfig(&Config, "../config", "micropagosDbConfig.json")
	if err != nil {
		ins_log.Fatalf(ctx, "error in Gconfiguration.GetConfig() ", err)
		return err
	}
	return nil
}
func (m MicropagosDatabaseConfiguration) ConfigurationString() string {
	configJSON, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("Error al convertir la configuración a JSON: %v", err)
	}
	return string(configJSON)
}
