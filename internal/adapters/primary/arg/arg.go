package argparser

import (
	"errors"
	"os"
	request "webpcdn/internal/core/domain"
	"webpcdn/internal/ports"
)

var (
	errInvalidArgumentCount = errors.New("invalid argument count, it should be the file name")
)

func New() ports.Arger {
	return &parser{}
}

type parser struct {
}

func (t *parser) FileName() (ports.RequestFile, error) {
	if len(os.Args) > 1 {
		return request.NewFile(os.Args[1])
	}

	return nil, errInvalidArgumentCount
}
