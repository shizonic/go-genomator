package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var commands = []cli.Command{
	{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a task to the list",
		Action: func(c *cli.Context) error {
			fmt.Println("added task: ", c.Args().First())
			return nil
		},
	},
}
