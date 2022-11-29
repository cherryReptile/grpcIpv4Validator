package base

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type App struct {
	Router *mux.Router
	Server *http.Server
}

func (a *App) Init() {
	a.Router = mux.NewRouter()
}

func (a *App) ApiRun(port string, ch chan error) {
	a.Server = &http.Server{
		Handler:      a.Router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	defer a.Server.Close()

	log.Printf("[DEBUG] Running server on port %s", port)

	if err := a.Server.ListenAndServe(); err != nil {
		ch <- errors.New(fmt.Sprintf("Error server: %s", err.Error()))
	}
}

func (a *App) Close() {
	if err := a.Server.Close(); err != nil {
		log.Fatalf("[FATAL] Unable to close server: %v", err)
		return
	}
}
