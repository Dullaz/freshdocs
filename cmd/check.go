package cmd

import (
	"fmt"

	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/processor"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for docs affected by uncommitted changes",
	Long:  `Scan for documentation files that are affected by uncommitted (staged or unstaged) changes in your code. This helps you see what docs will need updating before you commit.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := "freshdocs.yaml"

		// Load configuration
		cfg, err := config.Load(configPath)
		if err != nil {
			return fmt.Errorf("failed to load configuration: %w", err)
		}

		// Process documents and find stale ones
		processor := processor.NewProcessor(cfg)
		stale, err := processor.Check()
		if err != nil {
			return fmt.Errorf("check failed: %w", err)
		}

		// Print results
		for _, doc := range stale {
			fmt.Printf("%s affected by %s\n", doc.Path, doc.AffectedBy)
		}

		return nil
	},
}
