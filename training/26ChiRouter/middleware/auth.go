package middleware

import (
	"context"
	"fmt"
	"net/http"
)

func authMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth Middleware is running ...")
		ctx := context.WithValue(r.Context(), "user", "Admin") // this is for storing user information in the context
		r = r.WithContext(ctx)
		fmt.Println("Auth Middleware is finished ...")
		next.ServeHTTP(w, r) // next.ServeHTTP(w, r) is used to pass the request to the next handler in the chain

	})
}