package ports

type Configuration interface {
	GetStorageType() StorageType
}
