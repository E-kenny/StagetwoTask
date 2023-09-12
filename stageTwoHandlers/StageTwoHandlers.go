package stageTwoHandlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"stagaTwoCrud/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func Listpersons(w http.ResponseWriter, r *http.Request) {

	if err := render.RenderList(w, r, NewpersonListResponse(repository.PersonList())); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

}

func PersonCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var person *repository.Person
		var err error

		if param := chi.URLParam(r, "param"); param != "" {
			person, err = repository.DbGetPerson(param)

		// } else if personSlug := chi.URLParam(r, "personSlug"); personSlug != "" {
		// 	person, err = repository.DbGetPersonBySlug(personSlug)
		// 
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "person", person)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func Createperson(w http.ResponseWriter, r *http.Request) {
	data := &personRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	person := data.Person
	repository.DbNewperson(person)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewpersonResponse(person))
}

func Getperson(w http.ResponseWriter, r *http.Request) {

	person := r.Context().Value("person").(*repository.Person)

	if err := render.Render(w, r, NewpersonResponse(person)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// Updateperson updates an existing person in our persistent store.
func Updateperson(w http.ResponseWriter, r *http.Request) {
	person := r.Context().Value("person").(*repository.Person)

	data := &personRequest{Person: person}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	person = data.Person
	repository.DbUpdatePerson(person.ID, person)

	render.Render(w, r, NewpersonResponse(person))
}

// Deleteperson removes an existing person from our persistent store.
func Deleteperson(w http.ResponseWriter, r *http.Request) {
	var err error
	person := r.Context().Value("person").(*repository.Person)

	_ , err = repository.DbRemovePerson(person.ID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewpersonResponse(person))
}

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}

// This is entirely optional, but I wanted to demonstrate how you could easily
// add your own logic to the render.Respond method.
func init() {
	render.Respond = func(w http.ResponseWriter, r *http.Request, v interface{}) {
		if err, ok := v.(error); ok {

			// We set a default error status response code if one hasn't been set.
			if _, ok := r.Context().Value(render.StatusCtxKey).(int); !ok {
				w.WriteHeader(400)
			}

			// We log the error
			fmt.Printf("Logging err: %s\n", err.Error())

			render.DefaultResponder(w, r, render.M{"status": "error"})
			return
		}

		render.DefaultResponder(w, r, v)
	}
}

type personRequest struct {
	*repository.Person
}

func (a *personRequest) Bind(r *http.Request) error {
	// a.person is nil if no person fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	if a.Person == nil {
		return errors.New("missing required person fields")
	}
	
	return nil
}

func NewpersonResponse(person *repository.Person) *personResponse {
	resp := &personResponse{Person: person}

	return resp
}

type personResponse struct {
	*repository.Person	
}

func (rd *personResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewpersonListResponse(persons []*repository.Person) []render.Renderer {
	list := []render.Renderer{}
	for _, person := range persons {
		list = append(list, NewpersonResponse(person))
	}
	return list
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

// type UserPayload struct {
// 	*repository.User
// 	Role string `json:"role"`
// }

// func NewUserPayloadResponse(user *repository.User) *UserPayload {
// 	return &UserPayload{User: user}
// }