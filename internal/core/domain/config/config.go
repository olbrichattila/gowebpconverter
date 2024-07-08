package config

import (
	"os"
	"webpcdn/internal/ports"
)

func New() ports.Configuration {
	return &conf{}
}

type conf struct {
}

func (t *conf) GetStorageType() ports.StorageType {
	var stype ports.StorageType
	cahceType := os.Getenv("CACHE_TYPE")
	switch cahceType {
	case "file":
		stype = ports.TypeFile
	case "redis":
		stype = ports.TypeRedis
	case "db":
		stype = ports.TypeDB
	default:
		stype = ports.TypeFile
	}
	return stype
}
