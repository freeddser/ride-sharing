package services

import (
	"encoding/json"
	"github.com/freeddser/ride-sharing/flags"
	"github.com/freeddser/ride-sharing/models"
	"github.com/freeddser/rs-common/util"
	"strconv"
)

func TrackerDriver(driver models.Driver) (*GenericResponse, error) {
	// check Driver status, if not IDIE, not tracker in driverlist
	// because this Driverlist will use for dispacher new orders
	if driver.Status != flags.DriverPrebook {
		return NewGenericResponse(-1003), nil
	}

	var CachedriverList []models.Driver
	var driverList []models.Driver

	//get driver list from redis
	driverData, err := util.Redis.GetBytes(flags.CacheDriverList)
	if err != nil {
		return NewGenericResponse(-1000), nil
	}
	if len(driverData) < 0 {
		driverList = append(driverList, driver)
	} else {
		// Duplicate removal
		json.Unmarshal(driverData, &CachedriverList)
		driverList = append(driverList, driver)
		for _, dr := range CachedriverList {
			if dr.DriverID != driver.DriverID {
				driverList = append(driverList, dr)
			}
		}
	}
	// save new driver list into redis
	driverInfos, err := json.Marshal(&driverList)
	if err != nil {
		return NewGenericResponse(-1000), nil
	}

	util.Redis.SetBytes(flags.CacheDriverList, driverInfos, 86400)
	genResp := NewGenericResponse(0)
	(*genResp).Data = driverList
	return genResp, nil
}

func TrackerOrder(driver models.Driver) (*GenericResponse, error) {

	var CachedriverList []models.Driver
	var driverList []models.Driver

	// the driver offered a order, tracker order
	if driver.OrderID != 0 {
		driverData, err := util.Redis.GetBytes(flags.CacheTrackerOrderPrefix + strconv.FormatInt(driver.OrderID, 10))
		if err != nil {
			return NewGenericResponse(-1000), nil
		}
		if len(driverData) < 1 {
			driverList = append(driverList, driver)
		} else {
			json.Unmarshal(driverData, &CachedriverList)
			driverList = append(driverList, driver)
			for _, dr := range CachedriverList {
				driverList = append(driverList, dr)
			}
		}
		// save new driver list into redis
		driverInfos, err := json.Marshal(&driverList)
		if err != nil {
			return NewGenericResponse(-1000), nil
		}

		util.Redis.SetBytes(flags.CacheTrackerOrderPrefix+strconv.FormatInt(driver.OrderID, 10), driverInfos, 86400)
	}
	//todo handle other status
	genResp := NewGenericResponse(0)
	(*genResp).Data = driverList
	return genResp, nil
}

// get driver last location with orderID
func GetTrackerByOrderID(orderID int64) (*GenericResponse, error) {
	var CachedriverList []models.Driver
	var driverList []models.Driver

	driverData, err := util.Redis.GetBytes(flags.CacheTrackerOrderPrefix + strconv.FormatInt(orderID, 10))
	if err != nil {
		return NewGenericResponse(-1000), nil
	}
	if len(driverData) < 1 {
		return NewGenericResponse(-1002), nil
	} else {
		json.Unmarshal(driverData, &CachedriverList)
		for _, dr := range CachedriverList {
			driverList = append(driverList, dr)
		}
	}

	genResp := NewGenericResponse(0)
	if len(driverList) > 0 {
		(*genResp).Data = driverList[len(driverList)-1]
	}
	return genResp, nil
}
