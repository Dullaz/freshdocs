package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the main FreshDocs configuration
type Config struct {
	Version        int                   `yaml:"version"`
	Repositories   map[string]Repository `yaml:"repositories"`
	DocumentGroups []DocumentGroup       `yaml:"documentGroups"`
}

// Repository represents a code repository
type Repository struct {
	Path string `yaml:"path"`
}

// DocumentGroup represents a group of documents to monitor
type DocumentGroup struct {
	Path string `yaml:"path"`
	Ext  string `yaml:"ext"`
}

// Load reads a configuration file from the given path
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// Save writes the configuration to the given path
func (c *Config) Save(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
