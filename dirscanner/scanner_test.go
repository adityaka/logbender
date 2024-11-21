package dirscanner

import (
	"regexp"
	"strconv"
	"testing"
)

func TestDirScanner(t *testing.T) {
	filterPattern, err := regexp.Compile(`.+\.go`)
	if err != nil {
		log.Fatal(err.Error())
	}
	files, error := GetAllFilesWithFilter("/home/addy/src", *filterPattern)
	if error != nil {
		log.Error("Error enumerating files")
	}
	if len(files) == 0 {
		panic("No files found")
	} else {
		log.Info("we got " + strconv.Itoa((len(files))))

	}
}
