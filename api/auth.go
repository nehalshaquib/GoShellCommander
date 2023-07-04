package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nehalshaquib/GoShellCommander/config"
)

// authMiddleWare is a middleware function that checks for a valid authorization token.
func authMiddleWare(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	// Check if the token is missing
	if token == "" {
		// If the token is missing an unauthorized response and abort the request
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("authorization token missing")))
		ctx.Abort()
		return
	}

	// Check if the token is valid
	if !isTokenValid(token) {
		// If the token is invalid, return an unauthorized response and abort the request
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid authorization token")))
		ctx.Abort()
		return
	}

	// If the token is valid, proceed to the next middleware or handler
	ctx.Next()
}

// isTokenValid checks if the provided token is valid.
func isTokenValid(token string) bool {
	//TODO: Implement JWT token
	_, ok := config.AuthorizedTokens[token]
	return ok
}
