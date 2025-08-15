package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dullaz/freshdocs/util"
	"gopkg.in/yaml.v3"
)

type FreshConfig struct {
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

func LoadConfig() (*FreshConfig, error) {
	// default config path is in the root of the git repo
	repoPath, err := util.GetRepoPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get repo path: %w", err)
	}

	configPath := filepath.Join(repoPath, "freshdocs.yaml")
	return load(configPath)
}

// Load reads a configuration file from the given path
func load(path string) (*FreshConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config FreshConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}
