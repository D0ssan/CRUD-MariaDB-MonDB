package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

func New(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			os.Exit(1)
		}
	}(file)

	cfg := Config{}
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
