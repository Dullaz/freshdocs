/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/model"
	"github.com/dullaz/freshdocs/service"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates current state of documents",
	Long: `Validates current state of documents by checking for annotations that
	have no hash, are stale, or point to files that do not exist anymore.
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}

		parser := service.NewParser(config)
		documents, err := parser.Parse()
		if err != nil {
			log.Fatalf("failed to parse documents: %v", err)
		}

		validator := service.NewValidator(config)
		allResults := &model.ValidateResults{}
		for _, document := range documents {
			results, err := validator.Validate(document)
			if err != nil {
				log.Fatalf("failed to validate document: %v", err)
			}
			allResults.Append(results)
		}

		for _, result := range allResults.StaleResults {
			fmt.Println(result)
		}
		for _, result := range allResults.FileMissingResults {
			fmt.Println(result)
		}
		for _, result := range allResults.InvalidResults {
			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
