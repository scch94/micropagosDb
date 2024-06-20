package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/database"
	"github.com/scch94/micropagosDb/models/request"
	"github.com/scch94/micropagosDb/models/response"
)

func (handler *Handler) UpdateUsersLastLogin(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = ins_log.SetPackageNameInContext(ctx, "handler")

	var request request.UsersToUpdate
	// Utiliza el método BindJSON de Gin para vincular los datos del cuerpo de la solicitud a la estructura request
	if err := c.BindJSON(&request); err != nil {
		{
			ins_log.Errorf(ctx, "error when we try to get the json petition")
			response := response.NewResponse(1, err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	ins_log.Tracef(ctx, "this is the data that we recibed in the petition to update last login  %+v", request)

	//si no hay usuarios para actualizar devolvemos q no se actualizaron usuarios y no hacemos el intento en la base
	if len(request.Users) == 0 {
		ins_log.Infof(ctx, "no rows afectted")
		responses := response.UpdateLastLoginResponse{
			Response: response.Response{
				Result:  0,
				Message: "updated successfully",
			},
			RowsAffected: 0,
		}
		c.JSON(http.StatusOK, responses)
		return
	}
	ins_log.Info(ctx, "starting to update")

	rowsAffected, err := database.UpdateUsersLastLogin(&request, ctx)
	if err != nil {
		ins_log.Errorf(ctx, "error when we try to UpdateUsersLastLogin :%v", err)
		response := response.NewResponse(1, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	responses := response.UpdateLastLoginResponse{
		Response: response.Response{
			Result:  0,
			Message: "updated successfully",
		},
		RowsAffected: rowsAffected,
	}

	c.JSON(http.StatusOK, responses)

}
