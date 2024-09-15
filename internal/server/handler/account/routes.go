package account

import (
	"net/http"

	"proj/internal/server/middleware"
	"proj/internal/service"

	"github.com/danielgtaylor/huma/v2"
	"github.com/supabase-community/supabase-go"
	"go.uber.org/zap"
)

func RegisterHumaRoutes(
	accountService service.AccountService,
	humaApi huma.API,
	logger *zap.Logger,
	supabaseClient *supabase.Client,
) {

	handler := &httpHandler{
		accountService: accountService,
		logger:         logger,
	}

	huma.Register(humaApi, huma.Operation{
		OperationID: "get-account-by-id",
		Method:      http.MethodGet,
		Path:        "/account/{id}",
		Summary:     "Get account by ID",
		Description: "Get account by ID.",
		Tags:        []string{"Account"},
	}, handler.getByID)

	huma.Register(humaApi, huma.Operation{
		OperationID: "get-account-by-user-id",
		Method:      http.MethodGet,
		Path:        "/account/user/{userId}",
		Summary:     "Get account by user ID",
		Description: "Get account by user ID.",
		Tags:        []string{"Account"},
		Security: []map[string][]string{
			{"bearerAuth": {}},
		},
		Middlewares: huma.Middlewares{
			func(ctx huma.Context, next func(huma.Context)) {
				middleware.WithUser(humaApi)(ctx, next, logger, supabaseClient)
			},
		},
	}, handler.getByUserId)

	huma.Register(humaApi, huma.Operation{
		OperationID: "get-all-accounts",
		Method:      http.MethodGet,
		Path:        "/account",
		Summary:     "Get all accounts",
		Description: "Get all accounts.",
		Tags:        []string{"Account"},
	}, handler.getAll)

	huma.Register(humaApi, huma.Operation{
		OperationID: "create-account",
		Method:      http.MethodPost,
		Path:        "/account",
		Summary:     "Create a account",
		Description: "Create a account.",
		Tags:        []string{"Account"},
		Security: []map[string][]string{
			{"bearerAuth": {}},
		},
		Middlewares: huma.Middlewares{
			func(ctx huma.Context, next func(huma.Context)) {
				middleware.WithUser(humaApi)(ctx, next, logger, supabaseClient)
			},
		},
	}, handler.create)

	huma.Register(humaApi, huma.Operation{
		OperationID: "update-account",
		Method:      http.MethodPut,
		Path:        "/account/{id}",
		Summary:     "Update a account",
		Description: "Update a account.",
		Tags:        []string{"Account"},
	}, handler.update)

	huma.Register(humaApi, huma.Operation{
		OperationID: "delete-account",
		Method:      http.MethodDelete,
		Path:        "/account/{id}",
		Summary:     "Delete a account",
		Description: "Delete a account.",
		Tags:        []string{"Account"},
	}, handler.delete)

}
