package cmdify

import (
	"context"
	"fmt"
	"github.com/amirimatin/cmdify/command"
	"os"
)

type CmdRunner struct {
	name   string
	args   []string
	ctx    context.Context
	useTTY bool
}

func Run(name string, args ...string) *CmdRunner {
	return &CmdRunner{
		name: name,
		args: args,
		ctx:  context.Background(),
	}
}

func (c *CmdRunner) WithContext(ctx context.Context) *CmdRunner {
	c.ctx = ctx
	return c
}

func (c *CmdRunner) WithTTY() *CmdRunner {
	c.useTTY = true
	return c
}

func (c *CmdRunner) Must() {
	if c.useTTY {
		err := command.ExecuteWithTTY(c.name, c.args...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		res := command.Execute(c.ctx, c.name, c.args...)
		fmt.Print(res.Output)
		if res.ExitCode != 0 {
			os.Exit(res.ExitCode)
		}
	}
}
