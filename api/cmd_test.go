package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nehalshaquib/GoShellCommander/config"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestCmdHandler(t *testing.T) {
	config.Logger = zap.S()
	gin.SetMode(gin.TestMode)

	// Create a test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Create a request body
	reqBody := CmdHandlerRequest{
		CommandName: "go",
		Arguments:   []string{"version"},
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	require.NoError(t, err, "failed to encode request body")

	// Create a request
	c.Request, err = http.NewRequest("POST", "cmd", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a server and call the handler
	server := Server{}
	server.cmdHandler(c)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// TODO: Add more checks here, such as checking the response body
}
