package repository

import (
	"encoding/json"
	"fmt"
	"github.com/freeddser/ride-sharing/flags"
	"github.com/freeddser/ride-sharing/models"
	"strconv"
)

const (
	getPassengerByID = `SELECT customer_id,customer_name,customer_mobile,status,created_at,updated_at FROM rs_passenger where customer_id=? and status=?`
)

type PassengerRepository interface {
	GetPassengerByID(int64) (models.Passenger, error)
}

type passengerRepository struct {
	db    *DataSource
	redis redisconection
}

var passengerRepo *passengerRepository

func GetPassengerRepository() PassengerRepository {
	if passengerRepo == nil {
		passengerRepo = &passengerRepository{db: mysqldatasource, redis: Redis}
	}
	return passengerRepo

}

func (repo *passengerRepository) GetPassengerByID(passengerID int64) (models.Passenger, error) {
	cacheData, _ := repo.redis.getBytes(flags.CachePassengerPrefix + strconv.FormatInt(passengerID, 10))
	passenger := models.Passenger{}
	if len(cacheData) < 1 {
		fmt.Println("GET FROM DB")
		err := repo.db.QueryRow(getPassengerByID, passengerID, flags.ACTIVE).Scan(&passenger.CustomerID, &passenger.CustomerName, &passenger.CustomerMobile, &passenger.Status, &passenger.CreateAt, &passenger.UpdateAt)
		if err != nil {
			log.Error(err)
			return passenger, err
		}
		data, err := json.Marshal(passenger)
		if err != nil {
			return passenger, err
		}
		repo.redis.setBytes(flags.CachePassengerPrefix+strconv.FormatInt(passengerID, 10), []byte(data), flags.CACHETTL)
	} else {
		fmt.Println("GET FROM CACHE")
		json.Unmarshal(cacheData, &passenger)
	}

	return passenger, nil
}
