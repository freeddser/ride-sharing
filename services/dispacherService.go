package services

import (
	"encoding/json"
	"github.com/freeddser/rs-common/config"
	"github.com/freeddser/ride-sharing/flags"
	"github.com/freeddser/ride-sharing/models"
	"github.com/freeddser/ride-sharing/repository"
	"github.com/freeddser/rs-common/util"
)

func GetDispacherOrderListByDriverID(driverID int64) (*GenericResponse, error) {
	if driverID == 0 {
		return NewGenericResponse(-1001), nil
	}
	repo := repository.GetDispacherRepository()
	data, err := repo.GetDispacherOrderByDriverID(driverID)
	if err != nil {
		return NewGenericResponse(-1000), nil
	}
	genResp := NewGenericResponse(0)
	(*genResp).Data = data
	return genResp, nil
}

func ModifyDispacherOrderStatus(order models.DispacherOrderStatus) (*GenericResponse, error) {
	repo := repository.GetDispacherRepository()
	err := repo.UpdateDispacherOrderStatus(order.OrderID, order.Status)
	if err != nil {
		return NewGenericResponse(-1000), nil
	}

	orderStatusRequest := models.OrderDispacherStatus{OrderID: order.OrderID, DriverID: order.DriverID, Status: order.Status, DriverLatitude: order.DriverLatitude, DriverLongitude: order.DriverLongitude}
	payload, err := json.Marshal(orderStatusRequest)
	if err != nil {
		log.Error(err)
	}
	// get Order service url
	mode := config.MustGetString("server.mode")
	orderServiceEndpoint := config.MustGetString(mode + ".order_base_url")
	response := new(models.GeneratedResponse)
	util.DoPost("Modify dispacher order status", orderServiceEndpoint+"/orders/status", string(payload), 200, &response)
	// order service will handle  status
	genResp := NewGenericResponse(0)
	return genResp, nil
}

func DispacherOrder(orderID int64) (*GenericResponse, error) {
	if orderID == 0 {
		return NewGenericResponse(-1001), nil
	}
	//get order info by orderID
	repo := repository.GetOrderRepository()
	order, err := repo.GetOrderByID(orderID)
	if err != nil || order.OrderID == 0 {
		return NewGenericResponse(-1002), nil
	}

	//if status!=Prebook should return Order: Scheduled
	if order.Status != flags.Prebook {
		return NewGenericResponse(-1005), nil
	}

	// Get Driver List from Redis and check Which Driver can get new order info
	//var driverIDs []string

	//get driver list from redis
	var CachedriverList []models.DispacherDriver
	var driverPrebookList []models.DispacherDriver

	//get driver list from redis
	driverData, err := util.Redis.GetBytes(flags.CacheDriverList)
	if err != nil {
		return NewGenericResponse(-1000), nil
	}
	if len(driverData) < 1 {
		return NewGenericResponse(-1004), nil
	} else {
		// get driverPrebookList from redis, then check location
		json.Unmarshal(driverData, &CachedriverList)
		//fmt.Println(driverData)
		for _, dr := range CachedriverList {
			//fmt.Println(dr)
			//if vehicle around you 5km, will recive the order
			distance := util.EarthDistance(order.PickupLatitude, order.PickupLongitude, dr.Latitude, dr.Longitude)
			//fmt.Println("-------", order.PickupLatitude, order.PickupLongitude, dr.Latitude, dr.Longitude)
			//fmt.Println("###############", distance)
			if distance < flags.DriverPassengerDistance {
				dr.Distance = distance
				driverPrebookList = append(driverPrebookList, dr)
				// write into dispacher log table
				// dispacher order struct.
				repo := repository.GetDispacherRepository()
				dispacher := models.Dispacher{
					OrderID:          orderID,
					DriverID:         dr.DriverID,
					PickupTime:       order.PickupTime,
					PickupLongitude:  order.PickupLongitude,
					PickupLatitude:   order.PickupLatitude,
					DropoffLongitude: order.DropoffLongitude,
					DropoffLatitude:  order.DropoffLatitude,
					Distance:         distance,
					Status:           flags.DispacherPrebook,
				}
				err := repo.WriteDispacher(dispacher)
				if err != nil {
					return NewGenericResponse(-1000), nil
				}

			}
		}

		if len(driverPrebookList) < 1 {
			return NewGenericResponse(-1004), nil
		}

		genResp := NewGenericResponse(0)
		(*genResp).Data = driverPrebookList
		return genResp, nil
	}

}
