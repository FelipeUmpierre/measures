package handler

import (
	"net/http"

	"github.com/FelipeUmpierre/measures/pkg/domain"
)

type (
	getUserRepository interface {
		FindAll() ([]domain.User, error)
	}
)

// AllUsers return the users saved
func AllUsers(repo getUserRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
