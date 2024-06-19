package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/ins_log"

	handler "github.com/scch94/micropagosDb/utils/handlers"
	"github.com/scch94/micropagosDb/utils/middleware"
)

func SetupRouter(ctx context.Context) *gin.Engine {

	// create a new gin router and register the handlers
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Agregar middleware global
	router.Use(gin.Recovery())
	router.Use(middleware.GlobalMiddleware())

	h := handler.Handler{}
	//rutas
	router.GET("/", h.Welcome)
	router.POST("/insertMessage/:utfi", h.InserMessage)
	router.GET("/filter/:mobile/:shortNumber/:utfi", h.IsUserFilter)
	router.GET("/masks/:utfi", h.GetMaskPatterns)
	router.GET("/userdomain/:username/:utfi", h.GetUserDomain)
	router.GET("/usersInfo/:utfi", h.GetUsersInfo)
	router.NoRoute(notFoundHandler)

	return router
}

// Controlador para manejar rutas no encontradas
func notFoundHandler(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = ins_log.SetPackageNameInContext(ctx, "handler")
	ins_log.Errorf(ctx, "Route  not found: url: %v, method: %v", c.Request.RequestURI, c.Request.Method)
	c.JSON(http.StatusNotFound, nil)
}
