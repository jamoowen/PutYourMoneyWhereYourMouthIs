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

func GetUserFromCtx(ctx context.Context) User {
	return ctx.Value("user").(User)
}
