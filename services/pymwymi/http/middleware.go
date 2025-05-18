package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
)

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		page, err := strconv.ParseInt(q.Get("page"), 10, 64)
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := strconv.ParseInt(q.Get("limit"), 10, 64)
		if err != nil || limit <= 0 || limit > 100 {
			limit = 20
		}
		ctx := context.WithValue(r.Context(), "pagination", pymwymi.PageOpts{page, limit})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// so here we are checking the token has been signed by the intended wallet address?
		name := "test"
		walletAddress := "test"
		ctx := context.WithValue(r.Context(), "user", pymwymi.User{Name: name, WalletAddress: walletAddress})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
