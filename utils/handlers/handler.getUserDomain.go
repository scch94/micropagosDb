package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/database"
	"github.com/scch94/micropagosDb/models/request"
	"github.com/scch94/micropagosDb/models/response"
)

func (handler *Handler) GetUserDomain(c *gin.Context) {

	ctx := c.Request.Context()
	ctx = ins_log.SetPackageNameInContext(ctx, "handler")

	request := request.GetUserDomain{
		UserName: c.Param("username"),
	}

	ins_log.Infof(ctx, "starting to get the domain for the user %s", request.UserName)

	//REALISAMOS LA CONSULTA
	domainResponse, err := database.GetUserDomain(request, ctx)
	if err != nil {
		ins_log.Errorf(ctx, "error getting Domain response")
		response := response.NewResponse(1, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//creamos la respuesta
	response := response.DomainResponse{
		Response: response.Response{
			Result:  0,
			Message: domainResponse.Result,
		},
		DomainName: domainResponse.Domainname,
		Username:   domainResponse.Username,
		Password:   domainResponse.Password,
	}
	c.JSON(http.StatusOK, response)

}
