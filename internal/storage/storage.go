package storage

import (
	"errors"
	"github.com/MadiyarOne/oz1/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File // map of files, where key is uuid and value is
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) UploadFile(name string, blob []byte) (uuid.UUID, error) {
	var newFile = file.NewFile(name, blob)
	if s.files == nil {
		s.files = make(map[uuid.UUID]*file.File)
	}
	s.files[newFile.Id] = newFile
	return newFile.Id, nil
}

func (s *Storage) GetFile(id uuid.UUID) (*file.File, error) {
	foundFile, ok := s.files[id]
	if !ok {
		return nil, errors.New("file not found")
	}

	return foundFile, nil
}
