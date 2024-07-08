package storerfactory

import (
	"errors"
	"webpcdn/internal/adapters/secondary/storage/dbstorage"
	"webpcdn/internal/adapters/secondary/storage/filestorage"
	redisstorer "webpcdn/internal/adapters/secondary/storage/redisstorage"
	"webpcdn/internal/ports"
)

var (
	errInvalidStorer = errors.New("invalid storer type, cannot resolve data storer")
)

func New(t ports.StorageType) (ports.Storer, error) {
	switch t {
	case ports.TypeFile:
		return filestorage.New(), nil
	case ports.TypeRedis:
		return redisstorer.New(), nil
	case ports.TypeDB:
		return dbstorage.New(), nil
	default:
		return nil, errInvalidStorer
	}
}
