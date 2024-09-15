package storage

import (
	"context"

	"proj/internal/entities/account"
	"proj/internal/entities/foo"
	"proj/internal/storage/postgres/shared"

	"github.com/google/uuid"
)

type FooRepository interface {
	Create(ctx context.Context, params foo.CreateFooParams) (foo.Foo, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, paginationParams shared.PaginationRequest) ([]foo.Foo, error)
	Update(ctx context.Context, id int, params foo.UpdateFooParams) (foo.Foo, error)
	GetById(ctx context.Context, id int) (foo.Foo, error)
}

type AccountRepository interface {
	Create(ctx context.Context, params account.CreateAccountParams) (account.Account, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, paginationParams shared.PaginationRequest) ([]account.Account, error)
	Update(ctx context.Context, id int, params account.UpdateAccountParams) (account.Account, error)
	GetById(ctx context.Context, id int) (account.Account, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) (account.Account, error)
}

type RepositoryProvider interface {
	Foo() FooRepository
	Account() AccountRepository
}

type Transaction interface {
	RepositoryProvider
	Commit() error
	Rollback() error
	SubTransaction() (Transaction, error)
}

type Repository interface {
	RepositoryProvider
	HealthCheck(ctx context.Context) error
	NewTransaction() (Transaction, error)
	RunInTx(ctx context.Context, fn func(ctx context.Context, tx Transaction) error) error
}
