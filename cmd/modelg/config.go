package main

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Models []*ModelConfig `yaml:"models"`

	Converters []struct {
		Name    string `yaml:"name"`
		Scanner string `yaml:"scanner"`
		Valuer  string `yaml:"valuer"`
	} `yaml:"converters"`
}

type ModelConfig struct {
	Name                 string `yaml:"name"`
	Queries              string `yaml:"queries"`
	SkipCreateParameters bool   `yaml:"skip_create_parameters"`
	SkipUpdateParameters bool   `yaml:"skip_update_parameters"`
	TableName            string `yaml:"table_name"`
	DisplayName          string `yaml:"display_name"`
	PluralDisplayName    string `yaml:"plural_display_name"`
}

func LoadConfig(r io.Reader) (*Config, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
