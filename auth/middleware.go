package auth

import (
	"context"
	"net/http"
)

// identify the user when the WebSocket connection is established.
// to make this happen we will create a Middleware function that
// we can add to the WebSocket HTTP endpoint.

type contextKey string

const UserContextKey = contextKey("user")


func AuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, tok := r.URL.Query()["bearer"]

		if tok && len(token) == 1 {

			user, err := ValidateToken(token[0])
			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)

			} else {
				ctx := context.WithValue(r.Context(), UserContextKey, user)
				f(w, r.WithContext(ctx))
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("we dont know you... Please login first ;)"))
		}
	})
}
