package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/database"
	"github.com/scch94/micropagosDb/models/response"
)

func (handler *Handler) GetUsersInfo(c *gin.Context) {

	ctx := c.Request.Context()
	ctx = ins_log.SetPackageNameInContext(ctx, "handler")

	ins_log.Info(ctx, "starting to get users info")

	//realizamos la consulta
	users, err := database.GetUsersInfo(ctx)
	if err != nil {
		ins_log.Errorf(ctx, "error getting users info in the database: %v", err)
		response := response.NewResponse(1, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.GetUsersInfoResponse{
		Response: response.Response{
			Result:  0,
			Message: "the number of users is: " + strconv.Itoa(len(users)),
		},
		Users: users,
	}
	c.JSON(http.StatusOK, response)
}
