package repository

import (
	"database/sql"
	"github.com/freeddser/ride-sharing/flags"
	"github.com/freeddser/ride-sharing/models"
	"time"
)

const (
	addDispacher = `insert into rs_dispacher_logs(order_id,driver_id,pick_time,pickup_longitude,pickup_latitude,
				dropoff_longitude,dropoff_latitude,distance,status,created_at,updated_at) values (?,?,?,?,?,?,?,?,?,?,?)`

	getDispacherByDriverID = `SELECT dispacher_id,order_id,driver_id,pick_time,pickup_longitude,pickup_latitude,dropoff_longitude,dropoff_latitude,distance,status,created_at,updated_at
						FROM rs_dispacher_logs where status=? and driver_id=? group by order_id `

	updateDispacherOrderStatus = `update rs_dispacher_logs set status=?,updated_at=? where order_id=?`
)

type DispacherRepository interface {
	WriteDispacher(models.Dispacher) error
	GetDispacherOrderByDriverID(int64) ([]models.Dispacher, error)
	UpdateDispacherOrderStatus(orderID int64, status int) error
}

type dispacherRepository struct {
	db    *DataSource
	redis redisconection
}

var dispacherRepo *dispacherRepository

func GetDispacherRepository() DispacherRepository {
	if dispacherRepo == nil {
		dispacherRepo = &dispacherRepository{db: mysqldatasource, redis: Redis}
	}
	return dispacherRepo

}

func (repo *dispacherRepository) UpdateDispacherOrderStatus(orderID int64, status int) error {
	t := time.Now().String()
	_, err := repo.db.Exec(updateDispacherOrderStatus, status, t, orderID)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (repo *dispacherRepository) WriteDispacher(dispacher models.Dispacher) error {
	t := time.Now().String()
	_, err := repo.db.Exec(addDispacher, dispacher.OrderID, dispacher.DriverID, dispacher.PickupTime, dispacher.PickupLongitude, dispacher.PickupLatitude, dispacher.DropoffLongitude,
		dispacher.DropoffLatitude, dispacher.Distance, dispacher.Status, t, t)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Don't add cache for this API, you can try use other system to store it, like bigquery
func (repo *dispacherRepository) GetDispacherOrderByDriverID(driverID int64) ([]models.Dispacher, error) {
	dispacherOrders := []models.Dispacher{}
	var rows *sql.Rows
	var err error
	rows, err = repo.db.Query(getDispacherByDriverID, flags.DispacherPrebook, driverID)

	defer rows.Close()
	for rows.Next() {
		var dispacher models.Dispacher
		err = rows.Scan(&dispacher.DispacherID, &dispacher.OrderID, &dispacher.DriverID, &dispacher.PickupTime, &dispacher.PickupLongitude,
			&dispacher.PickupLatitude, &dispacher.DropoffLongitude,
			&dispacher.DropoffLatitude, &dispacher.Distance,
			&dispacher.Status, &dispacher.CreateAt, &dispacher.UpdateAt)
		if err != nil {
			log.Error(err)
			return dispacherOrders, err
		}
		dispacherOrders = append(dispacherOrders, dispacher)
	}

	return dispacherOrders, nil
}
