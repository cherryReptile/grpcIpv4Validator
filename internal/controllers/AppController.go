package controllers

import (
	"encoding/json"
	"errors"
	"grcpValidatorIPv4/api"
	"grcpValidatorIPv4/validator"
	"net/http"
)

type AppController struct {
	BaseController
}

func (c *AppController) RequestToGRPCServer(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Ip string `json:"ip"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Error(w, http.StatusBadRequest, err)
		return
	}

	if request.Ip == "" {
		c.Error(w, http.StatusBadRequest, errors.New("field ip is required"))
		return
	}

	res, err := validator.ClientDo(&api.ValRequest{Ipv4: request.Ip})
	if err != nil {
		c.Error(w, http.StatusBadRequest, err)
		return
	}

	c.JSON(w, http.StatusOK, map[string]any{"status": http.StatusText(http.StatusOK), "validIP": res.Response})
}
