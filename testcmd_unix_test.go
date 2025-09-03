//go:build !windows

package proc_test

import (
	"errors"
	"os/exec"
)

type TestCommand struct {
	cmd *exec.Cmd
}

func NewTestCommand() (*TestCommand, error) {
	// Use cross-platform compatible sleep command on Unix-like systems
	cmd := exec.Command("sleep", "10")

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	if cmd.Process == nil {
		return nil, errors.New("failed to start the test process")
	}

	return &TestCommand{cmd}, nil
}

func (c *TestCommand) Name() string {
	// On Unix-like systems, sleep appears as "sleep" in process list
	return "sleep"
}

func (c *TestCommand) Kill() {
	_ = c.cmd.Process.Kill()
}