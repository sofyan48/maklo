package ssmlib

import (
	"log"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/sofyan48/maklo/entity"
	"github.com/sofyan48/maklo/libs/aws"
	"github.com/sofyan48/maklo/libs/tool"
)

// SSManager ...
type SSManager struct {
	AWS   aws.AWSLibraryInterface
	Tools tool.ToolsInterface
}

// SSManagerHandler ...
func SSManagerHandler() *SSManager {
	return &SSManager{
		AWS:   aws.AWSLibraryHandler(),
		Tools: tool.ToolsHandler(),
	}
}

// SSManagerInterface ...
type SSManagerInterface interface {
	InsertParameter(dataJSON []entity.InsertDataModels, overwrite bool)
	DeleteParameter(dataJSON []entity.InsertDataModels)
	GenerateParameters(path, stage, appname string, decryption bool) ([]entity.GenerateOutputs, error)
}

// GenerateParameters ...
func (ssmtool *SSManager) GenerateParameters(path, stage, appname string, decryption bool) ([]entity.GenerateOutputs, error) {
	svc := ssmtool.AWS.GetSystemManagerStore()
	results := []entity.GenerateOutputs{}
	input := &ssm.GetParametersByPathInput{}
	input.SetPath(path)
	input.SetWithDecryption(decryption)
	data, err := svc.GetParametersByPath(input)
	if err != nil {
		return results, err
	}
	for _, i := range data.Parameters {
		dataTemps := entity.GenerateOutputs{}
		dataTemps.Name = *i.Name
		dataTemps.Value = *i.Value
		dataTemps.Version = *i.Version
		dataTemps.Type = *i.Type
		dataTemps.LastModified = *&i.LastModifiedDate
		results = append(results, dataTemps)
	}
	return results, nil
}

// InsertParameter ...
func (ssmtool *SSManager) InsertParameter(dataJSON []entity.InsertDataModels, overwrite bool) {
	svc := ssmtool.AWS.GetSystemManagerStore()
	for _, i := range dataJSON {
		inputFormat := &ssm.PutParameterInput{}
		inputFormat.SetName(i.Path)
		inputFormat.SetValue(i.Value)
		inputFormat.SetOverwrite(overwrite)
		if i.IsEncrypt {
			inputFormat.SetType("SecureString")
		} else {
			inputFormat.SetType("String")
		}
		result, err := svc.PutParameter(inputFormat)
		if err != nil {
			log.Println("Not Uploaded Parameter: ", err)
			return
		}
		log.Println("Uploaded Parameter: ", result)
	}
}

// DeleteParameter ..
func (ssmtool *SSManager) DeleteParameter(dataJSON []entity.InsertDataModels) {
	svc := ssmtool.AWS.GetSystemManagerStore()
	for _, i := range dataJSON {
		inputFormat := &ssm.DeleteParameterInput{}
		inputFormat.SetName(i.Path)
		result, err := svc.DeleteParameter(inputFormat)
		if err != nil {
			log.Println("Not Deleted Parameter: ", err)
		}
		log.Println("Deleted Parameter: ", result)
	}

}
