package ports

import (
	"image"
	"io"
)

type ImageProccessing interface {
	ReadHeif(heicImage io.Reader) (image.Image, error)
	ReadAvif(avifImage io.Reader) (image.Image, error)
	ImageToPng(img image.Image, out io.Writer) error
}
