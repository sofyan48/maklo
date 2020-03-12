package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/sofyan48/maklo/entity"
	"gopkg.in/yaml.v2"
)

// Tools ...
type Tools struct{}

// ToolsHandler ...
func ToolsHandler() *Tools {
	return &Tools{}
}

// ToolsInterface ...
type ToolsInterface interface {
	Storage(stage, appname string) *os.File
	GeneralSplit(envPath string) string
	ParsingYAML(path string) (*entity.TemplatesModels, error)
	ParsingJSON(path string) (*entity.TemplatesModels, error)
	CheckFile(path string) bool
	GenerateAWK(pathName string, decryption bool, file *os.File) (int, error)
	GenerateDotEnvirontment(pathName string, file *os.File) (int, error)
	CreateFolder(path string) error
}

// GenerateAWK ...
func (tl *Tools) GenerateAWK(pathName string, decryption bool, file *os.File) (int, error) {
	name := tl.GeneralSplit(pathName)
	format := ""
	if decryption {
		format = name + "=$(aws ssm get-parameter --name " + pathName + " --query  \"Parameter.{Value:Value}\" --with-decryption | grep Value | awk -F '\"' '{print $4}')\n"
	} else {
		format = name + "=$(aws ssm get-parameter --name " + pathName + " --query  \"Parameter.{Value:Value}\"| grep Value | awk -F '\"' '{print $4}')\n"

	}
	return file.Write([]byte(format))
}

// GenerateDotEnvirontment ...
func (tl *Tools) GenerateDotEnvirontment(pathName string, file *os.File) (int, error) {
	name := tl.GeneralSplit(pathName)
	format := name + "=" + pathName + "\n"
	return file.Write([]byte(format))
}

// Storage ...
func (tl *Tools) Storage(stage, appname string) *os.File {
	if _, err := os.Stat("./result/" + appname); os.IsNotExist(err) {
		os.MkdirAll("./result/"+appname, 0777)
	}
	fileName := "./result/" + appname + "/.env-" + stage
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Println(err)
		file.Close()
		os.Exit(0)
	}
	return file
}

// CreateFolder ...
func (tl *Tools) CreateFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	return os.MkdirAll(path, 0777)
}

// GeneralSplit ...
func (tl *Tools) GeneralSplit(envPath string) string {
	fmt.Println("path: ", envPath)
	data := strings.SplitN(envPath, "/", -1)
	result := data[len(data)-1]
	return result
}

// ParsingYAML ...
func (tl *Tools) ParsingYAML(path string) (*entity.TemplatesModels, error) {
	yamlObject := &entity.TemplatesModels{}
	ymlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return yamlObject, err
	}
	err = yaml.Unmarshal(ymlFile, yamlObject)
	if err != nil {
		return yamlObject, err
	}
	return yamlObject, nil
}

// ParsingJSON ...
func (tl *Tools) ParsingJSON(path string) (*entity.TemplatesModels, error) {
	jsonObject := &entity.TemplatesModels{}
	ymlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return jsonObject, err
	}
	err = json.Unmarshal(ymlFile, jsonObject)
	if err != nil {
		return jsonObject, err
	}
	return jsonObject, nil
}

// CheckFile ...
func (tl *Tools) CheckFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal(err)
		return false
	}
	return true
}
