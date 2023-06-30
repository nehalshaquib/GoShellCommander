package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nehalshaquib/GoShellCommander/util"
)

type CmdHandlerRequest struct {
	CommandName string   `json:"command_name" binding:"required"`
	Arguments   []string `json:"arguments"`
}

type CmdHandlerResponse struct {
	Result string `json:"result"`
}

func (server *Server) cmdHandler(ctx *gin.Context) {
	req := CmdHandlerRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	out := ``
	out, err = util.RunCommand(ctx, req.CommandName, req.Arguments)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "not recognized") {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":  "command not found",
				"detail": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, CmdHandlerResponse{Result: out})
}
