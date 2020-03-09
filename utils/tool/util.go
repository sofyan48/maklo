package tool

import (
	"log"
	"os"
	"strings"
)

func Storage(stage string) *os.File {
	fileName := ".env-" + stage
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
	return data[5]
}
