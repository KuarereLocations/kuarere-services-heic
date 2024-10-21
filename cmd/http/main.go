package main

import (
	imageprocessing "kuarere/internal/adapter/Imageprocessing"
	"kuarere/internal/adapter/config"
	"kuarere/internal/adapter/handler/http/himage"
	"kuarere/internal/adapter/http"
	"kuarere/internal/core/domain/static"
	"kuarere/internal/core/services"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			// routes
			http.AsRoute(himage.NewHandlerImageConverter),

			// core
			config.NewConfig,

			// http server
			fx.Annotate(
				http.NewHttpServer,
				fx.ParamTags(static.FxGroupRoutes),
			),

			// image proccessing
			imageprocessing.NewImageProccessing,

			// services
			services.NewImageConverterService,
		),
		fx.Invoke(func(
			*http.HttpServer,
		) {
		}),
	)
	app.Run()
}
