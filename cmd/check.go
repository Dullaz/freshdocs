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

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks if any annotations are affected by staged/unstaged changes",
	Long: `Checks if any annotations are affected by staged/unstaged changes.
This includes annotations that are now invalid as a result of staged/unstaged changes.

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

		checker := service.NewChecker(config)
		results := []model.CheckResult{}
		for _, document := range documents {
			documentResults, err := checker.Check(document)
			if err != nil {
				log.Fatalf("failed to check document: %v", err)
			}
			results = append(results, documentResults...)
		}

		for _, result := range results {
			fmt.Println(result.String())
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
