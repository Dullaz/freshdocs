package cmd

import (
	"fmt"

	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/processor"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Show which docs are stale (out of date with committed code)",
	Long:  `Check which documentation files are stale, meaning their linked code has changed in committed history since the last update. Use this to find docs that need updating after code changes have been committed.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := "freshdocs.yaml"

		// Load configuration
		cfg, err := config.Load(configPath)
		if err != nil {
			return fmt.Errorf("failed to load configuration: %w", err)
		}

		// Process documents and find affected ones
		processor := processor.NewProcessor(cfg)

		// Check for missing files first
		missing, err := processor.CheckMissingFiles()
		if err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}

		// Print missing file errors
		for _, doc := range missing {
			fmt.Printf("%s references %s that doesn't exist\n", doc.Path, doc.AffectedBy)
		}

		// Check for stale documents
		affected, err := processor.Validate()
		if err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}

		// Print stale document results
		for _, doc := range affected {
			fmt.Printf("%s affected by %s\n", doc.Path, doc.AffectedBy)
		}

		return nil
	},
}
