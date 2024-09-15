package http

import (
	"net/http"

	"proj/internal/service"

	"github.com/supabase-community/supabase-go"

	"go.uber.org/zap"
)

type Server struct {
	services       *service.Service
	apiName        string
	apiVersion     string
	logger         *zap.Logger
	supabaseClient *supabase.Client
}

func NewServer(
	services *service.Service,
	apiName, apiVersion string,
	logger *zap.Logger,
	supabaseClient *supabase.Client,
) *Server {
	return &Server{
		services:       services,
		apiName:        apiName,
		apiVersion:     apiVersion,
		logger:         logger,
		supabaseClient: supabaseClient,
	}
}

func (s *Server) Serve(port string) error {
	router := s.routes()

	return http.ListenAndServe(port, router)
}
