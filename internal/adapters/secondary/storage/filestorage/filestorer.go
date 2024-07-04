package filestorage

import (
	"os"
	"webpcdn/internal/ports"
)

func New() ports.Storer {
	return &store{}
}

type store struct {
}

// isExists implements ports.Storer.
func (s *store) IsExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// Read implements ports.Storer.
func (s *store) Read(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Write implements ports.Storer.
func (s *store) Write(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}
