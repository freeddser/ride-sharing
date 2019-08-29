package handlers

import (
	"github.com/freeddser/ride-sharing/services"
	"github.com/freeddser/rs-common/util"
	"net/http"
)

func GetDriverByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	res, err := services.GetDriverByID(util.StringToInt64(ID))
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ToJSON(w, http.StatusOK, res)
}
