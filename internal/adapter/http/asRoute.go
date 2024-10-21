package http

import (
	"kuarere/internal/core/domain/static"

	"go.uber.org/fx"
)

func AsRoute(f any, params ...fx.Annotation) any {
	var annotate = []fx.Annotation{
		fx.As(new(Route)),
		fx.ResultTags(static.FxGroupRoutes),
	}
	annotate = append(annotate, params...)

	return fx.Annotate(
		f,
		annotate...,
	)
}
