package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	http_server "github.com/amagkn/golang-production-ready-reference/gen/http/profile_v1/server"
	ver1 "github.com/amagkn/golang-production-ready-reference/internal/controller/http/v1"
	"github.com/amagkn/golang-production-ready-reference/internal/usecase"
	"github.com/amagkn/golang-production-ready-reference/pkg/logger"
	"github.com/amagkn/golang-production-ready-reference/pkg/metrics"
	"github.com/amagkn/golang-production-ready-reference/pkg/otel"
)

func ProfileRouter(r *chi.Mux, uc *usecase.UseCase, m *metrics.HTTPServer) {
	v1 := ver1.New(uc)

	r.Handle("/metrics", promhttp.Handler())

	r.Route("/mnepryakhin/my-app/api", func(r chi.Router) {
		r.Use(otel.Middleware)
		r.Use(logger.Middleware)
		r.Use(metrics.NewMiddleware(m))

		r.Route("/v1", func(r chi.Router) {
			mux := http_server.NewStrictHandler(v1, []http_server.StrictMiddlewareFunc{})
			http_server.HandlerFromMux(mux, r)
		})
	})
}
