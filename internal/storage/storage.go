package storage

import (
	"errors"
	"github.com/MadiyarOne/oz1/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[string]*file.File // map of files, where key is uuid and value is
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) UploadFile(name string, blob []byte) (uuid.UUID, error) {
	var newFile = file.NewFile(name, blob)
	if s.files == nil {
		s.files = make(map[string]*file.File)
	}
	s.files[name] = newFile
	return newFile.Id, nil
}

func (s *Storage) GetFile(fileName string) (*file.File, error) {
	foundFile, ok := s.files[fileName]
	if !ok {
		return nil, errors.New("file not found")
	}

	return foundFile, nil
}
