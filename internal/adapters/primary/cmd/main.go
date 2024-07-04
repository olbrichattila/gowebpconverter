package cmdadapter

import (
	"os"

	"webpcdn/internal/ports"
)

func New(
	arger ports.Arger,
	converter ports.Converter,
	cacher ports.Cacher,
) ports.MainAdapter {
	return &adapter{
		arger:     arger,
		converter: converter,
		cacher:    cacher,
	}
}

type adapter struct {
	arger       ports.Arger
	converter   ports.Converter
	cacher      ports.Cacher
	fileStorer  ports.Storer
	cacheStorer ports.Storer
}

func (t *adapter) Run() error {
	rf, err := t.arger.FileName()
	if err != nil {
		return err
	}

	buf, err := t.cacher.Retrieve(rf, func() ([]byte, error) {
		return t.converter.Convert(rf)
	})
	if err != nil {
		return err
	}

	if _, err := os.Stdout.Write(buf); err != nil {
		return err
	}

	return nil
}
