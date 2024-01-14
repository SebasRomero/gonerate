/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"sebasromero/github.com/api-generator/cmd/generator"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gonerate",
	Short: "Gonerate is a cli application to generate a basic api.",
	Long: `Gonerate is a cli application to generate a basic api, chossing
	among rest-api or a graphql api.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(generator.GeneratorCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.api-generator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommandPalettes()
}
