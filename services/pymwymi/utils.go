package pymwymi

import (
	"context"
	"time"
)

type PageOpts struct {
	Page  int64
	Limit int64
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

func GetPageOptsFromCtx(ctx context.Context) *PageOpts {
	paginationKey := ctx.Value(PaginationKey)
	if paginationKey == nil {
		return nil
	}
	opts := paginationKey.(PageOpts)
	return &opts
}

func GetIsoTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

func IsoNow() string {
	return GetIsoTime(time.Now())
}
