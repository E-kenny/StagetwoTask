package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"stagaTwoCrud/stageTwoHandlers"

	"github.com/go-chi/docgen"
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
		r.With(stageTwoHandlers.Paginate).Get("/", stageTwoHandlers.Listpersons)
		r.Post("/", stageTwoHandlers.Createperson)  

		r.Route("/{param}", func(r chi.Router) {
			r.Use(stageTwoHandlers.PersonCtx)      
			r.Get("/", stageTwoHandlers.Getperson)   
			r.Put("/", stageTwoHandlers.Updateperson)    
			r.Delete("/", stageTwoHandlers.Deleteperson) 
		})

	})

	
		// fmt.Println(docgen.JSONRoutesDoc(r))
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/E-kenny/StagetwoTask",
			Intro:       "Welcome to the task two/rest generated docs.",
		}))
		return
	

	http.ListenAndServe(":3333", r)
}




