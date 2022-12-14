package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type BaseController struct{}

func (c *BaseController) Error(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status": http.StatusText(code),
		"error":  err.Error(),
	})

	if err != nil {
		log.Printf("[ERROR] Error response error: %v", err)
		return
	}

	log.Printf("[INFO] Error response: %v", err)
}

func (c *BaseController) JSON(w http.ResponseWriter, code int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Printf("[ERROR] Error response json encoding: %v", err)
		return
	}

	log.Printf("[INFO] JSON response: %v", response)
}
