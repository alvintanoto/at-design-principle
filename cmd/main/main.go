package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"alvintanoto.id/design-principle/internal/controller"
	"github.com/gorilla/mux"
)

type Application struct {
	Configurations *Configurations

	Controller controller.Controller

	Router *mux.Router
}

func (app *Application) initConfigs() {
	configs := &Configurations{}
	if err := configs.ReadConfigs(); err != nil {
		log.Println("failed to read configuration: ", err.Error())
		return
	}

	app.Configurations = configs
}

func (app *Application) initController() {
	app.Controller = controller.NewController()
}

func main() {
	app := &Application{}
	app.initConfigs()
	app.initController()
	app.registerRoutes()

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", app.Configurations.Server.Port),
		Handler:      app.Router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		log.Println("starting server on", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("server shutting down")

	os.Exit(0)
}
