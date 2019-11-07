package main

import (
	"github.com/shizonic/go-genomator"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:        "directory, d",
		Value:       genomator.GetRecipesDir(),
		Usage:       "language for the greeting",
		Destination: &genomator.RecipesDir,
	},
}
