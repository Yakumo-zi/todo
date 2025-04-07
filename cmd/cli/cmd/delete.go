package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo item",
	Long:  `Delete a todo item from the list.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
