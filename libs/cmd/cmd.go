package cmd

import (
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
	CreateEnvironment()
	LoadEnvironment(path string)
}
