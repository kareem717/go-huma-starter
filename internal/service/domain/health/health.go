package health

import (
	"context"
	"time"

	"proj/internal/storage"
)

type HealthService struct {
	repositories storage.Repository
}

// NewHealthService returns a new instance of health service.
func NewHealthService(repositories storage.Repository) *HealthService {
	return &HealthService{
		repositories: repositories,
	}
}

func (s *HealthService) HealthCheck(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repositories.HealthCheck(ctx)
}
