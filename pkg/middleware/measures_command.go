package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
)

// CollectUserID ...
func CollectUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get(`User-ID`)
		if userID == `` {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, `missing user id`)
			return
		}

		ctx := context.WithValue(r.Context(), `aggregate_id`, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
