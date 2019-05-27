package foo

import "github.com/mitchellh/cli"

type cmd struct{}

func New() *cmd {
	return &cmd{}
}

func (c *cmd) Help() string {
	return "Usage: fooCommand..."
}

func (c *cmd) Run(args []string) int {

	return cli.RunResultHelp
}

func (c *cmd) Synopsis() string {
	return "Command foo makes a demo of foo"
}
