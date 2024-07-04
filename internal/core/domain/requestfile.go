package request

import (
	"errors"
	"path/filepath"
	"strings"
	"webpcdn/internal/ports"
)

var (
	errInvalidImageFileExtension = errors.New("invalid file extension")
)

// NewFile creates a new file domain object, validate file extension
func NewFile(n string) (ports.RequestFile, error) {

	rf := &rFile{}

	if rf.isImageFile(n) {
		rf.fPath = n
		return rf, nil
	}

	return nil, errInvalidImageFileExtension
}

type rFile struct {
	fPath string
}

func (t rFile) FileName() string {
	return t.fPath
}

func (t rFile) isImageFile(n string) bool {
	validExtensions := []string{".jpg", ".jpeg", ".gif", ".png", ".svg", ".tiff", ".webp"}

	ext := strings.ToLower(filepath.Ext(n))
	for _, validExt := range validExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}
