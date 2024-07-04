package ports

// The converter will convert the image to webp
type Converter interface {
	Convert(RequestFile) ([]byte, error)
}
