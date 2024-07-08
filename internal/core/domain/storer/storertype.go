package storertype

import (
	"errors"
	"webpcdn/internal/ports"
)

var (
	errInvalidStorageType = errors.New("invalid storage type")
)

func New(t ports.StorageType) (ports.Storager, error) {
	if t == ports.TypeDB || t == ports.TypeFile || t == ports.TypeRedis {
		return &storageType{
			t: t,
		}, nil
	}

	return nil, errInvalidStorageType
}

type storageType struct {
	t ports.StorageType
}

func (t *storageType) Get() ports.StorageType {
	return t.t
}
