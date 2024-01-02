package file

import "github.com/google/uuid"

func NewFile(name string, data []byte) *File {
	var id = uuid.New()
	return &File{Id: id, Name: name, Data: data}
}
