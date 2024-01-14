package generator

import (
	"log"
	filegenerator "sebasromero/github.com/api-generator/cmd/generator/file-generator"
	types "sebasromero/github.com/api-generator/cmd/generator/types"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Initialize the generation",
	Long:  `It will start the api generation`,
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()
	},
}

var rest = &cobra.Command{
	Use:   "rest",
	Short: "To generate a rest api",
	Long:  `It will generate a basic rest api`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Must specify a project name")
		}
		filegenerator.InitProject(args, string(types.Rest))
	},
}

var graph = &cobra.Command{
	Use:   "graph",
	Short: "To generate a graphql api",
	Long:  `It will generate a basic graphql api`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Must specify a project name")
		}
		filegenerator.InitProject(args, string(types.Graph))
	},
}

func init() {
	GeneratorCmd.AddCommand(newCmd)
	newCmd.AddCommand(rest)
	newCmd.AddCommand(graph)
}
