package http

import (
	"io"
	"mime/multipart"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type IRouteConfig interface {
	Pattern() string
	Method() string
	GetBodyLimit() uint
	SetBodyLimit(bodyLimit uint)
}

type Route interface {
	Handler(c *Ctx) error
	Config() IRouteConfig
}

type RouteConfig struct {
	pattern   string
	method    string
	BodyLimit uint //MB
}

func (rc *RouteConfig) Method() string {
	return rc.method
}

func (rc *RouteConfig) Pattern() string {
	return rc.pattern
}

func (rc *RouteConfig) GetBodyLimit() uint {
	if rc.BodyLimit == 0 {
		return 4
	}
	return rc.BodyLimit
}

func (rc *RouteConfig) SetBodyLimit(bodyLimit uint) {
	rc.BodyLimit = bodyLimit
}

func NewRouteConfig(pattern string, method string) IRouteConfig {
	return &RouteConfig{
		pattern: pattern,
		method:  method,
	}
}

type Ctx struct {
	c *fiber.Ctx
}

func (ctx *Ctx) Error(message interface{}) error {
	ctx.c.Status(500)
	var messageParse string

	if messageValue, ok := message.(string); ok {
		messageParse = messageValue
	}
	if messageValue, ok := message.(error); ok {
		messageParse = messageValue.Error()
	}

	return ctx.c.JSON(map[string]any{
		"error": messageParse,
	})
}

func (ctx *Ctx) JSON(value interface{}) error {
	return ctx.c.JSON(value)
}

func (ctx *Ctx) SendMessage(value string) error {
	return ctx.JSON(map[string]any{
		"message": value,
	})
}

func (ctx *Ctx) Send(body []byte) error {
	return ctx.c.Send(body)
}

func (ctx *Ctx) BodyParser(out interface{}) error {
	if err := ctx.c.BodyParser(out); err != nil {
		return err
	}
	return validate.Struct(out)
}

func (ctx *Ctx) Writer() io.Writer {
	return ctx.c
}
func (ctx *Ctx) Params(name string) string {
	return ctx.c.Params(name)
}

func (ctx *Ctx) FormFile(key string) (*multipart.FileHeader, error) {
	return ctx.c.FormFile(key)
}

func (ctx *Ctx) Redirect(location string, status ...int) error {
	return ctx.c.Redirect(location, status...)
}

func (ctx *Ctx) Query(key string, defaultValue ...string) string {
	return ctx.c.Query(key, defaultValue...)
}

func (ctx *Ctx) SetContentType(contentType string) {
	ctx.c.Response().Header.SetContentType(contentType)
}
