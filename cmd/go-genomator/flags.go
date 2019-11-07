package main

import (
	"github.com/shizonic/go-genomator/internal"
	"github.com/shizonic/go-genomator/internal/helpers"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:        "directory, d",
		Value:       helpers.GetRecipesDir(),
		Usage:       "language for the greeting",
		Destination: &internal.RecipesDir,
	},
}
