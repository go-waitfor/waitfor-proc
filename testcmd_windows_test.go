//go:build windows

package proc_test

import (
	"errors"
	"os/exec"
)

type TestCommand struct {
	cmd *exec.Cmd
}

func NewTestCommand() (*TestCommand, error) {
	// Use timeout command on Windows as it's always available and waits for specified time
	// timeout /t 10 waits for 10 seconds but doesn't accept input
	cmd := exec.Command("timeout", "/t", "10", "/nobreak")

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	if cmd.Process == nil {
		return nil, errors.New("failed to start the test process")
	}

	return &TestCommand{cmd}, nil
}

func (c *TestCommand) Name() string {
	// On Windows, timeout appears as "timeout.exe" in process list
	return "timeout.exe"
}

func (c *TestCommand) Kill() {
	_ = c.cmd.Process.Kill()
}