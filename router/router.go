package router

import (
	"github.com/freeddser/ride-sharing/handlers"
	"github.com/freeddser/rs-common/middleware"
	"github.com/gorilla/mux"
)

var recoverHandler = middleware.RecoverHandler()

var log = middleware.LogerClient()
var timer = middleware.Timer()

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", middleware.Chain(log(handlers.GetHomePage), timer, recoverHandler))

	// order Service
	r.HandleFunc("/orders/create", middleware.Chain(log(handlers.CreateOrder), timer, recoverHandler))
	r.HandleFunc("/orders", middleware.Chain(log(handlers.GetOrderByID), timer, recoverHandler))
	r.HandleFunc("/orders/edit", middleware.Chain(log(handlers.UpdateOrder), timer, recoverHandler))
	r.HandleFunc("/orders/status", middleware.Chain(log(handlers.UpdateOrderDispacherStatus), timer, recoverHandler))
	r.HandleFunc("/orders/gettripbyorderid", middleware.Chain(log(handlers.GetTripByOrderID), timer, recoverHandler))

	// dispacher Service
	// Dispacher order by orderID
	r.HandleFunc("/dispacher", middleware.Chain(log(handlers.DispacherOrder), timer, recoverHandler))
	// get dispacher order list by driverID
	r.HandleFunc("/dispacher/getorders", middleware.Chain(log(handlers.GetDispacherOrderListByDriverID), timer, recoverHandler))
	// update dispacher order status by orderID,DriverID
	r.HandleFunc("/dispacher/updateorder", middleware.Chain(log(handlers.UpdateDispacherOrderStatus), timer, recoverHandler))

	// Tracker Service
	r.HandleFunc("/tracker/driver", middleware.Chain(log(handlers.TrackerDriver), timer, recoverHandler))
	r.HandleFunc("/tracker/order/driver", middleware.Chain(log(handlers.TrackerOrder), timer, recoverHandler))
	r.HandleFunc("/tracker/order", middleware.Chain(log(handlers.GetTrackerByOrderID), timer, recoverHandler))

	// Passenger Service todo CURD
	r.HandleFunc("/passenger", middleware.Chain(log(handlers.GetPassengerByID), timer, recoverHandler))

	// Driver Service todo CURD
	r.HandleFunc("/driver", middleware.Chain(log(handlers.GetDriverByID), timer, recoverHandler))

	return r
}
