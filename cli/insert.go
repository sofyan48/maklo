package cli

import (
	"github.com/sofyan48/maklo/libs/cmd"
	"github.com/urfave/cli"
)

// InsertParameter ..
func InsertParameter() cli.Command {
	command := cli.Command{}
	command.Name = "insert"
	command.Usage = "insert [Option]"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "File Template Path",
			Destination: &Args.TemplatePath,
			Required:    true,
		},
		cli.StringFlag{
			Name:        "overwrite, w",
			Usage:       "Overwirte Option",
			Destination: &Args.OverWrites,
		},
		cli.StringFlag{
			Name:        "format, f",
			Usage:       "Format Type Templates Option",
			Destination: &Args.Type,
		},
	}
	command.Action = func(c *cli.Context) error {
		initEnvirontment()
		overwrites := false
		if Args.OverWrites != "" {
			overwrites = true
		}
		cmdHandler := cmd.CMDLibraryHandler()
		return cmdHandler.InsertParametersByPath(Args.TemplatePath, Args.Type, overwrites)
	}

	return command
}
