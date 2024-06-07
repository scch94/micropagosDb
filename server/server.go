package server

import (
	"context"
	"net/http"

	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/config"
	"github.com/scch94/micropagosDb/utils/router"
)

func StartServer(ctx context.Context) error {

	// Agregamos el valor "packageName" al contexto
	ctx = ins_log.SetPackageNameInContext(ctx, "server")

	ins_log.Infof(ctx, "Starting server on Port: %s", config.Config.ServerPort)

	router := router.SetupRouter(ctx)
	serverConfig := &http.Server{
		Addr:    config.Config.ServerPort,
		Handler: router,
	}
	err := serverConfig.ListenAndServe()
	if err != nil {
		ins_log.Errorf(ctx, "cant connect to the server: %+v", err)
		return err
	}
	return nil
}
