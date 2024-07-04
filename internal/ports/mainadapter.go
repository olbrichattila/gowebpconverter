package ports

// MainAdapter is the entry point of the application, main adapters
type MainAdapter interface {
	Run() error
}
