package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi"
	"github.com/jamoowen/PutYourMoneyWhereYourMouthIs/services/pymwymi/services/auth"
)

// extracts page and limit from query params and adds to ctx
// page defaults to 1, limit defaults to 20
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
		ctx := context.WithValue(r.Context(), pymwymi.PaginationKey, pymwymi.PageOpts{Page: page, Limit: limit})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func authMiddleware(authService *auth.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("pymwymi_auth_token")
			if err != nil {
				http.Error(w, "Unauthorized: missing auth cookie", http.StatusUnauthorized)
				return
			}

			user, authErr := authService.AuthenticateUserToken(cookie.Value)
			if authErr != nil {
				http.Error(w, authErr.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), pymwymi.UserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
