package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ideaspaper/mer/controllers"
	"github.com/ideaspaper/mer/services"
	"github.com/ideaspaper/mer/views"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	service    services.IColService
	view       views.IColView
	controller controllers.IColController
)

var rootCmd = &cobra.Command{
	Use:   "mer",
	Short: "Mer is a small dictionary app based on Merriam-Webster free Collegiate Dictionary API",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func init() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	exPath := filepath.Dir(ex)
	configPath := filepath.Join(exPath, "mer_config.env")

	viper.SetConfigName("mer_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(exPath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("config file %v does not exist\n", configPath)
			fmt.Printf("creating %v\n", configPath)
			os.WriteFile(configPath, []byte(""), 0666)
			fmt.Printf("file %v created successfully\n", configPath)
			os.Exit(1)
		}
	}

	service = services.NewColService("https://www.dictionaryapi.com/api/v3/references/collegiate/json")
	view = views.NewColView()
	controller = controllers.NewColController(service, view)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
