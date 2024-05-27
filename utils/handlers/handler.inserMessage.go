package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/database"
	"github.com/scch94/micropagosDb/models/request"
	"github.com/scch94/micropagosDb/models/response"
)

func (handler *Handler) InserMessage(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = ins_log.SetPackageNameInContext(ctx, "handler")

	var request request.InsertMessageModel
	// Utiliza el método BindJSON de Gin para vincular los datos del cuerpo de la solicitud a la estructura request
	if err := c.BindJSON(&request); err != nil {
		{
			ins_log.Errorf(ctx, "error when we try to get the json petition")
			response := response.NewResponse(1, err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}

	}
	ins_log.Tracef(ctx, "this is the data that we recibed in the petition to insert the message %s", request)
	ins_log.Info(ctx, "starting to insert message")

	//realizamos la consulta
	err := database.InsertMessage(&request, ctx)
	if err != nil {
		ins_log.Error(ctx, "error inserting message insertMessage(request) :")
		response := response.NewResponse(1, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	message := "el mensaje se inserto correctamente con el id" + strconv.FormatUint(request.Id, 10)

	response := response.NewResponseMessage(0, message, request.Id)
	c.JSON(http.StatusOK, response)

}
