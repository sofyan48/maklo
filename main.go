package main

import (
	"github.com/joho/godotenv"
	"github.com/sofyan48/ssm_go/utils/ssm"
)

func main() {
	godotenv.Load()
	ssm.GeneralParametersByPath("dev", "/rll/dev/secret/sdk_js")
}
