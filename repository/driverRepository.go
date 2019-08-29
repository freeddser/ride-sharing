package repository

import (
	"encoding/json"
	"github.com/freeddser/ride-sharing/flags"
	"github.com/freeddser/ride-sharing/models"
	"strconv"
)

const (
	getDriverByID = `SELECT driver_id,driver_name,driver_mobile,vehicle_id,vehicle_no,status,created_at,updated_at FROM rs_driver where driver_id=? and status=?`
)

type DriverRepository interface {
	GetDriverByID(int64) (models.Driver, error)
}

type driverRepository struct {
	db    *DataSource
	redis redisconection
}

var driverRepo *driverRepository

func GetDriverRepository() DriverRepository {
	if driverRepo == nil {
		driverRepo = &driverRepository{db: mysqldatasource, redis: Redis}
	}
	return driverRepo

}

func (repo *driverRepository) GetDriverByID(driverID int64) (models.Driver, error) {
	cacheData, _ := repo.redis.getBytes(flags.CacheDriverPrefix + strconv.FormatInt(driverID, 10))
	driver := models.Driver{}
	if len(cacheData) < 1 {
		err := repo.db.QueryRow(getDriverByID, driverID, flags.ACTIVE).Scan(&driver.DriverID, &driver.DriverName, &driver.DriverMobile, &driver.VehicleID, &driver.VehicleNo, &driver.Status, &driver.CreateAt, &driver.UpdateAt)
		if err != nil {
			log.Error(err)
			return driver, err
		}
		data, err := json.Marshal(driver)
		if err != nil {
			return driver, err
		}

		repo.redis.setBytes(flags.CacheDriverPrefix+strconv.FormatInt(driverID, 10), []byte(data), flags.CACHETTL)
	} else {
		json.Unmarshal(cacheData, &driver)
	}

	return driver, nil
}
