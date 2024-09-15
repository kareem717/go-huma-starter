package foo

import (
	"context"
	"database/sql"
	"errors"

	"proj/internal/entities/foo"
	"proj/internal/server/handler/shared"
	"proj/internal/service"

	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/zap"
)

type httpHandler struct {
	fooService service.FooService
	logger         *zap.Logger
}

func newHTTPHandler(fooService service.FooService, logger *zap.Logger) *httpHandler {
	return &httpHandler{
		fooService: fooService,
		logger:         logger,
	}
}

type SingleFooResponse struct {
	Body struct {
		shared.MessageResponse
		Foo *foo.Foo `json:"foo"`
	}
}

func (h *httpHandler) getByID(ctx context.Context, input *shared.PathIDParam) (*SingleFooResponse, error) {
	foo, err := h.fooService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Foo not found")
		default:
			h.logger.Error("failed to fetch foo", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the foo")
		}
	}

	resp := &SingleFooResponse{}
	resp.Body.Message = "Foo fetched successfully"
	resp.Body.Foo = &foo

	return resp, nil
}

type GetAllFooOutput struct {
	Body struct {
		shared.MessageResponse
		Foos []foo.Foo `json:"foos"`
		shared.PaginationResponse
	}
}

func (h *httpHandler) getAll(ctx context.Context, input *shared.PaginationRequest) (*GetAllFooOutput, error) {
	LIMIT := input.Limit + 1

	foos, err := h.fooService.GetAll(ctx, LIMIT, input.Cursor)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Foos not found")
		default:
			h.logger.Error("failed to fetch foos", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the foos")
		}
	}

	resp := &GetAllFooOutput{}
	resp.Body.Message = "Foos fetched successfully"
	resp.Body.Foos = foos

	if len(foos) == LIMIT {
		resp.Body.NextCursor = &foos[len(foos)-1].ID
		resp.Body.HasMore = true
		resp.Body.Foos = resp.Body.Foos[:len(resp.Body.Foos)-1]
	}

	return resp, nil
}

type CreateFooInput struct {
	Body foo.CreateFooParams `json:"foo"`
}

func (h *httpHandler) create(ctx context.Context, input *CreateFooInput) (*SingleFooResponse, error) {
	foo, err := h.fooService.Create(ctx, input.Body)
	if err != nil {
		h.logger.Error("failed to create foo", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while creating the foo")
	}

	resp := &SingleFooResponse{}
	resp.Body.Message = "Foo created successfully"
	resp.Body.Foo = &foo

	return resp, nil
}

type UpdateFooInput struct {
	shared.PathIDParam
	Body foo.UpdateFooParams `json:"foo"`
}

func (h *httpHandler) update(ctx context.Context, input *UpdateFooInput) (*SingleFooResponse, error) {
	_, err := h.fooService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Foo not found")
		default:
			h.logger.Error("failed to fetch foo", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the foo")
		}
	}

	foo, err := h.fooService.Update(ctx, input.ID, input.Body)

	if err != nil {
		h.logger.Error("failed to update foo", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while updating the foo")
	}

	resp := &SingleFooResponse{}
	resp.Body.Message = "Foo updated successfully"
	resp.Body.Foo = &foo

	return resp, nil
}

type DeleteFooResponse struct {
	Body shared.MessageResponse
}

func (h *httpHandler) delete(ctx context.Context, input *shared.PathIDParam) (*DeleteFooResponse, error) {
	_, err := h.fooService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Foo not found")
		default:
			h.logger.Error("failed to fetch foo", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the foo")
		}
	}

	err = h.fooService.Delete(ctx, input.ID)
	if err != nil {
		h.logger.Error("failed to delete foo", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while deleting the foo")
	}

	resp := &DeleteFooResponse{}
	resp.Body.Message = "Foo deleted successfully"

	return resp, nil
}
