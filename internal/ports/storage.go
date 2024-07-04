package ports

// Storer is an interface to store and read data, Can be filesystem, database, cache, S3, or whatever
// It is used in the app as source for images to converd, and source for images to cache, they can be separate implementations of the storer
type Storer interface {
	Read(string) ([]byte, error)
	Write(string, []byte) error
	IsExists(string) bool
}
