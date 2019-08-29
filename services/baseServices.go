package services

import (
	"github.com/freeddser/rs-common/logging"
	"github.com/freeddser/rs-common/util"
)

const (
	BASE_URL  = "baseUrl"
	CLIENT_ID = "clientId"
)

var Props map[string]string
var log = logging.MustGetLogger()

var BadRequest *GenericResponse
var RequestNotFound *GenericResponse

func InitializeService(name string) error {
	initErrorResponse()
	return nil
}

func initErrorResponse() {
	BadRequest = NewGenericResponse(-1001)
	RequestNotFound = NewGenericResponse(-1002)
}

func NewGenericResponse(errorId int) *GenericResponse {

	messageEn := util.GetErrorString(util.EN, errorId)
	messageCn := util.GetErrorString(util.CN, errorId)

	return &GenericResponse{
		//if errorID >=0, then true
		Success: errorId >= 0,
		Error: ErrorContext{
			ErrorId: errorId,
			Message: Message{
				CN: messageCn,
				EN: messageEn,
			},
		},
	}
}

type GenericResponse struct {
	Success bool         `json:"success"`
	Error   ErrorContext `json:"error"`
	Data    interface{}  `json:"results,omitempty"`
}

type ErrorContext struct {
	ErrorId int     `json:"error_id"`
	Message Message `json:"message"`
}

type Message struct {
	CN string `json:"cn"`
	EN string `json:"en"`
}

type ServiceCallback interface {
	checkParam()
	process()
}
