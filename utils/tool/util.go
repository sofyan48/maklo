package tool

import (
	"fmt"
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
