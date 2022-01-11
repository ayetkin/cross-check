package cfg

import (
	"cross-check/model"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var Values = model.Config{}

func NewConfig(configPath string) error {
	if err := ValidateConfigPath(configPath); err != nil {
		return err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&Values); err != nil {
		return err
	}
	return nil
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}