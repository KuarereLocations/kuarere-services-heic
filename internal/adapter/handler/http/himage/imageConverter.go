package himage

import (
	"bytes"
	"kuarere/internal/adapter/http"
	"kuarere/internal/core/services"
	"mime/multipart"
)

type HandlerImageConverter struct {
	imageConverterService *services.ImageConverterService
}

func (*HandlerImageConverter) Pattern() string {
	return "api/v1/image/converter"
}

func (h *HandlerImageConverter) Config() http.IRouteConfig {
	return http.NewRouteConfig(h.Pattern(), http.MethodPost)
}

func NewHandlerImageConverter(
	imageConverterService *services.ImageConverterService,
) http.Route {
	return &HandlerImageConverter{
		imageConverterService: imageConverterService,
	}
}

func (h *HandlerImageConverter) Handler(c *http.Ctx) error {
	var imageFileHeader *multipart.FileHeader
	var err error

	if imageFileHeader, err = c.FormFile("img"); err != nil {
		return err
	}

	postFile, err := imageFileHeader.Open()
	if err != nil {
		return err
	}

	var image = new(bytes.Buffer)
	if err = h.imageConverterService.ConvertToPng(postFile, image); err != nil {
		return err
	}
	c.SetContentType("image/png")
	_, err = c.Writer().Write(image.Bytes())
	return err
}
