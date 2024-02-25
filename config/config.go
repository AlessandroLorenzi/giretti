package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Title string `yaml:"title"`
}

func ParseConfig(filename string) (*Config, error) {
	// Read the YAML file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config data: %v", err)
		return nil, err
	}

	return &config, nil
}
