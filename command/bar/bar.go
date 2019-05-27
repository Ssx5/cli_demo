package bar

import (
	"fmt"
)

type cmd struct{}

func New() *cmd {
	return &cmd{}
}

func (c *cmd) Help() string {
	return "Usage: barCommand"
}

func (c *cmd) Run(args []string) int {
	if len(args) > 0 {
		fmt.Print("Args:")
		for _, v := range args {
			fmt.Print(" ", v)
		}
		fmt.Println()
	}
	fmt.Println("Run barCommand")
	return 0
}

func (c *cmd) Synopsis() string {
	return "This is command barCommand"
}
