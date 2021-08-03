package monitor

import (
	"errors"
	"os/exec"
)

func Exec(args ...string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("no command passed for monitor.Exec")
	}

	baseCmd := args[0]
	cmdArgs := args[1:]

	cmd := exec.Command(baseCmd, cmdArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
