package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/FelipeUmpierre/measures/pkg/business/command"
	"github.com/go-chi/render"
)

func ValidateUserIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get(`user`)
		if userID == `` {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, struct {
				Message string `json:"message"`
			}{`Missing user ID`})
			return
		}

		ctx := context.WithValue(r.Context(), `userID`, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ValidateMeasurePayloadMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		measure := new(command.MeasureCommand)
		if err := json.NewDecoder(r.Body).Decode(measure); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, struct {
				Message string `json:"message"`
				Error   error  `json:"error"`
			}{
				`Incorrect payload body`,
				err,
			})
			return
		}

		ctx := context.WithValue(r.Context(), `measure`, measure)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
