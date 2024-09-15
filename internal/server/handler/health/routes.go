package health

import (
	"net/http"

	"proj/internal/service"

	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/zap"
)

func RegisterHumaRoutes(
	healthService service.HealthService,
	humaApi huma.API,
	logger *zap.Logger,
) {

	handler := &httpHandler{
		healthService: healthService,
		logger:        logger,
	}

	// Register GET /test/{id}
	huma.Register(humaApi, huma.Operation{
		OperationID: "health-check",
		Method:      http.MethodGet,
		Path:        "/health",
		Summary:     "Health check",
		Description: "Health check.",
		Tags:        []string{"Health"},
	}, handler.healthCheck)
}
