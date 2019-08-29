package handlers

import (
	"encoding/json"
	"github.com/freeddser/rs-common/logging"
	"net/http"
)

const jsonContentType = "application/json; charset=utf-8"

var log = logging.MustGetLogger()

func BindJSON(req *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(obj); err != nil {
		return err
	}

	return nil
}

func ToJSON(w http.ResponseWriter, statusCode int, obj interface{}) error {
	w.Header().Add("Content-Type", jsonContentType)
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(obj)
}
