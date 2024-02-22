package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) registerRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/", app.Controller.ViewController.HomepageHandler())
	router.HandleFunc("/color", app.Controller.ViewController.ColorpageHandler())
	router.HandleFunc("/font", app.Controller.ViewController.FontpageHandler())
	router.HandleFunc("/button", app.Controller.ViewController.ButtonpageHandler())
	router.HandleFunc("/chart", app.Controller.ViewController.ChartpageHandler())
	router.HandleFunc("/alert", app.Controller.ViewController.AlertPageHandler())

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets/"))))

	app.Router = router
}
