package converterservice

import (
	"bytes"
	"errors"
	"image"
	"webpcdn/internal/ports"

	"github.com/chai2010/webp"

	_ "image/gif"  // GIF format
	_ "image/jpeg" // JPEG format
	_ "image/png"  // PNG format

	_ "github.com/ajstarks/svgo" // Import SVG support
	_ "github.com/chai2010/tiff" // Import TIFF support
	_ "github.com/chai2010/webp" // Import WebP support
)

var (
	errImageConvert = errors.New("image convect error")
)

func New(storer ports.Storer) ports.Converter {
	return &convert{
		storer: storer,
	}
}

type convert struct {
	storer ports.Storer
}

func (t *convert) Convert(rf ports.RequestFile) ([]byte, error) {

	data, err := t.storer.Read(rf.FileName())
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, errors.Join(err, errImageConvert)
	}

	var buf bytes.Buffer
	err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: 80})
	if err != nil {
		return nil, err
	}

	// // Resize the image to width = 800px preserving the aspect ratio
	// resizedImg := imaging.Resize(img, 800, 0, imaging.Lanczos)

	// // Crop the image to a 400x400 square from the center
	// croppedImg := imaging.CropCenter(resizedImg, 400, 400)
	return buf.Bytes(), nil
}
