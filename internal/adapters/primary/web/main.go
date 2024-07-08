package webadapter

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	request "webpcdn/internal/core/domain/request"
	"webpcdn/internal/ports"
)

func New(
	converter ports.Converter,
	cacher ports.Cacher,
) ports.MainAdapter {
	return &adapter{
		converter: converter,
		cacher:    cacher,
	}
}

type adapter struct {
	converter ports.Converter
	cacher    ports.Cacher
}

func (t *adapter) Run() error {
	http.HandleFunc("/", t.handler)
	return http.ListenAndServe(":8000", nil)
}

func (t *adapter) handler(w http.ResponseWriter, r *http.Request) {
	imagePath := strings.TrimPrefix(r.URL.Path, "/")
	// TODO add parsing ?wh and other parameters to crop

	rf, err := request.NewFile(imagePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch image %s, %s", imagePath, err.Error()), http.StatusNotFound)
		return
	}

	buf, err := t.cacher.Retrieve(rf, func() ([]byte, error) {
		return t.converter.Convert(rf)
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch image %s", imagePath), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/webp")
	if _, err := w.Write(buf); err != nil {
		log.Printf("failed to write image to ResponseWriter: %v", err)
	}
}
