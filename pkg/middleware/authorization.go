package middleware

import (
	"errors"
	"net/http"

	"github.com/Ndraaa15/workshop-bcc/pkg/response"
	"github.com/gin-gonic/gin"
)

func (m *middleware) OnlyAdmin(ctx *gin.Context) {
	user, err := m.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		response.Error(ctx, http.StatusForbidden, "failed to authorize user", err)
		ctx.Abort()
		return
	}

	if user.Role != 1 {
		response.Error(ctx, http.StatusForbidden, "this endpoint can't be access", errors.New("user don't have access"))
		ctx.Abort()
		return
	}

	ctx.Next()
}
