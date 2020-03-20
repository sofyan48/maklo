package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvironment ...
func (cmd *CMDLibrary) LoadEnvironment(path string) {
	godotenv.Load(path)
}

// CreateEnvironment ...
func (cmd *CMDLibrary) CreateEnvironment() {
	homePath, _ := os.UserHomeDir()
	err := cmd.Tools.CreateFolder(homePath + "/.maklo/")
	if err == nil {
		fileName := homePath + "/.maklo/environtment"
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Println(err)
			file.Close()
			os.Exit(0)
		}
		log.Println("Environment Created")
		return
	}
	log.Println("Environment Not Found")
}
