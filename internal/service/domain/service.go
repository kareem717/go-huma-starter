package domain

import (
	"proj/internal/service"
	"proj/internal/service/domain/account"
	"proj/internal/service/domain/health"
	"proj/internal/service/domain/foo"
	"proj/internal/storage"
)

// NewService implementation for storage of all services.
func NewService(
	repositories storage.Repository,
) *service.Service {
	return &service.Service{
		FooService: foo.NewFooService(repositories),
		AccountService: account.NewAccountService(repositories),
		HealthService:  health.NewHealthService(repositories),
	}
}
