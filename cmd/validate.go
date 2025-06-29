package cmd

import (
	"fmt"

	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/processor"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate which documents are affected by code changes",
	Long:  `Check which documentation files are affected by changes in the linked code files.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := "freshdocs.yaml"

		// Load configuration
		cfg, err := config.Load(configPath)
		if err != nil {
			return fmt.Errorf("failed to load configuration: %w", err)
		}

		// Process documents and find affected ones
		processor := processor.NewProcessor(cfg)
		affected, err := processor.Validate()
		if err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}

		// Print results
		for _, doc := range affected {
			fmt.Printf("%s affected by %s\n", doc.Path, doc.AffectedBy)
		}

		return nil
	},
}
