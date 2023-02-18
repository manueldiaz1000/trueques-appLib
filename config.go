package truequeslib

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	GinMode   string    `yaml:"gin_mode"`
	Server    Server    `yaml:"server"`
	Endpoints Endpoints `yaml:"endpoints"`
}

type Server struct {
	Port         int `yaml:"port"`
	ReadTimeout  int `yaml:"read_timeout"`
	WriteTimeout int `yaml:"write_timeout"`
}

type Endpoints struct {
	Adverts string `yaml:"adverts"`
}

func LoadConfig(cfgFileName string) (Config, error) {
	var config Config

	yamlFile, err := os.ReadFile(cfgFileName)
	if err != nil {
		return config, fmt.Errorf("err-reading_yamlFile %s: %w", cfgFileName, err)
	}

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		return config, fmt.Errorf("err-unmarshalling_yamlFile %s: %w", cfgFileName, err)
	}

	return config, nil
}
