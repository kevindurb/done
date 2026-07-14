package app

import (
	"log"
	"net/http"

	"github.com/kevindurb/done/html/pages"
	"github.com/kevindurb/done/sqlcgen"
	"golang.org/x/crypto/bcrypt"
)

type signupBody struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func (a *App) signupShow(w http.ResponseWriter, r *http.Request) {
	pages.Signup().Render(w)
}

func (a *App) signup(w http.ResponseWriter, r *http.Request) {
	var body signupBody
	if err := a.fp.Parse(&body, r); err != nil {
		log.Printf("Error parsing body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = a.q.CreateUser(r.Context(), sqlcgen.CreateUserParams{
		Email: body.Email,
		Hash:  hash,
	})
	if err != nil {
		log.Printf("Error creating user: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
