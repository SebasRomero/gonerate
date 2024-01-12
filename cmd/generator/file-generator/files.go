package filegenerator

import (
	"log"
	"os"
)

func InitProject() {
	createFolder(getCurrentWorkDirectory())
	createFile(getCurrentWorkDirectory())
}

func getCurrentWorkDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	currentDirectory := dir + "/yournewAPi/"
	return currentDirectory
}

func createFile(currentDirectory string) {
	newFile := currentDirectory + "example.go"
	os.Create(newFile)
}

func createFolder(currentDirectory string) {
	if err := os.MkdirAll(getCurrentWorkDirectory(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
