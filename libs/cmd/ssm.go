package cmd

import (
	"github.com/sofyan48/maklo/entity"
	"github.com/sofyan48/maklo/libs/ssmlib"
	"github.com/sofyan48/maklo/libs/tool"
)

// CMDLibrary ...
type CMDLibrary struct {
	SSMLib ssmlib.SSManagerInterface
	Tools  tool.ToolsInterface
}

// CMDLibraryHandler ...
func CMDLibraryHandler() *CMDLibrary {
	return &CMDLibrary{
		SSMLib: ssmlib.SSManagerHandler(),
		Tools:  tool.ToolsHandler(),
	}
}

// CMDLibraryInterface ...
type CMDLibraryInterface interface {
	InsertParametersByPath(path, types string, overwrite bool) error
	GenerateByTemplates(path, types string) error
	DeleteParameterByTemplate(path, types string)
	GenerateParametersByPath(appname, stage, path string, decryption bool) error
	GenerateParametersToEnvirontment(appname, stage, path string, decryption bool) error
}

// InsertParametersByPath ...
func (cmd *CMDLibrary) InsertParametersByPath(path, types string, overwrite bool) error {
	data := &entity.TemplatesModels{}
	if types == "json" {
		data, _ = cmd.Tools.ParsingJSON(path)
	} else {
		data, _ = cmd.Tools.ParsingYAML(path)
	}
	cmd.SSMLib.InsertParameter(data.Parameters, overwrite)
	return nil
}

// GenerateByTemplates ...
func (cmd *CMDLibrary) GenerateByTemplates(path, types string) error {
	data := &entity.TemplatesModels{}
	if types == "json" {
		data, _ = cmd.Tools.ParsingJSON(path)
	} else {
		data, _ = cmd.Tools.ParsingYAML(path)
	}
	file := cmd.Tools.Storage(data.Stage, data.Name)
	for _, i := range data.Parameters {
		cmd.Tools.GenerateAWK(i.Path, i.IsEncrypt, file)
	}
	return nil
}

// DeleteParameterByTemplate ...
func (cmd *CMDLibrary) DeleteParameterByTemplate(path, types string) {
	data := &entity.TemplatesModels{}
	if types == "json" {
		data, _ = cmd.Tools.ParsingJSON(path)
	} else {
		data, _ = cmd.Tools.ParsingYAML(path)
	}
	cmd.SSMLib.DeleteParameter(data.Parameters)
}

// GenerateParametersByPath ...
func (cmd *CMDLibrary) GenerateParametersByPath(appname, stage, path string, decryption bool) error {
	file := cmd.Tools.Storage(stage, appname)
	data, err := cmd.SSMLib.GenerateParameters(path, stage, appname, decryption)
	if err != nil {
		return err
	}
	for _, i := range data {
		cmd.Tools.GenerateAWK(i.Name, decryption, file)
	}
	return nil
}

// GenerateParametersToEnvirontment ...
func (cmd *CMDLibrary) GenerateParametersToEnvirontment(appname, stage, path string, decryption bool) error {
	file := cmd.Tools.Storage(stage, appname)
	data, err := cmd.SSMLib.GenerateParameters(path, stage, appname, decryption)
	if err != nil {
		return err
	}
	for _, i := range data {
		cmd.Tools.GenerateDotEnvirontment(i.Name, file)
	}
	return nil
}
