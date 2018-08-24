package main

import (
	"net/http"
	"os"

	"github.com/FelipeUmpierre/measures/pkg/handler"
	"github.com/FelipeUmpierre/measures/pkg/middleware"
	"github.com/FelipeUmpierre/measures/pkg/repository"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := sqlx.Connect(`postgres`, os.Getenv(`DB_DSN`))
	failOnError(err, `Unable to connect with database`)

	measureRepo := repository.NewMeasuresRepository(db)

	// usersRepo := repository.NewUsersRepository(db)

	r := chi.NewRouter()
	// r.Route(`/user`, func(r chi.Router) {
	// 	r.Get(`/all`, handler.AllUsers(usersRepo))
	// 	r.Get(`/{id}`, handler.GetUser(usersRepo))
	// })

	r.Route(`/measures`, func(r chi.Router) {
		// command
		r.With(middleware.CollectUserID).
			Post(`/`, handler.SaveMeasures(measureRepo))

		// query
		r.Get(`/{aggregateID}`, handler.GetMeasureState(measureRepo))
		r.Get(`/{aggregateID}/all`, handler.GetMeasures(measureRepo))
	})

	http.ListenAndServe(`:8000`, r)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.WithError(err).Fatal(msg)
	}
}
