package main

import (
	"fmt"
	"github.com/MadiyarOne/oz1/internal/storage"
)

func main() {
	fileStorage := storage.NewStorage()
	fileName := "test.txt"
	fileStorage.UploadFile(fileName, []byte("Hello World!"))

	foundFile, err := fileStorage.GetFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundFile)
}
