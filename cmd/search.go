package cmd

import "github.com/spf13/cobra"

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the meaning of a word",
	Long:  "Search the meaning of a word\nOnly accept exactly 1 argument, which is the word you want to search",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controller.Search(args[0])
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
