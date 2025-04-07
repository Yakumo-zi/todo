package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo item",
	Long:  `Add a new todo item to the list.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
