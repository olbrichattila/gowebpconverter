package storage

func NewMock() *storemock {
	return &storemock{}
}

type storemock struct {
	isExist     bool
	readCalled  int
	writeCalled int
}

func (s *storemock) WithExists(v bool) *storemock {
	s.isExist = v
	return s
}

func (s *storemock) ReadCalled() int {
	return s.readCalled
}

func (s *storemock) WriteCalled() int {
	return s.writeCalled
}

// isExists implements ports.Storer.
func (s *storemock) IsExists(filename string) bool {
	return s.isExist
}

// Read implements ports.Storer.
func (s *storemock) Read(filename string) ([]byte, error) {
	s.readCalled++
	return []byte("test"), nil
}

// Write implements ports.Storer.
func (s *storemock) Write(filename string, data []byte) error {
	s.writeCalled++
	return nil
}
