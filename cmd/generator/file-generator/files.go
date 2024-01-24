package filegenerator

import (
	"log"
	"os"
	"sebasromero/github.com/api-generator/cmd/generator/types"
)

func InitProject(args []string, typeApi string) {
	projectName := args[0]
	if typeApi == string(types.Rest) { //TODO Rest
		getCurrentDirectory := getCurrentWorkDirectory(projectName)
		routeServer := string(types.RouteServer)

		createFolder(getCurrentDirectory, args[0])
		createFolder(getCurrentDirectory, projectName+"/server")

		createFile(getCurrentDirectory, "", string(types.Main))
		createFile(getCurrentDirectory, routeServer, string(types.Server))
		createFile(getCurrentDirectory, routeServer, projectName)
		createFile(getCurrentDirectory, routeServer, string(types.Routes))

	} else { //TODO Graph
		createFolder(getCurrentWorkDirectory(projectName), args[0])
		createFile(getCurrentWorkDirectory(projectName), "", "")

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

func createFile(currentDirectory string, route string, nameFile string) {
	newFile := currentDirectory + route + nameFile
	file, err := os.Create(newFile)
	if err != nil {
		log.Fatal(err)
	}

	switch nameFile {
	case string(types.Server):
		writting := []byte(initServerWriter())
		file.Write(writting)

	case string(types.Main):
		writting := []byte(initMainWriter())
		file.Write(writting)

	case string(types.Routes):
		writting := []byte(initRoutes(nameFile))
		file.Write((writting))

	default:
		os.Rename(newFile, newFile+".go")
		writting := []byte(initTopic(nameFile)) //This isn't working well
		file.Write(writting)
	}
}

func createFolder(currentDirectory string, projectName string) {
	if err := os.MkdirAll(getCurrentWorkDirectory(projectName), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
