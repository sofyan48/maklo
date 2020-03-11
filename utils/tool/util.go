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

// Storage ...
func Storage(stage, appname string) *os.File {
	if _, err := os.Stat("./result/" + appname); os.IsNotExist(err) {
		os.MkdirAll("./result/"+appname, 0777)
	}
	fileName := "./result/" + appname + "/.env-" + stage
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Println(err)
		f.Close()
		os.Exit(0)
	}
	return f
}

// GeneralSplit ...
func GeneralSplit(envPath string) string {
	fmt.Println("path: ", envPath)
	data := strings.SplitN(envPath, "/", -1)
	result := data[len(data)-1]
	return result
}

// ParsingYAML ...
func ParsingYAML(path string) (*entity.TemplatesModels, error) {
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
func ParsingJSON(path string) (*entity.TemplatesModels, error) {
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
func CheckFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal(err)
		return false
	}
	return true
}
