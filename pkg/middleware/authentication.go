package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Ndraaa15/workshop-bcc/model"
	"github.com/Ndraaa15/workshop-bcc/pkg/response"
	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.Error(ctx, http.StatusUnauthorized, "empty token", errors.New(""))
		ctx.Abort()
		return
	}

	token := strings.Split(bearer, " ")[1]
	userId, err := m.jwtAuth.ValidateToken(token)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed validate token", err)
		ctx.Abort()
		return
	}

	user, err := m.service.UserService.GetUser(model.UserParam{
		ID: userId,
	})
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed get user", err)
		ctx.Abort()
		return
	}

	ctx.Set("user", user)

	ctx.Next()
}
