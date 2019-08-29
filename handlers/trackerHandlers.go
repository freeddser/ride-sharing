package handlers

import (
	"github.com/freeddser/ride-sharing/models"
	"github.com/freeddser/ride-sharing/services"
	"github.com/freeddser/rs-common/util"
	"net/http"
)

func TrackerDriver(w http.ResponseWriter, r *http.Request) {
	// parse json body
	params := models.Driver{}
	err := BindJSON(r, &params)
	if err != nil {
		log.Error(err)
		ToJSON(w, http.StatusOK, &models.GeneraResponse{Response: err.Error()})
		return
	}

	// create order
	res, err := services.TrackerDriver(params)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}

func TrackerOrder(w http.ResponseWriter, r *http.Request) {
	// parse json body
	params := models.Driver{}
	err := BindJSON(r, &params)
	if err != nil {
		log.Error(err)
		ToJSON(w, http.StatusOK, &models.GeneraResponse{Response: err.Error()})
		return
	}

	// create order
	res, err := services.TrackerOrder(params)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}

func GetTrackerByOrderID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	// create order
	res, err := services.GetTrackerByOrderID(util.StringToInt64(ID))
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}
