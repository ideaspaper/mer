package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var readFlag *bool

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "create config file which holds the api key",
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if *readFlag {
			fmt.Println("API_KEY:", viper.GetString("API_KEY"))
		} else {
			if len(args) == 1 {
				f, err := os.OpenFile(viper.ConfigFileUsed(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				_, err = f.Write([]byte(fmt.Sprintf("API_KEY=%s", args[0])))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println("API_KEY:", args[0])
				defer f.Close()
			} else {
				fmt.Println("only accepts 1 argument")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(keyCmd)
	readFlag = keyCmd.Flags().BoolP("read", "r", false, "task id")
}
