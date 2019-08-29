package handlers

import (
	"github.com/freeddser/ride-sharing/models"
	"github.com/freeddser/ride-sharing/services"
	"github.com/freeddser/rs-common/util"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	// parse json body
	params := models.Order{}
	err := BindJSON(r, &params)
	if err != nil {
		log.Error(err)
		ToJSON(w, http.StatusOK, &models.GeneraResponse{Response: err.Error()})
		return
	}

	// create order
	res, err := services.CreateOrder(params)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	res, err := services.GetOrderByID(util.StringToInt64(ID))
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}

func GetTripByOrderID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	res, err := services.GetTripByOrderID(util.StringToInt64(ID))
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	params := models.Order{}
	err := BindJSON(r, &params)
	if err != nil {
		log.Error(err)
		ToJSON(w, http.StatusOK, &models.GeneraResponse{Response: err.Error()})
		return
	}

	// create order
	res, err := services.UpdateOrder(params)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}

func UpdateOrderDispacherStatus(w http.ResponseWriter, r *http.Request) {
	params := models.OrderDispacherStatus{}
	err := BindJSON(r, &params)
	if err != nil {
		log.Error(err)
		ToJSON(w, http.StatusOK, &models.GeneraResponse{Response: err.Error()})
		return
	}
	// dispacher order
	res, err := services.UpdateOrderDispacherStatus(params)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}
