package ports

type Config interface {
	Get() Configuration
}
