package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	//"github.com/go-chi/docgen"
	"stagaTwoCrud/stageTwoHandlers"
)

func main() {
	os.Setenv("DATABASE_URL", "postgres://postgres:E_kenny246810@localhost:5432/StageTwoDB")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// RESTy routes for "persons" resource
	r.Route("/api", func(r chi.Router) {
		r.With(stageTwoHandlers.Paginate).Get("/", stageTwoHandlers.Listpersons) // GET /persons
		r.Post("/", stageTwoHandlers.Createperson)       // POST /persons

		r.Route("/{param}", func(r chi.Router) {
			r.Use(stageTwoHandlers.PersonCtx)            // Load the *person on the request context
			r.Get("/", stageTwoHandlers.Getperson)       // GET /persons/123
			r.Put("/", stageTwoHandlers.Updateperson)    // PUT /persons/123
			r.Delete("/", stageTwoHandlers.Deleteperson) // DELETE /persons/123
		})

		// GET /persons/whats-up
		r.With(stageTwoHandlers.PersonCtx).Get("/{personSlug:[a-z-]+}", stageTwoHandlers.Getperson)
	})

	http.ListenAndServe(":3333", r)
}




