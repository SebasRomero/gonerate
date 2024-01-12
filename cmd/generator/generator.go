/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package generator

import (
	"github.com/spf13/cobra"
)

// generatorCmd represents the generator command
var GeneratorCmd = &cobra.Command{
	Use:   "gonerator",
	Short: "Generate a rest-api for you",
	Long:  "This command will generate a rest-api for you, using the http package.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
