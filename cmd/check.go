package cmd

import (
	"fmt"

	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/processor"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for stale documentation",
	Long:  `Perform a quick check to identify documentation that may be stale.`,
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
