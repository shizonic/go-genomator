package utils

import (
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/shizonic/go-genomator/versioninfo"
)

func DirectoryExist(dir string) bool {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return true
	}
	return false
}

func RecipesDirectory() string {
	var (
		filePath string
		name     string = strings.ToLower(versioninfo.ProjectName)
	)

	if len(os.Getenv("XDG_CONFIG_HOME")) > 0 {
		filePath, _ = filepath.Abs(path.Join(os.Getenv("XDG_CONFIG_HOME"), name))
	}

	usr, _ := user.Current()
	filePath, _ = filepath.Abs(path.Join(usr.HomeDir, ".config", name))

	return filePath
}

func RecipesInDirectory(directory string) ([]string, error) {
	var recipes []string

	err := filepath.Walk(directory,
		func(directory string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filepath.Ext(directory) == ".json" {
				recipes = append(recipes, directory)
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	return recipes, nil
}
