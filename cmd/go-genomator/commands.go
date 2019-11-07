package main

import (
	"fmt"

	"github.com/shizonic/go-genomator/internal"
	"github.com/urfave/cli"
)

var commands = []cli.Command{
	{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "List available recipes",
		Action: func(c *cli.Context) error {
			return nil
		},
	},
	{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "List available recipes",
		Action: func(c *cli.Context) error {
			fmt.Println(internal.RecipesDir)
			return nil
		},
	},
}
