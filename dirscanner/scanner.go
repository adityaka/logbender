package dirscanner

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/adityaka/logbender/logging"
)

var log *logging.Log = logging.GetLogger(logging.LoggingConfig{Name: "dir_scanner", FileFullPath: "logbender_logscanner.log"})

func GetAllFilesWithFilter(rootDir string, filter regexp.Regexp) ([]string, error) {
	files := make([]string, 0)
	err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		currentPath := filepath.Join(path, d.Name())
		log.Info("Visiting " + filepath.Join(rootDir, d.Name()))
		if !d.IsDir() {
			if filter.MatchString(currentPath) {
				files = append(files, currentPath)
			}
		}
		return nil
	})
	return files, err
}
