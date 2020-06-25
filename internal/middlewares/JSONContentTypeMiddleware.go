package middlewares

import "net/http"

// JSONContentTypeMiddleware will add the header 'content-type: application/json; charset=UTF-8' for all endpoints
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json; charset=UTF-8")
		next.ServeHTTP(w, r)
	})
}
