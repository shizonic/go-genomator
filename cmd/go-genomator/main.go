package main

import (
	"log"
	"os"

	"github.com/shizonic/go-genomator/versioninfo"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = versioninfo.Version
	app.Name = versioninfo.ProjectName
	app.Usage = versioninfo.ProjectDescription
	app.Flags = flags
	app.Commands = commands

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
