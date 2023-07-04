package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nehalshaquib/GoShellCommander/util"
)

// CmdHandlerRequest represents the JSON structure of a command request.
type CmdHandlerRequest struct {
	CommandName string   `json:"command_name" binding:"required"`
	Arguments   []string `json:"arguments"`
}

// CmdHandlerResponse represents the JSON structure of a command response.
type CmdHandlerResponse struct {
	Result string `json:"result"`
}

// cmdHandler handles command execution requests.
func (server *Server) cmdHandler(ctx *gin.Context) {
	// Bind the JSON request to the CmdHandlerRequest struct
	req := CmdHandlerRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Execute the command using the util.RunCommand function
	out, err := util.RunCommand(ctx, req.CommandName, req.Arguments)
	if err != nil {
		// If the error indicates that the command was not found or recognized, return a not found response
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "not recognized") {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":  "command not found",
				"detail": err.Error(),
			})
			return
		}
		// Otherwise, return an internal server error response
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// Return the command execution result in a successful response
	ctx.JSON(http.StatusOK, CmdHandlerResponse{Result: out})
}
