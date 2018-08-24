package handler

// import (
// 	"net/http"

// 	"github.com/FelipeUmpierre/measures/pkg/domain"
// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/render"
// )

// type (
// 	findAllUserRepository interface {
// 		FindAll() (*[]domain.User, error)
// 	}

// 	getUserRepository interface {
// 		FindByID(ID string) (*domain.User, error)
// 	}
// )

// // AllUsers return the users saved
// func AllUsers(repo findAllUserRepository) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		users, err := repo.FindAll()
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			render.JSON(w, r, struct {
// 				Error string `json:"error"`
// 			}{
// 				err.Error(),
// 			})

// 			return
// 		}

// 		render.JSON(w, r, users)
// 	}
// }

// // GetUser return the users saved
// func GetUser(repo getUserRepository) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id := chi.URLParam(r, `id`)

// 		user, err := repo.FindByID(id)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			render.JSON(w, r, struct {
// 				Error string `json:"error"`
// 			}{
// 				err.Error(),
// 			})

// 			return
// 		}

// 		render.JSON(w, r, user)
// 	}
// }
