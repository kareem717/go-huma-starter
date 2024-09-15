package foo

import (
	"context"
	"database/sql"

	"proj/internal/entities/foo"
	"proj/internal/storage/postgres/shared"

	"github.com/uptrace/bun"
)

type FooRepository struct {
	db  bun.IDB
	ctx context.Context
}

// NewFooRepository returns a new instance of the repository.
func NewFooRepository(db bun.IDB, ctx context.Context) *FooRepository {
	return &FooRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *FooRepository) Create(ctx context.Context, params foo.CreateFooParams) (foo.Foo, error) {
	resp := foo.Foo{}

	err := r.db.
		NewInsert().
		Model(&params).
		ModelTableExpr("foos").
		Returning("*").
		Scan(ctx, &resp)

	return resp, err
}

func (r *FooRepository) Update(ctx context.Context, id int, params foo.UpdateFooParams) (foo.Foo, error) {
	resp := foo.Foo{}

	err :=
		r.db.
			NewUpdate().
			Model(&params).
			ModelTableExpr("foos").
			Where("id = ?", id).
			Returning("*").
			OmitZero().
			Scan(ctx, &resp)

	return resp, err
}

func (r *FooRepository) Delete(ctx context.Context, id int) error {
	res, err :=
		r.db.
			NewDelete().
			Model(&foo.Foo{}).
			Where("id = ?", id).
			Exec(ctx)

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (r *FooRepository) GetById(ctx context.Context, id int) (foo.Foo, error) {
	resp := foo.Foo{}

	err := r.db.
		NewSelect().
		Model(&resp).
		Where("id = ?", id).
		Scan(ctx)

	return resp, err
}

func (r *FooRepository) GetAll(ctx context.Context, paginationParams shared.PaginationRequest) ([]foo.Foo, error) {
	resp := []foo.Foo{}

	err := r.db.
		NewSelect().
		Model(&resp).
		Where("id >= ?", paginationParams.Cursor).
		Order("id").
		Limit(paginationParams.Limit).
		Scan(ctx)

	return resp, err
}
