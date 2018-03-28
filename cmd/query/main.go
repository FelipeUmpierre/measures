package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get(`/`, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(`another service`)
	})

	http.ListenAndServe(`:8001`, r)
}
