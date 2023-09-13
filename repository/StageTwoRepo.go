package repository

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"github.com/jackc/pgx/v4"
)



func PersonList() []*Person {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	allPerson := []*Person{}
	rows, _ := conn.Query(context.Background(), "select * from persons")

	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil
		}

		allPerson = append(allPerson, &Person{id,name})
	}

	return allPerson
}


func DbNewperson(person *Person) (string, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	person.ID = fmt.Sprintf("%d", rand.Intn(10000000)+10)
	
	_, err = conn.Exec(context.Background(), "insert into persons(id, name) values($1, $2) ", person.ID, person.Name)

	if err != nil {
		return "Can't save", err 
	}

	defer conn.Close(context.Background())

	
	//Persons = append(Persons, person)
   
	return person.ID, nil
}

func DbGetPerson(param string) (*Person, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
		
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	
	defer conn.Close(context.Background())

	a := &Person{}
	
	err = conn.QueryRow(context.Background(), "select * from persons where id=$1 or name=$1", param).Scan(&a.ID, &a.Name)
	
	if err != nil {
		return nil, errors.New("Person not found")
	}
	return a, nil
}

// func DbGetPersonBySlug(slug string) (*Person, error) {
// 	for _, a := range Persons {
// 		if a.Name == slug {
// 			return a, nil
// 		}
// 	}
// 	return nil, errors.New("Person not found")
// }

func DbUpdatePerson(param string, Person *Person) (*Person, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "update persons set name=$1 where id=$2 or name=$2", Person.Name, param)
	
	if err != nil {
		return nil, errors.New("Person not found")
	}
	return Person, nil
}

func DbRemovePerson(param string) (string, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "delete from persons where id=$1 or name=$1", param)

	if err != nil {
		return "Unable to delete person", errors.New("Person not found")
	}
	
	return param, nil
}

type Person struct {
	ID     string `db:"id"`
	Name   string `db:"name"`
}


// User data model
// type User struct {
// 	ID   int64  `json:"id"`
// 	Name string `json:"name"`
// }

// Person data model. I suggest looking at https://upper.io for an easy
// and powerful data persistence adapter.




