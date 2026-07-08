package app

import (
	"errors"
	"log"
	"net/http"
)

func (a *App) setUserID(w http.ResponseWriter, r *http.Request, id int64) error {
	log.Printf("Setting user_id in session: %d", id)
	s, err := a.s.Get(r, "done-session")
	if err != nil {
		return err
	}
	s.Values["user_id"] = id
	return a.s.Save(r, w, s)
}

func (a *App) userID(r *http.Request) (int64, error) {
	s, err := a.s.Get(r, "done-session")
	if err != nil {
		return 0, err
	}

	id, ok := s.Values["user_id"].(int64)
	if !ok {
		return 0, errors.New("failed getting user_id from session")
	}
	log.Printf("Got user_id in session: %d", id)
	return id, nil
}

func (a *App) isLoggedIn(r *http.Request) bool {
	id, err := a.userID(r)
	return err == nil && id > 0
}

func (a *App) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !a.isLoggedIn(r) {
			log.Printf("Logged out, redirecting to /login")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		log.Printf("Logged in, continuing")
		next.ServeHTTP(w, r)
	})
}
