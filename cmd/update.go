/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/service"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update all annotations in all documents",
	Long: `Update all annotations in all documents.
This will update the annotations in all documents to the latest hash of the file.
some mod
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}

		updater := service.NewUpdater(config)
		for _, documentGroup := range config.DocumentGroups {
			err := updater.Update(documentGroup)
			if err != nil {
				log.Fatalf("failed to update document group %s: %v", documentGroup.Path, err)
			}
		}

		log.Printf("updated all annotations")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
