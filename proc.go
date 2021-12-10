package proc

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-waitfor/waitfor"

	"github.com/mitchellh/go-ps"
)

const Scheme = "proc"

type Process struct {
	name string
}

func Use() waitfor.ResourceConfig {
	return waitfor.ResourceConfig{
		Scheme:  []string{Scheme},
		Factory: New,
	}
}

func New(u *url.URL) (waitfor.Resource, error) {
	if u == nil {
		return nil, fmt.Errorf("%q: %w", "url", waitfor.ErrInvalidArgument)
	}

	return &Process{name: u.Host}, nil
}

func (p *Process) Test(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	list, err := ps.Processes()

	if err != nil {
		return err
	}

	var found bool

	for _, proc := range list {
		if proc.Executable() == p.name {
			found = true
			break
		}
	}

	if found {
		return nil
	}

	return fmt.Errorf("process not found: %s", p.name)
}
