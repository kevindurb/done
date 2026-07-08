package app

import (
	"log"
	"net/http"

	"github.com/kevindurb/done/html"
	"github.com/kevindurb/done/sqlcgen"
	"golang.org/x/crypto/bcrypt"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type signupBody struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func (a *App) signupShow(w http.ResponseWriter, r *http.Request) (g.Node, error) {
	return html.Layout(
		h.H1(g.Text("Signup")),
		h.Form(
			h.Method(http.MethodPost),
			h.Action("/signup"),
			h.Div(
				h.Label(h.For("email"), g.Text("Email")),
				h.Input(h.Type("email"), h.ID("email"), h.Name("Email")),
			),
			h.Div(
				h.Label(h.For("password"), g.Text("Password")),
				h.Input(h.Type("password"), h.ID("password"), h.Name("Password")),
			),
			h.Button(h.Type("submit"), g.Text("Signup")),
			h.A(h.Href("/login"), g.Text("Login")),
		),
	), nil
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
