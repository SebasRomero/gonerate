package filegenerator

import (
	"log"
	"os"
	"sebasromero/github.com/api-generator/cmd/generator/types"
)

func InitProject(args []string, typeApi string) {
	projectName := args[0]
	if typeApi == string(types.Rest) { //TODO Rest
		createFolder(getCurrentWorkDirectory(projectName), args[0])
		createFolder(getCurrentWorkDirectory(projectName), projectName+"/server")
		createFile(getCurrentWorkDirectory(projectName), "main.go")
		createFile(getCurrentWorkDirectory(projectName), "server/server.go")

	} else { //TODO Graph
		createFolder(getCurrentWorkDirectory(projectName), args[0])
		createFile(getCurrentWorkDirectory(projectName), "")

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

func createFile(currentDirectory string, nameFile string) {
	newFile := currentDirectory + nameFile
	file, err := os.Create(newFile)
	if err != nil {
		log.Fatal(err)
	}

	switch nameFile {
	case string("server/" + types.Server):
		writting := []byte(initServerWriter())
		file.Write(writting)

	case string(types.Main):
		writting := []byte(initMainWriter())
		file.Write(writting)

	}
}

func createFolder(currentDirectory string, projectName string) {
	if err := os.MkdirAll(getCurrentWorkDirectory(projectName), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
