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

func RunCommand(ctx context.Context, cmd string, args []string) (string, error) {
	log = config.Logger

	log.Infoln("Running command on platform:", runtime.GOOS)
	command := exec.CommandContext(ctx, cmd, args...)

	log.Infoln("command: ", command)

	out, err := command.CombinedOutput()
	if err != nil {
		return "", errors.New(err.Error())
	}

	log.Infoln("command output:", string(out))

	return string(out), nil
}
