package main

import "github.com/gorilla/mux"

func (app *Application) registerRoutes() {
	router := mux.NewRouter()

	app.Router = router
}
