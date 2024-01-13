package filegenerator

import (
	"log"
	"os"
	"sebasromero/github.com/api-generator/cmd/generator/types"
)

func InitProject(args []string, typeApi string) {
	if typeApi == string(types.Rest) { //TODO Rest
		projectName := args[0]
		createFolder(getCurrentWorkDirectory(projectName), args[0])
		createFile(getCurrentWorkDirectory(projectName), args[0])

	} else { //TODO Graph
		projectName := args[0]
		createFolder(getCurrentWorkDirectory(projectName), args[0])
		createFile(getCurrentWorkDirectory(projectName), args[0])

	}
}

func getCurrentWorkDirectory(projectName string) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	currentDirectory := dir + "/" + projectName + "/"
	return currentDirectory
}

func createFile(currentDirectory string, projectName string) {
	newFile := currentDirectory + projectName + ".go"
	os.Create(newFile)
}

func createFolder(currentDirectory string, projectName string) {
	if err := os.MkdirAll(getCurrentWorkDirectory(projectName), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
