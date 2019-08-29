package handlers

import (
	"github.com/freeddser/ride-sharing/models"
	"github.com/freeddser/ride-sharing/services"
	"github.com/freeddser/rs-common/util"
	"net/http"
)

func GetDispacherOrderListByDriverID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	res, err := services.GetDispacherOrderListByDriverID(util.StringToInt64(ID))
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}

func DispacherOrder(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	res, err := services.DispacherOrder(util.StringToInt64(ID))
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}

func UpdateDispacherOrderStatus(w http.ResponseWriter, r *http.Request) {
	params := models.DispacherOrderStatus{}
	err := BindJSON(r, &params)
	if err != nil {
		log.Error(err)
		ToJSON(w, http.StatusOK, &models.GeneraResponse{Response: err.Error()})
		return
	}
	res, err := services.ModifyDispacherOrderStatus(params)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}
