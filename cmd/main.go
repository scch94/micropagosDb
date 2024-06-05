package main

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/config"
	"github.com/scch94/micropagosDb/database"
	"github.com/scch94/micropagosDb/server"
)

func main() {

	// Creamos el contexto para esta ejecución
	ctx := context.Background()

	logFileName, err := initializeLogger()
	if err != nil {
		panic(err)
	}
	defer logFileName.Close()

	// Load configuration
	if err := config.Upconfig(ctx); err != nil {
		ins_log.Errorf(ctx, "error loading configuration: %v", err)
		return
	}

	// Inicializamos el logger
	ins_log.SetService("micropagosdatabase")
	ins_log.SetLevel(config.Config.LogLevel)
	ctx = ins_log.SetPackageNameInContext(ctx, "main")
	ins_log.Infof(ctx, "startig micropagos database module version : %+v", version())

	//inicialisamos la base de datos
	database.InitDb(ctx)

	//inicalisamos el servidor
	if err := server.StartServer(ctx); err != nil {
		ins_log.Errorf(ctx, "error starting server: %s", err.Error())
	}
}

func initializeLogger() (*os.File, error) {
	logDir := "../log"

	// Create the log directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}

	//Definimos el nombre del archivo
	today := time.Now().Format("2006-01-02_15")
	logFileName := filepath.Join(logDir, "micropagosdatabasegateway_"+today+".log")

	//abrimos el archivo del log
	file, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	//configuramos para abrir el log file y la consola
	multiWriter := io.MultiWriter(os.Stdout, file)
	ins_log.StartLoggerWithWriter(multiWriter)
	return file, nil
}

func version() string {
	return "1.0.0"
}
