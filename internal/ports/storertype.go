package ports

const (
	TypeFile  = 0
	TypeRedis = 1
	TypeDB    = 2
)

type StorageType int

type Storager interface {
	Get() StorageType
}
