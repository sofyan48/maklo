package ssm

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/sofyan48/ssm_go/utils/aws"
	"github.com/sofyan48/ssm_go/utils/tool"
)

// GeneralParametersByPath ...
func GeneralParametersByPath(stage, path string) {
	svc := aws.GetSSM()
	input := &ssm.GetParametersByPathInput{}
	input.SetPath(path)
	data, err := svc.GetParametersByPath(input)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	file := tool.Storage(stage)
	for _, i := range data.Parameters {
		name := tool.GeneralSplit(*i.Name)
		format := name + "=$(aws ssm get-parameter --name " + *i.Name + " --query  \"Parameter.{Value:Value}\"| grep Value | awk -F '\"' '{print $4}')\n"
		file.Write([]byte(format))
	}
}
