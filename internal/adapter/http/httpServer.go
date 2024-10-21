package http

import (
	"context"
	"fmt"
	"kuarere/internal/adapter/config"
	"log"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"go.uber.org/fx"
)

type HttpServer struct {
	app    *fiber.App
	config *config.Config
}

func NewHttpServer(routes []Route, config *config.Config, lc fx.Lifecycle) *HttpServer {
	var apiService = new(HttpServer)
	app := fiber.New(
		fiber.Config{
			BodyLimit: 500 * 1024 * 1024, //500 MB
		},
	)
	apiService.app = app
	app.Use(cors.New())
	apiService.SetRoutes(routes)
	apiService.config = config

	lc.Append(fx.StartHook(func(ctx context.Context) error {
		log.Println("NewApiService: Iniciando")
		// check port is open
		var ln, err = net.Listen("tcp", "127.0.0.1:"+config.Port())

		if err != nil {
			log.Println("NewApiService: open port error")
			return err
		}
		ln.Close()

		// init app
		var errChanelInit = make(chan error)

		go func() {
			errChanelInit <- app.Listen(":" + config.Port())
		}()

		select {
		case err := <-errChanelInit:
			return err
		case <-time.After(time.Millisecond * 200):
		}

		log.Println("NewApiService: Iniciado")
		return nil
	}))

	lc.Append(fx.StopHook(func(ctx context.Context) error {
		fmt.Println("api:init close")
		app.ShutdownWithTimeout(time.Millisecond * 300)
		fmt.Println("api:finish close")
		return nil
	}))

	return apiService
}

func (apiService *HttpServer) SetRoutes(routes []Route) {
	for i := range routes {
		fmt.Printf("Route: %s, Method: %s\n", routes[i].Config().Pattern(), routes[i].Config().Method())
		apiService.SetHandler(routes[i])
	}
}

func (apiService *HttpServer) SetHandler(route Route) {
	apiService.app.Add(route.Config().Method(), route.Config().Pattern(), func(c *fiber.Ctx) error {
		ctx := &Ctx{c: c}
		return route.Handler(ctx)
	})
}
