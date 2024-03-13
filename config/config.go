package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var Config *ConfigStruct

type ConfigStruct struct {
	Title   string `yaml:"title"`
	BaseUrl string `yaml:"base_url"`
	Author  string `yaml:"author"`
}

func Init(filename string) error {
	var err error
	Config, err = ParseConfig(filename)
	return err
}

func ParseConfig(filename string) (*ConfigStruct, error) {
	// Read the YAML file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil, err
	}

	var config ConfigStruct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config data: %v", err)
		return nil, err
	}

	return &config, nil
}
