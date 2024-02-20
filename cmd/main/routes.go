package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) registerRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/", app.Controller.ViewController.HomepageHandler())

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets/"))))

	app.Router = router
}
