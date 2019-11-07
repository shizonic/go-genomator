package helpers

import (
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/shizonic/go-genomator/internal"
	"github.com/shizonic/go-genomator/versioninfo"
)

func DirExist(dir string) bool {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return true
	}
	return false
}

func GetRecipesDir() string {
	name := strings.ToLower(versioninfo.ProjectName)
	if len(os.Getenv("XDG_CONFIG_HOME")) > 0 {
		return path.Join(os.Getenv("XDG_CONFIG_HOME"), name)
	}

	usr, _ := user.Current()
	return path.Join(usr.HomeDir, name)
}

func PrepareEnviron() {
	if !DirExist(internal.RecipesDir) {
		os.MkdirAll(internal.RecipesDir, 0644)
	}
}
