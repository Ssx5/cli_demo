package main

import (
	"log"
	"os"

	"./command"
	"github.com/mitchellh/cli"
)

func main() {

	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	c := &cli.CLI{
		Name:     "app",
		Version:  "1.0.0",
		Args:     os.Args,
		Commands: command.Map(ui),
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
