package zzz

import "github.com/mitchellh/cli"

type cmd struct{}

func New() *cmd {
	return &cmd{}
}

func (c *cmd) Help() string {
	return "Usage: foozzzCommand..."
}

func (c *cmd) Run(args []string) int {

	return cli.RunResultHelp
}

func (c *cmd) Synopsis() string {
	return "SubCommand foo zzz"
}
