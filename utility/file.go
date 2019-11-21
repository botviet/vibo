package utility

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// IsExist .
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ReadFile .
func ReadFile(file string) (content []byte, err error) {
	content = make([]byte, 0)
	content, err = ioutil.ReadFile(file)

	return
}

// ReadFiles from diretory
func ReadFiles(dir string) (content map[string][]byte, err error) {
	content = make(map[string][]byte, 0)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		b, err := ReadFile(path)
		if err != nil {
			return err
		}

		content[path] = b
		return nil
	})

	return
}

// WalkFileName .
func WalkFileName(dir string) (fileNames []string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileNames = append(fileNames, filepath.Base(path))
		return nil
	})

	return
}

// BasePath .
func BasePath(path string) string {
	return filepath.Base(path)
}

// WriteAppend text to exits file
func WriteAppend(text string, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Error(err)
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// WriteFile .
func WriteFile(filepath string, data []byte) error {
	return ioutil.WriteFile(filepath, data, 0644)
}

// CreateDir .
func CreateDir(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Error(err)
		}
	} else if err != nil {
		log.Error(err)
	}
}
