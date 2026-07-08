package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kevindurb/done/app"
	"github.com/kevindurb/done/config"
)

func main() {
	c := config.FromEnv()
	a := app.New(c)
	log.Printf("Listening: %s", c.ListenAddr)
	http.ListenAndServe(c.ListenAddr, a.Router())
}
