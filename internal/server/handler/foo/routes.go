package foo

import (
	"net/http"

	"proj/internal/service"

	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/zap"
)

func RegisterHumaRoutes(
	fooService service.FooService,
	humaApi huma.API,
	logger *zap.Logger,
) {

	handler := &httpHandler{
		fooService: fooService,
		logger:     logger,
	}

	huma.Register(humaApi, huma.Operation{
		OperationID: "get-foo-by-id",
		Method:      http.MethodGet,
		Path:        "/foo/{id}",
		Summary:     "Get foo by ID",
		Description: "Get foo by ID.",
		Tags:        []string{"Foo"},
	}, handler.getByID)

	huma.Register(humaApi, huma.Operation{
		OperationID: "get-all-foos",
		Method:      http.MethodGet,
		Path:        "/foo",
		Summary:     "Get all foos",
		Description: "Get all foos.",
		Tags:        []string{"Foo"},
	}, handler.getAll)

	huma.Register(humaApi, huma.Operation{
		OperationID: "create-foo",
		Method:      http.MethodPost,
		Path:        "/foo",
		Summary:     "Create a foo",
		Description: "Create a foo.",
		Tags:        []string{"Foo"},
	}, handler.create)

	huma.Register(humaApi, huma.Operation{
		OperationID: "update-foo",
		Method:      http.MethodPut,
		Path:        "/foo/{id}",
		Summary:     "Update a foo",
		Description: "Update a foo.",
		Tags:        []string{"Foo"},
	}, handler.update)

	huma.Register(humaApi, huma.Operation{
		OperationID: "delete-foo",
		Method:      http.MethodDelete,
		Path:        "/foo/{id}",
		Summary:     "Delete a foo",
		Description: "Delete a foo.",
		Tags:        []string{"Foo"},
	}, handler.delete)

}
