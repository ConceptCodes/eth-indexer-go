package middleware

import (
	"net/http"
	"strings"
)

func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			w.Header().Set("Content-Type", "application/json;charset=utf8")
		}
		next.ServeHTTP(w, r)
	})
}
