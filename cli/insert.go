package cli

import (
	"github.com/sofyan48/maklo/utils/ssm"
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
			Destination: &Args.Decryption,
		},
	}
	command.Action = func(c *cli.Context) error {
		return ssm.InsertParametersByPath(Args.TemplatePath)
	}

	return command
}
