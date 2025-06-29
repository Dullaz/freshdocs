package cmd

import (
	"fmt"
	"os"

	"github.com/dullaz/freshdocs/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new FreshDocs configuration",
	Long:  `Create a new freshdocs.yaml configuration file in the current directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := "freshdocs.yaml"

		// Check if config already exists
		if _, err := os.Stat(configPath); err == nil {
			return fmt.Errorf("configuration file %s already exists", configPath)
		}

		// Create default configuration
		defaultConfig := config.Config{
			Version: 1,
			Repositories: map[string]config.Repository{
				"core": {
					Path: "../my-service",
				},
				"utils": {
					Path: "../shared-utils",
				},
			},
			DocumentGroups: []config.DocumentGroup{
				{
					Path: "./docs/folder",
					Ext:  ".md",
				},
			},
		}

		// Write configuration to file
		if err := defaultConfig.Save(configPath); err != nil {
			return fmt.Errorf("failed to create configuration file: %w", err)
		}

		fmt.Printf("FreshDocs configuration created at %s\n", configPath)
		fmt.Println("Please update the configuration with your actual repository paths and document folders.")

		return nil
	},
}
