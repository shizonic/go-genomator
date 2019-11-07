package run

import (
	"fmt"
	"log"

	"github.com/shizonic/go-genomator/internal/helpers"
	"github.com/tucnak/climax"
)

var RunCmd = climax.Command{
	Name:  "run",
	Brief: "Run a recipe",
	Help:  "Run a recipe",
	Flags: []climax.Flag{
		{
			Name:     "recipe",
			Short:    "r",
			Usage:    `--recipe=<name>`,
			Help:     `The name of the recipe to run`,
			Variable: true,
		},
		{
			Name:     "directory",
			Short:    "d",
			Usage:    `--directory=<path-to-directory>`,
			Help:     `Directory to look for the recipe`,
			Variable: true,
		},
	},
	Handle: func(ctx climax.Context) int {
		var recipe, directory string

		if r, ok := ctx.Get("recipe"); ok {
			recipe = r
		} else {
			log.Fatal("Name of the recipe is required")
		}

		if _, ok := ctx.Get("directory"); !ok {
			directory = helpers.RecipesDirectory()
		}

		fmt.Println(recipe, directory)

		return 0
	},
}
