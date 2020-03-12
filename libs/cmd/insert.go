package cmd

import "github.com/sofyan48/maklo/entity"

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
