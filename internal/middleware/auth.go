package middleware

import (
	"net/http"

	"github.com/nurbekabilev/go-open-api/internal/app"
	"github.com/nurbekabilev/go-open-api/internal/handler/response"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		di := app.DI()
		token := r.Header.Get("Authorization")

		err := di.Auth.CheckAuth(token)
		if err != nil {
			response.WriteJsonResponse(w, response.NewUnauthroziedResponse("Unauthorized: invalid token"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
