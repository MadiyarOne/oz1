package main

import (
	"fmt"
	"github.com/MadiyarOne/oz1/internal/storage"
)

func main() {
	fileStorage := storage.NewStorage()
	uploadFileId, _ := fileStorage.UploadFile("test.txt", []byte("Hello World!"))

	foundFile, err := fileStorage.GetFile(uploadFileId)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundFile)
}
