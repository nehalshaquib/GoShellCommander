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

	// Define a test case
	testCases := []struct {
		name         string
		reqBody      CmdHandlerRequest
		expectedCode int
	}{
		{
			name: "Test valid command",
			reqBody: CmdHandlerRequest{
				CommandName: "go",
				Arguments:   []string{"version"},
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "Test command not found",
			reqBody: CmdHandlerRequest{
				CommandName: "invalidCmd",
				Arguments:   []string{},
			},
			expectedCode: http.StatusNotFound, // Assuming your handler returns 400 for invalid commands
		},
		{
			name: "Test internal error",
			reqBody: CmdHandlerRequest{
				CommandName: "go",
				Arguments:   []string{"run random.go"},
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a test context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Create a request body
			reqBodyBytes, err := json.Marshal(tc.reqBody)
			require.NoError(t, err, "failed to encode request body")

			// Create a request
			c.Request, err = http.NewRequest("POST", "cmd", bytes.NewBuffer(reqBodyBytes))
			require.NoError(t, err, "failed to create request")

			// Create a server and call the handler
			server := Server{}
			server.cmdHandler(c)
			// Check the status code
			require.Equal(t, tc.expectedCode, w.Code)
		})
	}
}
