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

func (f File) String() string {
	return fmt.Sprintf("File: %s, %s, %s", f.Id, f.Name, string(f.Data))
}
