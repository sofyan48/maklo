package ssm

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/sofyan48/ssm_go/utils/aws"
	"github.com/sofyan48/ssm_go/utils/tool"
)

// GeneralParametersByPath ...
func GeneralParametersByPath(appname, stage, path string, decryption bool) {
	svc := aws.GetSSM()
	input := &ssm.GetParametersByPathInput{}
	input.SetPath(path)
	input.SetWithDecryption(decryption)
	data, err := svc.GetParametersByPath(input)
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(0)
	}
	// fmt.Println("Data: ", data.String())
	file := tool.Storage(stage, appname)
	for _, i := range data.Parameters {
		name := tool.GeneralSplit(*i.Name)
		format := ""
		if decryption {
			format = name + "=$(aws ssm get-parameter --name " + *i.Name + " --query  \"Parameter.{Value:Value}\" --with-decryption | grep Value | awk -F '\"' '{print $4}')\n"
		} else {
			format = name + "=$(aws ssm get-parameter --name " + *i.Name + " --query  \"Parameter.{Value:Value}\"| grep Value | awk -F '\"' '{print $4}')\n"

		}
		file.Write([]byte(format))
	}
}
