package foo

import (
	"context"

	"proj/internal/entities/foo"
	"proj/internal/storage"
	"proj/internal/storage/postgres/shared"
)

type FooService struct {
	repositories storage.Repository
}

// NewTestService returns a new instance of test service.
func NewFooService(repositories storage.Repository) *FooService {
	return &FooService{
		repositories: repositories,
	}
}

func (s *FooService) GetById(ctx context.Context, id int) (foo.Foo, error) {
	return s.repositories.Foo().GetById(ctx, id)
}

func (s *FooService) GetAll(ctx context.Context, limit int, cursor int) ([]foo.Foo, error) {
	paginationParams := shared.PaginationRequest{
		Limit:  limit,
		Cursor: cursor,
	}

	return s.repositories.Foo().GetAll(ctx, paginationParams)
}

func (s *FooService) Create(ctx context.Context, params foo.CreateFooParams) (foo.Foo, error) {
	return s.repositories.Foo().Create(ctx, params)
}

func (s *FooService) Delete(ctx context.Context, id int) error {
	return s.repositories.Foo().Delete(ctx, id)
}

func (s *FooService) Update(ctx context.Context, id int, params foo.UpdateFooParams) (foo.Foo, error) {
	return s.repositories.Foo().Update(ctx, id, params)
}
