package file

import (
	"fmt"
	"github.com/google/uuid"
)

type File struct {
	Id   uuid.UUID
	Name string
	Data []byte
}

func NewFile(name string, data []byte) *File {
	var id = uuid.New()
	return &File{Id: id, Name: name, Data: data}
}

func (f File) String() string {
	return fmt.Sprintf("File: %s, %s, %s", f.Id, f.Name, string(f.Data))
}
