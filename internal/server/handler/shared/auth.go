package shared

import (
	"context"

	// "proj/internal/entities/account"

	"proj/internal/entities/account"

	"github.com/supabase-community/gotrue-go/types"
)

const (
	UserContextKey    = "user"
	AccountContextKey = "account"
)

func GetAuthenticatedUser(ctx context.Context) types.User {
	if ctxValue, ok := ctx.Value(UserContextKey).(types.User); ok {
		return ctxValue
	}

	return types.User{}
}

func GetAuthenticatedAccount(ctx context.Context) account.Account {
	if ctxValue, ok := ctx.Value(AccountContextKey).(account.Account); ok {
		return ctxValue
	}

	return account.Account{}
}
