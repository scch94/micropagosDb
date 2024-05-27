package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/ins_log"
	"github.com/scch94/micropagosDb/database"
	"github.com/scch94/micropagosDb/models/request"
	"github.com/scch94/micropagosDb/models/response"
)

func (handler *Handler) IsUserFilter(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = ins_log.SetPackageNameInContext(ctx, "handler")

	//creamos la structura donde vamos a guardar el request
	request := request.IsUserFilter{
		ShortNumber: c.Param("shortNumber"),
		Mobile:      c.Param("mobile"),
	}

	ins_log.Infof(ctx, "starting to check if the shortnumber %s with the mobile %s are filter", request.ShortNumber, request.Mobile)

	isUserFilterResponse, err := database.IsUserFilter(request, ctx)
	if err != nil {
		ins_log.Errorf(ctx, "error getting filter response")
		response := response.NewResponse(1, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := response.IsFilterResponse{
		Result:  isUserFilterResponse.Result,
		Message: isUserFilterResponse.Comment,
	}
	c.JSON(http.StatusOK, response)
}
