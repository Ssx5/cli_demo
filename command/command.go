package command

import (
	"./bar"
	"./foo"
	"./foo/zzz"
	"github.com/mitchellh/cli"
)

var commands map[string]cli.CommandFactory

func init() {
	Register("foo", func(ui cli.Ui) (cli.Command, error) { return foo.New(), nil })
	Register("foo zzz", func(ui cli.Ui) (cli.Command, error) { return zzz.New(), nil })
	Register("bar", func(ui cli.Ui) (cli.Command, error) { return bar.New(), nil })
}
