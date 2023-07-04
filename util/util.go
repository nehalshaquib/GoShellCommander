package util

import (
	"context"
	"errors"
	"os/exec"
	"runtime"

	"github.com/nehalshaquib/GoShellCommander/config"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

// RunCommand executes a shell command with the given command and arguments.
// It captures the combined output (stdout and stderr) of the command and returns it as a string.
func RunCommand(ctx context.Context, cmd string, args []string) (string, error) {
	// Get the logger instance from the config package
	log = config.Logger

	// current operating system
	log.Infoln("current os:", runtime.GOOS)

	// Create a new command with the given command and arguments
	command := exec.CommandContext(ctx, cmd, args...)

	log.Infoln("running command: ", command)

	// Execute the command and capture the combined output (stdout and stderr)
	out, err := command.CombinedOutput()
	if err != nil {
		return "", errors.New(string(out) + err.Error())
	}

	log.Infoln("command output:", string(out))

	// Return the command output as a string
	return string(out), nil
}
