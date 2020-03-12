package cmd

import "github.com/sofyan48/maklo/entity"

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
