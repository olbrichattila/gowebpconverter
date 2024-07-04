package ports

// Arger gets the first command line argument, returns RequestFile domain object
type Arger interface {
	FileName() (RequestFile, error)
}
