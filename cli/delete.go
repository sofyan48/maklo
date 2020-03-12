package cli

import (
	"github.com/sofyan48/maklo/libs/cmd"
	"github.com/urfave/cli"
)

// Deleted ...
func Deleted() cli.Command {
	command := cli.Command{}
	command.Name = "delete"
	command.Usage = "delete [option]"
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
		cmdHandler := cmd.CMDLibraryHandler()
		cmdHandler.DeleteParameterByTemplate(Args.Path, Args.Type)
		return nil
	}
	return command
}
