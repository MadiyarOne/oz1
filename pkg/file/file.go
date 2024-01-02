package file

import "github.com/google/uuid"
import "github.com/MadiyarOne/oz1/internal/file"

func NewFile(name string, data []byte) *file.File {
	var id = uuid.New()
	return &file.File{Id: id, Name: name, Data: data}
}
