package account

import (
	"context"

	"proj/internal/entities/account"
	"proj/internal/storage"
	"proj/internal/storage/postgres/shared"

	"github.com/google/uuid"
)

type AccountService struct {
	repositories storage.Repository
}

// NewTestService returns a new instance of test service.
func NewAccountService(repositories storage.Repository) *AccountService {
	return &AccountService{
		repositories: repositories,
	}
}

func (s *AccountService) GetById(ctx context.Context, id int) (account.Account, error) {
	return s.repositories.Account().GetById(ctx, id)
}

func (s *AccountService) GetByUserId(ctx context.Context, userId uuid.UUID) (account.Account, error) {
	return s.repositories.Account().GetByUserId(ctx, userId)
}

func (s *AccountService) GetAll(ctx context.Context, limit int, cursor int) ([]account.Account, error) {
	paginationParams := shared.PaginationRequest{
		Limit:  limit,
		Cursor: cursor,
	}

	return s.repositories.Account().GetAll(ctx, paginationParams)
}

func (s *AccountService) Create(ctx context.Context, params account.CreateAccountParams) (account.Account, error) {
	return s.repositories.Account().Create(ctx, params)
}

func (s *AccountService) Delete(ctx context.Context, id int) error {
	return s.repositories.Account().Delete(ctx, id)
}

func (s *AccountService) Update(ctx context.Context, id int, params account.UpdateAccountParams) (account.Account, error) {
	return s.repositories.Account().Update(ctx, id, params)
}
