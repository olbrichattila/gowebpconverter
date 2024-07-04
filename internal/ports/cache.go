package ports

type CacheFunc = func() ([]byte, error)

// Cacher will retriveve the data blob from cache or store if not exists, TODO: TTL
type Cacher interface {
	Retrieve(RequestFile, CacheFunc) ([]byte, error)
}
