package app

import (
	"log"
	"net/http"

	"github.com/kevindurb/done/html"
	"golang.org/x/crypto/bcrypt"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type loginBody struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func (a *App) loginShow(w http.ResponseWriter, r *http.Request) (g.Node, error) {
	return html.Layout(
		h.H1(g.Text("Login")),
	), nil
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {
	var body loginBody
	if err := a.fp.Parse(&body, r); err != nil {
		log.Printf("Error parsing body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := a.q.GetUserByEmail(r.Context(), body.Email)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = bcrypt.CompareHashAndPassword(user.Hash, []byte(body.Password)); err != nil {
		log.Printf("Error comparing password hash: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = a.setUserID(w, r, user.ID); err != nil {
		log.Printf("Error setting session user id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
