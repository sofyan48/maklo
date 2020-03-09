package tool

import (
	"log"
	"os"
	"strings"
)

// Storage ...
func Storage(stage, appname string) *os.File {
	if _, err := os.Stat("./result/" + appname); os.IsNotExist(err) {
		os.MkdirAll("./result/"+appname, 0777)
	}
	fileName := "./result/" + appname + "/.env-" + stage
	f, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		f.Close()
		os.Exit(0)
	}
	return f
}

// GeneralSplit ...
func GeneralSplit(envPath string) string {
	data := strings.SplitN(envPath, "/", -1)
	result := data[len(data)-1]
	return result
}
