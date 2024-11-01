package storage

import (
	"fmt"
	"github.com/Alex1472/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (st *Storage) Upload(name string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(name, blob)
	if err != nil {
		return nil, err
	}

	st.files[newFile.ID] = newFile
	return newFile, nil
}

func (st *Storage) GetByID(id uuid.UUID) (*file.File, error) {
	restoredFile, ok := st.files[id]
	if !ok {
		return nil, fmt.Errorf("file with id = %v not found", id)
	}

	return restoredFile, nil
}
