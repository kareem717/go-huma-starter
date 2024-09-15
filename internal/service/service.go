package service

import (
	"context"

	"proj/internal/entities/account"
	"proj/internal/entities/foo"

	"github.com/google/uuid"
)

type FooService interface {
	Create(ctx context.Context, params foo.CreateFooParams) (foo.Foo, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, params foo.UpdateFooParams) (foo.Foo, error)
	GetById(ctx context.Context, id int) (foo.Foo, error)
	GetAll(ctx context.Context, limit int, cursor int) ([]foo.Foo, error)
}

type AccountService interface {
	Create(ctx context.Context, params account.CreateAccountParams) (account.Account, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, params account.UpdateAccountParams) (account.Account, error)
	GetById(ctx context.Context, id int) (account.Account, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) (account.Account, error)
	GetAll(ctx context.Context, limit int, cursor int) ([]account.Account, error)
}

type HealthService interface {
	HealthCheck(ctx context.Context) error
}

// Service storage of all services.
type Service struct {
	AccountService AccountService
	FooService FooService
	HealthService  HealthService
}
