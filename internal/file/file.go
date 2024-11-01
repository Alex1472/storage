package file

import (
	"fmt"
	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(name string, blob []byte) (*File, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		ID:   id,
		Name: name,
		Data: blob,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("File(%v, %s)", f.ID, f.Name)
}
