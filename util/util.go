package util

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func RunCommand(ctx context.Context, cmd string, args []string) (string, error) {
	command := &exec.Cmd{}
	fullCmd := cmd
	for _, val := range args {
		fullCmd = fullCmd + " " + val
	}
	fmt.Println("full cmd: ", fullCmd)
	if runtime.GOOS == "windows" {
		fmt.Println("Running command in windowdss")
		command = exec.CommandContext(ctx, "powershell", fullCmd)
	} else if runtime.GOOS == "linux" {
		fmt.Println("Running command in linux")
		command = exec.CommandContext(ctx, "/bin/sh", "-c", fullCmd)
	} else {
		fmt.Println("Running command in :", runtime.GOOS)
		command = exec.CommandContext(ctx, cmd, args...)
	}

	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	fmt.Println("command: ", command)

	if err := command.Run(); err != nil {
		return "", errors.New(err.Error() + ": " + stdErr.String())
	}

	return stdOut.String(), nil
}
