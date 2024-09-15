package account

import (
	"context"
	"database/sql"

	"proj/internal/entities/account"
	"proj/internal/storage/postgres/shared"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AccountRepository struct {
	db  bun.IDB
	ctx context.Context
}

// NewAccountRepository returns a new instance of the repository.
func NewAccountRepository(db bun.IDB, ctx context.Context) *AccountRepository {
	return &AccountRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *AccountRepository) Create(ctx context.Context, params account.CreateAccountParams) (account.Account, error) {
	resp := account.Account{}

	err := r.db.
		NewInsert().
		Model(&params).
		ModelTableExpr("accounts").
		Returning("*").
		Scan(ctx, &resp)

	return resp, err
}

func (r *AccountRepository) Update(ctx context.Context, id int, params account.UpdateAccountParams) (account.Account, error) {
	resp := account.Account{}

	err :=
		r.db.
			NewUpdate().
			Model(&params).
			ModelTableExpr("accounts").
			Where("id = ?", id).
			Returning("*").
			OmitZero().
			Scan(ctx, &resp)

	return resp, err
}

func (r *AccountRepository) Delete(ctx context.Context, id int) error {
	res, err :=
		r.db.
			NewDelete().
			Model(&account.Account{}).
			Where("id = ?", id).
			Exec(ctx)

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (r *AccountRepository) GetById(ctx context.Context, id int) (account.Account, error) {
	resp := account.Account{}

	err := r.db.
		NewSelect().
		Model(&resp).
		Where("id = ?", id).
		Scan(ctx)

	return resp, err
}

func (r *AccountRepository) GetByUserId(ctx context.Context, userId uuid.UUID) (account.Account, error) {
	resp := account.Account{}

	err := r.db.
		NewSelect().
		Model(&resp).
		Where("user_id = ?", userId).
		Scan(ctx)

	return resp, err
}

func (r *AccountRepository) GetAll(ctx context.Context, paginationParams shared.PaginationRequest) ([]account.Account, error) {
	resp := []account.Account{}

	err := r.db.
		NewSelect().
		Model(&resp).
		Where("id >= ?", paginationParams.Cursor).
		Order("id").
		Limit(paginationParams.Limit).
		Scan(ctx)

	return resp, err
}
