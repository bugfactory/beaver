package utils

import (
	"path/filepath"
	"io/ioutil"
)

func ReadFile(fileName string) ([]byte, error) {
	filePath, err := filepath.Abs(fileName)
	if err != nil {
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

    return yamlFile, nil
}
