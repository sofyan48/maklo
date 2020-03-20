package cli

import (
	"os"

	"github.com/sofyan48/maklo/libs/cmd"
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
	OverWrites   string
	Environment  string
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
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "environtment, e",
			Usage:       "Load Environtment from path",
			Destination: &Args.Environment,
		},
	}

	return app
}

// AppCommands All Command line app
func AppCommands() *cli.App {
	app := Init()
	app.Commands = []cli.Command{
		Generate(),
		InsertParameter(),
		GeneratePath(),
		Deleted(),
	}
	return app
}

func initEnvirontment() {
	cmdHandler := cmd.CMDLibraryHandler()
	if Args.Environment == "" {
		cmdHandler.CreateEnvironment()
		homePath, _ := os.UserHomeDir()
		cmdHandler.LoadEnvironment(homePath + "/.maklo/environtment")
		return
	}
	cmdHandler.LoadEnvironment(Args.Environment)
	return
}
