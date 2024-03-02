package rest

import (
	"net/http"

	"github.com/Ndraaa15/workshop-bcc/model"
	"github.com/Ndraaa15/workshop-bcc/pkg/response"
	"github.com/gin-gonic/gin"
)

func (r *Rest) RentBook(ctx *gin.Context) {
	param := model.RentBook{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind request", err)
		return
	}

	err = r.service.RentService.RentBook(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to rent book", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "success rent book", nil)
}
