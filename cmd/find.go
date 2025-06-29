package cmd

import (
	"fmt"

	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/processor"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find <code-file>",
	Short: "Find documents linked to a code file",
	Long:  `Find all documentation files that are linked to a specific code file.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := "freshdocs.yaml"

		// Load configuration
		cfg, err := config.Load(configPath)
		if err != nil {
			return fmt.Errorf("failed to load configuration: %w", err)
		}

		// Find documents linked to the code file
		processor := processor.NewProcessor(cfg)
		docs, err := processor.Find(args[0])
		if err != nil {
			return fmt.Errorf("failed to find documents: %w", err)
		}

		// Print results
		for _, doc := range docs {
			fmt.Println(doc)
		}

		return nil
	},
}
