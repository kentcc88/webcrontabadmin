package file

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
