package utility

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

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

// WriteFile .
func WriteFile(filepath string, data []byte) error {
	return ioutil.WriteFile(filepath, data, 0644)
}
