package ssm

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/sofyan48/maklo/entity"
	"github.com/sofyan48/maklo/utils/aws"
	"github.com/sofyan48/maklo/utils/tool"
)

// GenerateByTemplates ...
func GenerateByTemplates(path, types string, decryption bool) error {
	data := &entity.TemplatesModels{}
	if types == "json" {
		data, _ = tool.ParsingJSON(path)
	} else {
		data, _ = tool.ParsingYAML(path)
	}
	fmt.Println(data)
	return nil
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
func InsertParametersByPath(path, types string) error {
	data := &entity.TemplatesModels{}
	if types == "json" {
		data, _ = tool.ParsingJSON(path)
	} else {
		data, _ = tool.ParsingYAML(path)
	}
	InsertParameter(data.Parameters)
	return nil
}

// InsertParameter ...
func InsertParameter(dataJSON []entity.InsertDataModels) {
	svc := aws.GetSSM()
	for _, i := range dataJSON {
		inputFormat := &ssm.PutParameterInput{}
		inputFormat.SetName(i.Path)
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
