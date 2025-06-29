package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fresh",
	Short: "FreshDocs - Keep your documentation as fresh as your code!",
	Long: `FreshDocs is an app that helps developers and users maintain a real, 
actionable connection between documentation and source code. 
By embedding hidden comment annotations in doc files, you can link docs to specific code files or folders. 
Whenever the code changes, FreshDocs detects which documentation is now stale — so your docs never fall behind.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(findCmd)
}
