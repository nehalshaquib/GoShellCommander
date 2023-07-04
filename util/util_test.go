package util

import (
	"context"
	"errors"
	"testing"

	"github.com/nehalshaquib/GoShellCommander/config"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestRunCommand(t *testing.T) {
	ctx := context.Background()
	config.Logger = zap.S()

	// Define a test case
	testCases := []struct {
		name    string
		cmd     string
		args    []string
		want    string
		wantErr error
	}{
		{
			name:    "Test valid command",
			cmd:     "go",
			args:    []string{"version"},
			want:    "go version",
			wantErr: nil,
		},
		{
			name:    "Test invalid command",
			cmd:     "invalidCmd",
			args:    []string{},
			want:    "",
			wantErr: errors.New("not found"),
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := RunCommand(ctx, tc.cmd, tc.args)

			if got != tc.want {
				require.Contains(t, got, tc.want)
			}

			if (err != nil && tc.wantErr == nil) || (err == nil && tc.wantErr != nil) || (err != nil && tc.wantErr != nil && err.Error() != tc.wantErr.Error()) {
				require.ErrorContains(t, err, tc.wantErr.Error())
			}
		})
	}
}
