package proc_test

import (
	"context"
	"errors"
	"net/url"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-proc"
)

type TestCommand struct {
	cmd *exec.Cmd
}

func NewTestCommand() (*TestCommand, error) {
	cmd := exec.Command("man", "cat")

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	if cmd.Process == nil {
		return nil, errors.New("failed to start the test process")
	}

	return &TestCommand{cmd}, nil
}

func (c *TestCommand) Name() string {
	return "man"
}

func (c *TestCommand) Kill() {
	_ = c.cmd.Process.Kill()
}

func TestProcess_Use(t *testing.T) {
	cmd, err := NewTestCommand()

	assert.NoError(t, err)

	defer cmd.Kill()

	uStr := "proc://" + cmd.Name()
	w := waitfor.New(proc.Use())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = w.Test(ctx, []string{uStr})

	assert.NoError(t, err)
}

func TestProcess_Test(t *testing.T) {
	cmd, err := NewTestCommand()

	assert.NoError(t, err)

	defer cmd.Kill()

	procUrl, err := url.Parse("proc://" + cmd.Name())

	assert.NoError(t, err)

	r, err := proc.New(procUrl)

	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = r.Test(ctx)

	assert.NoError(t, err)
}

func TestProcess_Test_Fail(t *testing.T) {
	procUrl, err := url.Parse("proc://foo")

	assert.NoError(t, err)

	r, err := proc.New(procUrl)

	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = r.Test(ctx)

	assert.Error(t, err)
}
