package pymwymi

import "context"

type PageOpts struct {
	Page  int64
	Limit int64
}

type User struct {
	Name          string
	WalletAddress string
}

type FieldToSet struct {
	Field string
	Value string
}

type ctxKey string

const (
	PaginationKey ctxKey = "pagination"
	UserKey       ctxKey = "user"
)

func GetUserFromCtx(ctx context.Context) User {
	return ctx.Value(UserKey).(User)
}

func GetPageOptsFromCtx(ctx context.Context) PageOpts {
	return ctx.Value(PaginationKey).(PageOpts)
}
