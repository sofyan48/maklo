package cli

import (
	"github.com/sofyan48/maklo/utils/ssm"
	"github.com/urfave/cli"
)

// Generate ...
func GeneratePath() cli.Command {
	command := cli.Command{}
	command.Name = "generatePath"
	command.Usage = "generatePath [option]"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "Path Parameters Aws System Manager",
			Destination: &Args.Path,
		},
		cli.StringFlag{
			Name:        "name, n",
			Usage:       "App Name",
			Destination: &Args.Name,
			Required:    true,
		},
		cli.StringFlag{
			Name:        "stage, s",
			Usage:       "Stage Parameters",
			Destination: &Args.Stage,
			Required:    true,
		},
		cli.StringFlag{
			Name:        "decrypt, d",
			Usage:       "Decryption Option",
			Destination: &Args.Decryption,
			Required:    true,
		},
	}
	command.Action = func(c *cli.Context) error {
		decrypt := false
		if Args.Decryption != "" {
			decrypt = true
		}
		return ssm.GeneralParametersByPath(Args.Name, Args.Stage, Args.Path, decrypt)
	}

	return command
}
