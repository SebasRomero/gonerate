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
		createMain(getCurrentWorkDirectory(projectName))

	} else { //TODO Graph
		projectName := args[0]
		createFolder(getCurrentWorkDirectory(projectName), args[0])
		createMain(getCurrentWorkDirectory(projectName))

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

func createMain(currentDirectory string) {
	newFile := currentDirectory + "main.go"
	file, err := os.Create(newFile)
	if err != nil {
		log.Fatal(err)
	}

	writting := []byte(initServerWriter())
	file.Write(writting)

}

func createFolder(currentDirectory string, projectName string) {
	if err := os.MkdirAll(getCurrentWorkDirectory(projectName), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
