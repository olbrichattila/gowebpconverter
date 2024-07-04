package ports

// RequestFile is the domain object of the file name
type RequestFile interface {
	FileName() string
}
