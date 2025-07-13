package http

import (
	"context"
	"net/http"
)

type contextKey string
const contextKeyCustomerID contextKey = "customerID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var customerID int64 = 1 // Contoh: ID Budi setelah login
		ctx := context.WithValue(r.Context(), contextKeyCustomerID, customerID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}