package http

import (
	"proj/internal/server/handler/account"
	"proj/internal/server/handler/foo"
	"proj/internal/server/handler/health"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

func (s *Server) routes() chi.Router {
	router := chi.NewMux()

	config := huma.DefaultConfig(s.apiName, s.apiVersion)
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearerAuth": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	humaApi := humachi.New(router, config)

	foo.RegisterHumaRoutes(
		s.services.FooService,
		humaApi,
		s.logger,
	)

	account.RegisterHumaRoutes(
		s.services.AccountService,
		humaApi,
		s.logger,
		s.supabaseClient,
	)

	health.RegisterHumaRoutes(
		s.services.HealthService,
		humaApi,
		s.logger,
	)

	return router
}
