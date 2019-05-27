package command

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type Factory func(ui cli.Ui) (cli.Command, error)

func Register(name string, fn Factory) {
	if registry == nil {
		registry = make(map[string]Factory)
	}
	if registry[name] != nil {
		panic(fmt.Errorf("Command %q is already registered", name))
	}
	registry[name] = fn
}

var registry map[string]Factory

func Map(ui cli.Ui) map[string]cli.CommandFactory {
	m := make(map[string]cli.CommandFactory)
	for name, fn := range registry {
		thisFn := fn
		m[name] = func() (cli.Command, error) {
			return thisFn(ui)
		}
	}
	return m
}
