package middleware

import (
	"net/http"
)

func setJSONNormailDataAtMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
