package configuration

import (
    "gopkg.in/yaml.v2"
    "github.com/bugfactory/beaver/utils"
)

type Configuration struct {
    Ip string
    Port int
}

func New (filename string) (Configuration, error) {
	var conf Configuration
	yamlFile, err := utils.ReadFile(filename)
	if err != nil {
		return Configuration{}, err
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return Configuration{}, err
	}

	return conf, nil
}
