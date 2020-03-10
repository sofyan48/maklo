package ssm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/sofyan48/ssm_go/utils/aws"
	"github.com/sofyan48/ssm_go/utils/tool"
)

// InsertDataModels ...
type InsertDataModels struct {
	Name      string
	Value     string
	IsEncrypt bool
}

// GeneralParametersByPath ...
func GeneralParametersByPath(appname, stage, path string, decryption bool) error {
	svc := aws.GetSSM()
	input := &ssm.GetParametersByPathInput{}
	input.SetPath(path)
	input.SetWithDecryption(decryption)
	data, err := svc.GetParametersByPath(input)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	// fmt.Println("Data: ", data.String())
	file := tool.Storage(stage, appname)
	for _, i := range data.Parameters {
		GenerateJSON(*i.Name, decryption, file)
	}
	return nil
}

// GenerateJSON ...
func GenerateJSON(pathName string, decryption bool, file *os.File) (int, error) {
	name := tool.GeneralSplit(pathName)
	format := ""
	if decryption {
		format = name + "=$(aws ssm get-parameter --name " + pathName + " --query  \"Parameter.{Value:Value}\" --with-decryption | grep Value | awk -F '\"' '{print $4}')\n"
	} else {
		format = name + "=$(aws ssm get-parameter --name " + pathName + " --query  \"Parameter.{Value:Value}\"| grep Value | awk -F '\"' '{print $4}')\n"

	}
	return file.Write([]byte(format))
}

// InsertParametersByPath ...
func InsertParametersByPath(path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	dataJSON := []InsertDataModels{}
	err = json.Unmarshal(byteValue, &dataJSON)
	if err != nil {
		return err
	}
	InsertParameter(dataJSON)
	return nil
}

// InsertParameter ...
func InsertParameter(dataJSON []InsertDataModels) {
	svc := aws.GetSSM()
	for _, i := range dataJSON {
		inputFormat := &ssm.PutParameterInput{}
		inputFormat.SetName(i.Name)
		inputFormat.SetValue(i.Value)
		if i.IsEncrypt {
			inputFormat.SetType("SecureString")
		} else {
			inputFormat.SetType("String")
		}

		result, err := svc.PutParameter(inputFormat)
		if err != nil {
			log.Println("Not Uploaded Parameter: ", err)
		}
		log.Println("Uploaded Parameter: ", result)
	}
}
