package middlewares

import "net/http"

type JSONContentTypeMiddleware struct{}

func NewJSONContentTypeMiddleware() *JSONContentTypeMiddleware {
	return &JSONContentTypeMiddleware{}
}

// JSONContentTypeMiddleware will add the header 'content-type: application/json; charset=UTF-8' for all endpoints
func (m *JSONContentTypeMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	// Add Content-Type header
	w.Header().Add("content-type", "application/json; charset=UTF-8")

	// Call the next middleware handler
	next(w, req)
}
