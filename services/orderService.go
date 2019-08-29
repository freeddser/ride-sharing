package services

import (
	"fmt"
	"github.com/freeddser/rs-common/config"
	"github.com/freeddser/ride-sharing/flags"
	"github.com/freeddser/ride-sharing/models"
	"github.com/freeddser/ride-sharing/repository"
	"github.com/freeddser/rs-common/util"
	"net/url"
	"strconv"
	"time"
)

func CreateOrder(order models.Order) (*GenericResponse, error) {
	if order.CustomerID == 0 || order.PickupLongitude == 0 || order.PickupLatitude == 0 {
		return NewGenericResponse(-1001), nil
	}
	repo := repository.GetOrderRepository()
	data, err := repo.WriteOrder(order)
	if err != nil {
		return NewGenericResponse(-1000), nil
	}

	genResp := NewGenericResponse(0)
	(*genResp).Data = data
	return genResp, nil
}

func UpdateOrder(order models.Order) (*GenericResponse, error) {
	if order.OrderID == 0 {
		return NewGenericResponse(-1001), nil
	}

	repo := repository.GetOrderRepository()
	err := repo.UpdateOrder(order)
	if err != nil {
		return NewGenericResponse(-1000), nil
	}

	genResp := NewGenericResponse(0)
	return genResp, nil
}

func UpdateOrderDispacherStatus(order models.OrderDispacherStatus) (*GenericResponse, error) {
	if order.OrderID == 0 {
		return NewGenericResponse(-1001), nil
	}
	t := time.Now().String()
	repo := repository.GetOrderRepository()
	err := repo.UpdateOrderDispacherStatus(order)
	if err != nil {
		return NewGenericResponse(-1000), nil
	}
	// get order detail by order ID
	or, _ := repo.GetOrderByID(order.OrderID)
	// handle order status for dispacher
	// if status=flags.booked, insert row into rs_trip table.
	if order.Status == flags.Booked {
		// todo Notification to passenger
		orderTrip := models.OrderTrip{OrderID: or.OrderID, CustomerID: or.CustomerID, DriverID: order.DriverID, StartTime: t, PickupLatitude: or.PickupLatitude, PickupLongitude: or.PickupLongitude, Status: order.Status, CreateAt: t}
		err := repo.WriteOrderTrip(orderTrip)
		if err != nil {
			return NewGenericResponse(-1000), nil
		}
	}

	if order.Status == flags.Arrived {
		// todo notification to passenger
	}

	if order.Status == flags.StartTrip {
		// todo notification to passenger
		//VacantDistance
		vacantDistance := util.EarthDistance(order.DriverLongitude, order.DriverLatitude, or.PickupLongitude, or.PickupLatitude)
		orderTrip := models.OrderTrip{StartTime: t, VacantDistance: vacantDistance, PickupLatitude: order.DriverLatitude, PickupLongitude: order.DriverLongitude, Status: order.Status, UpdateAt: t, OrderID: order.OrderID}
		err := repo.ModifyOrderTrip(orderTrip)
		if err != nil {
			return NewGenericResponse(-1000), nil
		}
	}

	if order.Status == flags.EndTrip {
		// todo notification to passenger
		orderTrip, _ := repo.GetOrderTripByOrderID(order.OrderID)
		fmt.Println(util.EarthDistance(order.DriverLongitude, order.DriverLatitude, orderTrip.PickupLongitude, orderTrip.PickupLatitude))
		fmt.Println("-----", order.DriverLongitude, order.DriverLatitude, orderTrip.PickupLongitude, orderTrip.PickupLatitude)
		engagedDistance := util.EarthDistance(order.DriverLongitude, order.DriverLatitude, orderTrip.PickupLongitude, orderTrip.PickupLatitude)
		tripInfo := models.OrderTrip{EndTime: t, EngagedDistance: engagedDistance, DropoffLatitude: order.DriverLatitude, DropoffLongitude: order.DriverLongitude, Status: order.Status, UpdateAt: t, OrderID: order.OrderID}
		err := repo.ModifyOrderTrip(tripInfo)
		if err != nil {
			return NewGenericResponse(-1000), nil
		}

	}

	genResp := NewGenericResponse(0)
	return genResp, nil
}

func GetTripByOrderID(orderID int64) (*GenericResponse, error) {
	if orderID == 0 {
		return NewGenericResponse(-1001), nil
	}
	repo := repository.GetOrderRepository()
	data, err := repo.GetOrderTripByOrderID(orderID)
	if err != nil {
		return NewGenericResponse(-1000), nil
	}
	if data.OrderID == 0 {
		return NewGenericResponse(-1002), nil
	}

	//// get passenger info
	//passenger := getPassengerInfoByID(data.CustomerID)
	//data.CustomerName = passenger.CustomerName
	//data.CustomerMobile = passenger.CustomerMobile
	//
	//if data.DriverID != 0 {
	//	driver := getDriveInfoByID(data.DriverID)
	//	data.DriverName = driver.DriverName
	//	data.DriverMobile = driver.DriverMobile
	//	data.VehicleNo = driver.VehicleNo
	//	data.VehicleID = driver.VehicleID
	//}

	genResp := NewGenericResponse(0)
	(*genResp).Data = data
	return genResp, nil
}

func GetOrderByID(orderID int64) (*GenericResponse, error) {
	if orderID == 0 {
		return NewGenericResponse(-1001), nil
	}
	repo := repository.GetOrderRepository()
	data, err := repo.GetOrderByID(orderID)
	if err != nil {
		return NewGenericResponse(-1000), nil
	}
	if data.OrderID == 0 {
		return NewGenericResponse(-1002), nil
	}

	// get passenger info
	passenger := getPassengerInfoByID(data.CustomerID)
	data.CustomerName = passenger.CustomerName
	data.CustomerMobile = passenger.CustomerMobile

	if data.DriverID != 0 {
		driver := getDriveInfoByID(data.DriverID)
		data.DriverName = driver.DriverName
		data.DriverMobile = driver.DriverMobile
		data.VehicleNo = driver.VehicleNo
		data.VehicleID = driver.VehicleID
	}

	genResp := NewGenericResponse(0)
	(*genResp).Data = data
	return genResp, nil
}

func getDriveInfoByID(ID int64) models.Driver {
	mode := config.MustGetString("server.mode")
	baseURL := config.MustGetString(mode + ".driver_base_url")
	requestURL := baseURL + "/driver"
	values := make(url.Values)
	values.Set("id", strconv.FormatInt(ID, 10))
	var data models.GeneratedDriverResponse
	err := util.DoGet("Get Driver info By ID", requestURL, &values, 200, &data)
	if err != nil {
		log.Error(err)
	}
	return data.Results
}

func getPassengerInfoByID(ID int64) models.Passenger {
	mode := config.MustGetString("server.mode")
	baseURL := config.MustGetString(mode + ".passenger_base_url")
	requestURL := baseURL + "/passenger"
	values := make(url.Values)
	values.Set("id", strconv.FormatInt(ID, 10))
	var data models.GeneratedPassengerResponse
	err := util.DoGet("Get passenger info By ID", requestURL, &values, 200, &data)
	if err != nil {
		log.Error(err)
	}
	return data.Results
}
