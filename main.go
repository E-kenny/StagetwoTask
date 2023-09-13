package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	"stagaTwoCrud/stageTwoHandlers"
	//"github.com/go-chi/docgen"
	// "fmt"
)

func main() {
	// os.Setenv("DATABASE_URL", "postgres://postgres:E_kenny246810@localhost:5432/StageTwoDB")
	// os.Setenv("DATABASE_URL", "postgres://persons_414z_user:NXghusPlovh5g49DGkBx84LwrU4h5df2@dpg-ck0c3mj6fquc73ch4log-a/persons_414z")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	  }
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	var testTable = `
        CREATE TABLE IF NOT EXISTS persons(
                id TEXT,
				name TEXT
        )
  `
	_, err = conn.Exec(context.Background(), testTable)

	if err != nil {
		fmt.Println(err)
	}

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

	http.ListenAndServe(":3333", r)
}




