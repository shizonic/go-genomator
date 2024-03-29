package main

import (
	"github.com/shizonic/go-genomator/cmd/list"
	"github.com/shizonic/go-genomator/cmd/run"
	"github.com/shizonic/go-genomator/versioninfo"
	"github.com/tucnak/climax"
)

func main() {
	app := climax.New(versioninfo.ProjectName)
	app.Name = versioninfo.ProjectName
	app.Brief = versioninfo.ProjectDescription
	app.Version = versioninfo.Revision

	app.AddCommand(run.RunCmd)
	app.AddCommand(list.ListCmd)

	app.Run()
}
