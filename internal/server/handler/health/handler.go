package health

import (
	"context"

	"proj/internal/service"

	"go.uber.org/zap"
)

type httpHandler struct {
	healthService service.HealthService
	logger        *zap.Logger
}

func newHTTPHandler(healthService service.HealthService, logger *zap.Logger) *httpHandler {
	return &httpHandler{
		healthService: healthService,
		logger:        logger,
	}
}

type HealthCheckOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

func (h *httpHandler) healthCheck(ctx context.Context, input *struct{}) (*HealthCheckOutput, error) {
	err := h.healthService.HealthCheck(ctx)
	if err != nil {
		h.logger.Error("failed to health check", zap.Error(err))
	}

	resp := &HealthCheckOutput{}
	resp.Body.Message = "Health check passed"

	return resp, nil
}
