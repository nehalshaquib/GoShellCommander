package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nehalshaquib/GoShellCommander/config"
)

func authMiddleWare(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("authorization token missing")))
		ctx.Abort()
		return
	}
	if !isTokenValid(token) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid authorization token")))
		ctx.Abort()
		return
	}

	ctx.Next()
}

func isTokenValid(token string) bool {
	//TODO: Implement JWT token
	_, ok := config.AuthorizedTokens[token]
	return ok
}
