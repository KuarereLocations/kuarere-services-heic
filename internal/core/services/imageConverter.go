package services

import (
	"bytes"
	"errors"
	"image"
	"io"
	"kuarere/internal/core/ports"
	"log"
	"time"

	"github.com/gabriel-vasile/mimetype"
)

type ImageConverterService struct {
	imageproccessing ports.ImageProccessing
}

func NewImageConverterService(imageproccessing ports.ImageProccessing) *ImageConverterService {
	return &ImageConverterService{
		imageproccessing: imageproccessing,
	}
}

func (ic *ImageConverterService) ConvertToPng(img io.Reader, out io.Writer) error {
	var initTime = time.Now()

	var bffCopyImg = new(bytes.Buffer)
	var err error
	var imgOut image.Image
	if _, err = io.Copy(bffCopyImg, img); err != nil {
		return err
	}
	log.Println("timeSinceCopy: ", time.Since(initTime))

	var mimeType = mimetype.Detect(bffCopyImg.Bytes())
	if mimeType == nil {
		return errors.New("mimeType: null")
	}

	log.Println("timeSinceDetect: ", time.Since(initTime))

	if mimeType.String() == "image/heif" ||
		mimeType.String() == "image/heic" {
		if imgOut, err = ic.imageproccessing.ReadHeif(bffCopyImg); err != nil {
			return err
		}
		log.Println("timeSinceRead: ", time.Since(initTime))
		err = ic.imageproccessing.ImageToPng(imgOut, out)
		log.Println("timeSinceToPng: ", time.Since(initTime))
		return err
	}

	if mimeType.String() == "image/avif" {
		if imgOut, err = ic.imageproccessing.ReadAvif(bffCopyImg); err != nil {
			return err
		}
		return ic.imageproccessing.ImageToPng(imgOut, out)
	}

	return errors.New("unsupported type " + mimeType.String())
}
