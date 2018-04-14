package main

import (
	"net/http"
	"os"

	"github.com/FelipeUmpierre/measures/pkg/application/bus"
	"github.com/FelipeUmpierre/measures/pkg/application/handler"
	"github.com/FelipeUmpierre/measures/pkg/business/command"
	"github.com/FelipeUmpierre/measures/pkg/business/handler"
	mid "github.com/FelipeUmpierre/measures/pkg/middleware"
	"github.com/FelipeUmpierre/measures/pkg/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := sqlx.Connect(`postgres`, os.Getenv(`DB_DSN`))
	failOnError(err, `Unable to connect with database`)

	usersRepo := repository.NewUsersRepository(db)
	eventRepo := repository.NewEventRepository(db)

	dispatcher := bus.New()
	dispatcher.Register(command.MeasureCommand{}, commandhandler.NewMeasureHandler(eventRepo))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route(`/users`, func(r chi.Router) {
		r.Get(`/`, handler.AllUsers(usersRepo))
		r.Get(`/{id}`, handler.GetUser(usersRepo))
	})

	r.Route(`/measure`, func(r chi.Router) {
		r.Use(
			mid.ValidateUserIDMiddleware,
			mid.ValidateMeasurePayloadMiddleware,
		)
		r.Post(`/`, handler.SaveMeasure(dispatcher))
	})

	http.ListenAndServe(`:8000`, r)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.WithError(err).Fatal(msg)
	}
}
