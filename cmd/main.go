package main

import (
	"fmt"
	"grcpValidatorIPv4/internal/base"
	"grcpValidatorIPv4/internal/controllers"
	"grcpValidatorIPv4/validator"
)

func main() {
	chErr := make(chan error, 1)
	app := new(base.App)
	app.Init()

	appController := new(controllers.AppController)
	app.Router.HandleFunc("/ip", appController.RequestToGRPCServer).Methods("POST")

	s := validator.GRPCServer{}

	go app.ApiRun("9000", chErr)
	go validator.ListenAndServe(&s, chErr)

	if err := <-chErr; err != nil {
		fmt.Println(err)
	}
}
