package cli

import (
	"github.com/sofyan48/maklo/libs/cmd"
	"github.com/urfave/cli"
)

// Generate ...
func Generate() cli.Command {
	command := cli.Command{}
	command.Name = "generate"
	command.Usage = "generate [option]"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "Templates Path",
			Destination: &Args.Path,
			Required:    true,
		},
		cli.StringFlag{
			Name:        "format, f",
			Usage:       "Templates Formats | yaml or json",
			Destination: &Args.Type,
		},
	}
	command.Action = func(c *cli.Context) error {
		initEnvirontment()
		cmdHandler := cmd.CMDLibraryHandler()
		return cmdHandler.GenerateByTemplates(Args.Path, Args.Type)
	}
	return command
}
