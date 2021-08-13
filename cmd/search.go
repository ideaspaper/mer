package cmd

import "github.com/spf13/cobra"

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search the word",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controller.Search(args[0])
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
