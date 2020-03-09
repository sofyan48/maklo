package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sofyan48/ssm_go/utils/ssm"
	"github.com/urfave/cli"
)

type ArgsMapping struct {
	Name  string
	Stage string
	Type  string
}

// Args ...
var Args ArgsMapping

func main() {
	godotenv.Load()

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name, n",
				Usage:       "--name | -n ",
				Destination: &Args.Name,
			},
			&cli.StringFlag{
				Name:        "stage, s",
				Usage:       "--stage | -s ",
				Destination: &Args.Stage,
			},
			&cli.StringFlag{
				Name:        "type, t",
				Usage:       "--type | -t ",
				Destination: &Args.Type,
			},
		},
		Action: func(c *cli.Context) error {
			if Args.Name == "" || Args.Stage == "" || Args.Type == "" {
				fmt.Println("Setting Name, Stage AND Type")
				os.Exit(0)
			}
			ssm.GeneralParametersByPath(Args.Name, Args.Stage, "/rll/"+Args.Stage+"/"+Args.Type+"/"+Args.Name)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
