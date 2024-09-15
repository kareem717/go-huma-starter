package account

import (
	"context"
	"database/sql"
	"errors"

	"proj/internal/entities/account"
	"proj/internal/server/handler/shared"
	"proj/internal/service"

	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/zap"
)

type httpHandler struct {
	accountService service.AccountService
	logger         *zap.Logger
}

func newHTTPHandler(accountService service.AccountService, logger *zap.Logger) *httpHandler {
	return &httpHandler{
		accountService: accountService,
		logger:         logger,
	}
}

type SingleAccountResponse struct {
	Body struct {
		shared.MessageResponse
		Account *account.Account `json:"account"`
	}
}

func (h *httpHandler) getByID(ctx context.Context, input *shared.PathIDParam) (*SingleAccountResponse, error) {
	account, err := h.accountService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Account not found")
		default:
			h.logger.Error("failed to fetch account", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the account")
		}
	}

	resp := &SingleAccountResponse{}
	resp.Body.Message = "Account fetched successfully"
	resp.Body.Account = &account

	return resp, nil
}

func (h *httpHandler) getByUserId(ctx context.Context, input *shared.PathUserIDParam) (*SingleAccountResponse, error) {
	h.logger.Info("getByUserId", zap.Any("input", input))
	h.logger.Error("getByUserId", zap.Any("input", input))
	if user := shared.GetAuthenticatedUser(ctx); user.ID != input.UserID {
		h.logger.Error("input user id does not match authenticated user id",
			zap.Any("input user id", input.UserID),
			zap.Any("authenticated user id", user.ID))

		return nil, huma.Error403Forbidden("Cannot get account for another user")
	}

	account, err := h.accountService.GetByUserId(ctx, input.UserID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Account not found")
		default:
			h.logger.Error("failed to fetch account", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the account")
		}
	}

	resp := &SingleAccountResponse{}
	resp.Body.Message = "Account fetched successfully"
	resp.Body.Account = &account

	return resp, nil
}

type GetAllAccountOutput struct {
	Body struct {
		shared.MessageResponse
		Accounts []account.Account `json:"accounts"`
		shared.PaginationResponse
	}
}

func (h *httpHandler) getAll(ctx context.Context, input *shared.PaginationRequest) (*GetAllAccountOutput, error) {
	LIMIT := input.Limit + 1

	accounts, err := h.accountService.GetAll(ctx, LIMIT, input.Cursor)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Accounts not found")
		default:
			h.logger.Error("failed to fetch accounts", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the accounts")
		}
	}

	resp := &GetAllAccountOutput{}
	resp.Body.Message = "Accounts fetched successfully"
	resp.Body.Accounts = accounts

	if len(accounts) == LIMIT {
		resp.Body.NextCursor = &accounts[len(accounts)-1].ID
		resp.Body.HasMore = true
		resp.Body.Accounts = resp.Body.Accounts[:len(resp.Body.Accounts)-1]
	}

	return resp, nil
}

type CreateAccountInput struct {
	Body account.CreateAccountParams `json:"account"`
}

func (h *httpHandler) create(ctx context.Context, input *CreateAccountInput) (*SingleAccountResponse, error) {
	if user := shared.GetAuthenticatedUser(ctx); user.ID != input.Body.UserID {
		h.logger.Error("input user id does not match authenticated user id",
			zap.Any("input user id", input.Body.UserID),
			zap.Any("authenticated user id", user.ID))

		return nil, huma.Error403Forbidden("Cannot create account for another user")
	}

	account, err := h.accountService.Create(ctx, input.Body)
	if err != nil {
		h.logger.Error("failed to create account", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while creating the account")
	}

	resp := &SingleAccountResponse{}
	resp.Body.Message = "Account created successfully"
	resp.Body.Account = &account

	return resp, nil
}

type UpdateAccountInput struct {
	shared.PathIDParam
	Body account.UpdateAccountParams `json:"account"`
}

func (h *httpHandler) update(ctx context.Context, input *UpdateAccountInput) (*SingleAccountResponse, error) {
	_, err := h.accountService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Account not found")
		default:
			h.logger.Error("failed to fetch account", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the account")
		}
	}

	account, err := h.accountService.Update(ctx, input.ID, input.Body)

	if err != nil {
		h.logger.Error("failed to update account", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while updating the account")
	}

	resp := &SingleAccountResponse{}
	resp.Body.Message = "Account updated successfully"
	resp.Body.Account = &account

	return resp, nil
}

type DeleteAccountResponse struct {
	Body shared.MessageResponse
}

func (h *httpHandler) delete(ctx context.Context, input *shared.PathIDParam) (*DeleteAccountResponse, error) {
	_, err := h.accountService.GetById(ctx, input.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, huma.Error404NotFound("Account not found")
		default:
			h.logger.Error("failed to fetch account", zap.Error(err))
			return nil, huma.Error500InternalServerError("An error occurred while fetching the account")
		}
	}

	err = h.accountService.Delete(ctx, input.ID)
	if err != nil {
		h.logger.Error("failed to delete account", zap.Error(err))
		return nil, huma.Error500InternalServerError("An error occurred while deleting the account")
	}

	resp := &DeleteAccountResponse{}
	resp.Body.Message = "Account deleted successfully"

	return resp, nil
}
