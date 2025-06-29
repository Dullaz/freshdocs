package cmd

import (
	"fmt"

	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/processor"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [file]",
	Short: "Update document hashes",
	Long:  `Update the hashes for all documents or a specific document file.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := "freshdocs.yaml"

		// Load configuration
		cfg, err := config.Load(configPath)
		if err != nil {
			return fmt.Errorf("failed to load configuration: %w", err)
		}

		// Process documents
		processor := processor.NewProcessor(cfg)

		if len(args) > 0 {
			// Update specific file
			if err := processor.UpdateFile(args[0]); err != nil {
				return fmt.Errorf("failed to update file %s: %w", args[0], err)
			}
			fmt.Printf("Updated hashes for %s\n", args[0])
		} else {
			// Update all files
			if err := processor.Update(); err != nil {
				return fmt.Errorf("failed to update hashes: %w", err)
			}
			fmt.Println("Updated hashes for all documents")
		}

		return nil
	},
}
