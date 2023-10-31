package commands

import (
	"bytes"
	"image"
	"os"

	"github.com/disintegration/imaging"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

const dimensions = 1024

func readPicture(filename string) (*ttnpb.Picture, error) {
	pictureFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer pictureFile.Close()
	img, format, err := image.Decode(pictureFile)
	if err != nil {
		return nil, err
	}
	var encodingFormat imaging.Format
	var mimeType string
	switch format {
	case "jpeg":
		encodingFormat, mimeType = imaging.JPEG, "image/jpeg"
	default:
		encodingFormat, mimeType = imaging.PNG, "image/png"
	}
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	if width < height {
		img = imaging.CropAnchor(img, width, width, imaging.Center)
	} else if width > height {
		img = imaging.CropAnchor(img, height, height, imaging.Center)
	}
	if width > dimensions || height > dimensions {
		img = imaging.Resize(img, dimensions, dimensions, imaging.Lanczos)
	}
	var buf bytes.Buffer
	if err = imaging.Encode(&buf, img, encodingFormat); err != nil {
		return nil, err
	}
	return &ttnpb.Picture{
		Embedded: &ttnpb.Picture_Embedded{
			MimeType: mimeType,
			Data:     buf.Bytes(),
		},
	}, nil
}
