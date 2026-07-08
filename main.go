package main

import (
	"net/http"

	"github.com/kevindurb/done/app"
	"github.com/kevindurb/done/config"
)

func main() {
	c := config.FromEnv()

	a := app.New(&c)

	http.ListenAndServe(c.ListenAddr, a.Router())
}
