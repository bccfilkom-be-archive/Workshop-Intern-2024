package rest

import (
	"net/http"

	"github.com/Ndraaa15/workshop-bcc/model"
	"github.com/Ndraaa15/workshop-bcc/pkg/response"
	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	token, err := r.service.UserService.Register(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register new user", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "success register new user", token)
}
