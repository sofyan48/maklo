package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sofyan48/ssm_go/cli"
)

func main() {
	godotenv.Load()

	app := cli.AppCommands()
	app.Run(os.Args)
}
