package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
)

type userIDKey struct{}

func NewAuthMiddleware(client *auth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				next.ServeHTTP(w, r)
				return
			}

			tElems := strings.SplitN(token, " ", 2)
			if len(tElems) < 2 {
				http.Error(w, `{"reason": "invalid token format"}`, http.StatusUnauthorized)
				return
			}

			if tElems[0] != "Bearer" {
				http.Error(w, `{"reason": "invalid token format"}`, http.StatusUnauthorized)
				return
			}

			idToken := tElems[1]
			user, err := client.VerifyIDToken(r.Context(), idToken)
			if err != nil {
				log.Println(err)
				http.Error(w, `{"reason": "invalid token"}`, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey{}, user.UID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey{}).(string)
	return userID, ok
}
