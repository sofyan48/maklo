package main

import (
	"os"

	"github.com/sofyan48/maklo/cli"
)

func main() {
	app := cli.AppCommands()
	app.Run(os.Args)
}
