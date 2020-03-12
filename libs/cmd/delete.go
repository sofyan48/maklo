package cmd

import "github.com/sofyan48/maklo/entity"

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
