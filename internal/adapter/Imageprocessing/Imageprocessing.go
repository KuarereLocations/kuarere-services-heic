package imageprocessing

import (
	"image"
	"image/png"
	"io"
	"kuarere/internal/core/ports"

	"github.com/gen2brain/avif"
	"github.com/jdeng/goheif"
)

type ImageProccessing struct {
}

func NewImageProccessing() ports.ImageProccessing {
	return &ImageProccessing{}
}

func (ip *ImageProccessing) ReadHeif(heicImage io.Reader) (image.Image, error) {
	return goheif.Decode(heicImage)
}

func (ip *ImageProccessing) ReadAvif(avifImage io.Reader) (image.Image, error) {
	return avif.Decode(avifImage)
}

func (ip *ImageProccessing) ImageToPng(img image.Image, out io.Writer) error {
	var encoder = new(png.Encoder)
	encoder.CompressionLevel = png.BestSpeed
	return encoder.Encode(out, img)
}
