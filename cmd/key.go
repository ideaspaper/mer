package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Read/store your api key from/to config file",
	Long:  "Read/store your API key from/to config file\nOnly accept maximum of 1 argument, which is the API key you want to store\nIf no argument given, your stored API key will be shown instead",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("API_KEY:", viper.GetString("API_KEY"))
		} else {
			f, err := os.OpenFile(viper.ConfigFileUsed(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				log.Fatalln(err)
			}
			_, err = f.Write([]byte(fmt.Sprintf("API_KEY=%s", args[0])))
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("API_KEY:", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(keyCmd)
}
