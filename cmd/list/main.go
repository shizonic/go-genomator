package list

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/shizonic/go-genomator/internal/utils"
	"github.com/tucnak/climax"
)

var ListCmd = climax.Command{
	Name:  "list",
	Brief: "List available recipes",
	Help:  "List available recipes",
	Flags: []climax.Flag{
		{
			Name:     "directory",
			Short:    "d",
			Usage:    `--directory=<path-to-directory>`,
			Help:     `Directory to look for the recipe`,
			Variable: true,
		},
	},
	Handle: func(ctx climax.Context) int {
		var (
			directory string
			recipes   []string
			err       error
			ok        bool
		)

		if directory, ok = ctx.Get("directory"); !ok {
			directory = utils.RecipesDirectory()
		}
		directory, err = filepath.Abs(directory)

		recipes, err = utils.RecipesInDirectory(directory)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(recipes)

		return 0
	},
}
