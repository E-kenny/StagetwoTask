package main

import (
	//"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	//"github.com/jackc/pgx/v4"

	"stagaTwoCrud/stageTwoHandlers"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
    "database/sql"
	//"github.com/go-chi/docgen"
	// "fmt"
)

var connGLOBAL *pgxpool.Pool

func main() {
	os.Setenv("DATABASE_URL", "postgres://postgres:E_kenny246810@localhost:5432/StageTwoDB")
	// os.Setenv("DATABASE_URL", "postgres://persons_414z_user:NXghusPlovh5g49DGkBx84LwrU4h5df2@dpg-ck0c3mj6fquc73ch4log-a.oregon-postgres.render.com/persons_414z")

	// conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
	// 	os.Exit(1)
	// }
	
	// _, err = conn.Exec(context.Background(), "create table personNew $1, $2", "id", "name")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    "https://dashboard.render.com/d/dpg-ck0c3mj6fquc73ch4log-a", 5432, "persons_414z_user", "persons_414z_user", "persons_414z")

  db, err := sql.Open("postgres", psqlInfo)

  if err != nil {
    panic(err)
  }

  err = db.Ping()

  if err != nil {
    panic(err)
  }


  defer db.Close()
  fmt.Println("Successfully connected!")
  var testTable = `
        CREATE TABLE IF NOT EXISTS persons(
                id Text
				name Text
        )
  `
  dbret, _ := db.Exec(testTable)
  fmt.Printf("value of db %v", dbret)
  

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
		// fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
		// 	ProjectPath: "github.com/E-kenny/StagetwoTask",
		// 	Intro:       "Welcome to the task two/rest generated docs.",
		// }))
		// return
	

	http.ListenAndServe(":3333", r)
}




