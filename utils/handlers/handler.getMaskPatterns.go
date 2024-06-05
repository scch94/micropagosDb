package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/database"
	"github.com/scch94/micropagosDb/models/response"
)

func (handler *Handler) GetMaskPatterns(c *gin.Context) {

	ctx := c.Request.Context()
	ctx = ins_log.SetPackageNameInContext(ctx, "handler")

	ins_log.Infof(ctx, "starting to get the maskpatterns")

	masksResponse, err := database.GetMask(ctx)
	if err != nil {
		ins_log.Errorf(ctx, "error getting masks")
		response := response.NewResponse(1, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//creamos la respuesta
	response := response.MaskResponse{
		Response: response.Response{
			Result:  0,
			Message: "getting mask completed!",
		},
		Masks: masksResponse,
	}
	c.JSON(http.StatusOK, response)
}
