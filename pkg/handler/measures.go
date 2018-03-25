package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

func GetMeasures(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
