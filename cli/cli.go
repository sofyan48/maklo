package cli

import (
	"github.com/urfave/cli"
)

// ArgsMapping ...
type ArgsMapping struct {
	Name         string
	TemplatePath string
	Path         string
	Stage        string
	Type         string
	Decryption   string
}

// Args ...
var Args ArgsMapping

// Init Initialise a CLI app
func Init() *cli.App {
	app := cli.NewApp()
	app.Name = "maklo"
	app.Usage = "maklo [command]"
	app.Author = "sofyan48"
	app.Email = "meongbego@gmail.com"
	app.Version = "0.0.1"
	return app
}

// AppCommands All Command line app
func AppCommands() *cli.App {
	app := Init()
	app.Commands = []cli.Command{
		Generate(),
		InsertParameter(),
	}
	return app
}
