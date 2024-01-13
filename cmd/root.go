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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommandPalettes()
}
