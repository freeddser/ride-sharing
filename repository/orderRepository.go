package repository

import (
	"encoding/json"
	"github.com/freeddser/ride-sharing/flags"
	"github.com/freeddser/ride-sharing/models"
	"strconv"
	"time"
)

const (
	addOrder = `insert into rs_order(pickup_time,customer_id ,driver_id,pickup_longitude ,pickup_latitude ,dropoff_longitude ,dropoff_latitude,status,created_at ,updated_at) values (?,?,?,?,?,?,?,?,?,?)`

	updateOrder = `update rs_order set pickup_time=?,customer_id=?,driver_id=?,pickup_longitude=? ,pickup_latitude=?,dropoff_longitude=?,dropoff_latitude=?,status=?,updated_at=? where order_id=?`

	updateDispacherStatus = `update rs_order set driver_id=?,status=?,updated_at=? where order_id=?`

	getOrderByID = `SELECT order_id,pickup_time,customer_id,driver_id,
		pickup_longitude,pickup_latitude,dropoff_longitude ,dropoff_latitude,status,created_at,updated_at FROM rs_order where order_id=?`
	getOrderByStatus = `SELECT order_id,pickup_time,customer_id,driver_id,pickup_longitude,pickup_latitude,dropoff_longitude ,dropoff_latitude,status,created_at,updated_at FROM rs_order where status=?`

	addOrderTrip = `insert into rs_trip (customer_id,order_id,driver_id,start_time,end_time,vacant_distance ,engaged_distance,
	pickup_longitude ,pickup_latitude  ,dropoff_longitude  ,dropoff_latitude,Status,created_at,updated_at) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	updateOrderTripStart = `update rs_trip set start_time=?,vacant_distance=?,pickup_longitude=? ,pickup_latitude=?,status=?,updated_at=? where order_id=?`
	updateOrderTripEnd   = `update rs_trip set end_time=?,engaged_distance=?,dropoff_longitude=?,dropoff_latitude=? ,status=?,updated_at=? where order_id=?`

	getOrderTripByOrderID = `select  trip_id,customer_id,order_id,driver_id,start_time,end_time,vacant_distance ,engaged_distance,
	pickup_longitude ,pickup_latitude  ,dropoff_longitude  ,dropoff_latitude,Status,created_at,updated_at from rs_trip where order_id=?`
)

type OrderRepository interface {
	WriteOrder(models.Order) (models.Order, error)
	UpdateOrder(models.Order) error

	UpdateOrderDispacherStatus(order models.OrderDispacherStatus) error

	GetOrderByID(int64) (models.Order, error)
	GetOrderByStatus(int) (models.Order, error)

	WriteOrderTrip(models.OrderTrip) error
	ModifyOrderTrip(order models.OrderTrip) error
	GetOrderTripByOrderID(orderID int64) (models.OrderTrip, error)
}

type orderRepository struct {
	db    *DataSource
	redis redisconection
}

var orderRepo *orderRepository

func GetOrderRepository() OrderRepository {
	if orderRepo == nil {
		orderRepo = &orderRepository{db: mysqldatasource, redis: Redis}
	}
	return orderRepo

}

func (repo *orderRepository) GetOrderTripByOrderID(orderID int64) (models.OrderTrip, error) {
	cacheData, _ := repo.redis.getBytes(flags.CacheOrderTripPrefix + strconv.FormatInt(orderID, 10))
	order := models.OrderTrip{}
	if len(cacheData) < 1 {
		//select  trip_id,customer_id,order_id,driver_id,start_time,end_time,vacant_distance ,engaged_distance,
		//	pickup_longitude ,pickup_latitude  ,dropoff_longitude  ,dropoff_latitude,Status,created_at,updated_at from rs_trip where order_id=
		err := repo.db.QueryRow(getOrderTripByOrderID, orderID).Scan(&order.TripID, &order.CustomerID, &order.OrderID, &order.DriverID,
			&order.StartTime, &order.EndTime, &order.VacantDistance, &order.EngagedDistance,
			&order.PickupLongitude, &order.PickupLatitude, &order.DropoffLongitude, &order.DropoffLatitude, &order.Status, &order.CreateAt, &order.UpdateAt)
		if err != nil {
			log.Error(err)
			return order, err
		}
		data, err := json.Marshal(order)
		if err != nil {
			return order, err
		}
		repo.redis.setBytes(flags.CacheOrderTripPrefix+strconv.FormatInt(orderID, 10), []byte(data), flags.CACHETTL)
	} else {
		json.Unmarshal(cacheData, &order)
	}

	return order, nil
}

func (repo *orderRepository) ModifyOrderTrip(orderTrip models.OrderTrip) error {
	if orderTrip.Status == flags.StartTrip {
		_, err := repo.db.Exec(updateOrderTripStart, orderTrip.StartTime, orderTrip.VacantDistance, orderTrip.PickupLongitude, orderTrip.PickupLatitude, orderTrip.Status, orderTrip.UpdateAt, orderTrip.OrderID)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	if orderTrip.Status == flags.EndTrip {
		_, err := repo.db.Exec(updateOrderTripEnd, orderTrip.EndTime, orderTrip.EngagedDistance, orderTrip.DropoffLongitude, orderTrip.DropoffLatitude, orderTrip.Status, orderTrip.UpdateAt, orderTrip.OrderID)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	//delete cache
	repo.redis.delCache(flags.CacheOrderTripPrefix + strconv.FormatInt(orderTrip.OrderID, 10))
	return nil
}

func (repo *orderRepository) WriteOrderTrip(orderTrip models.OrderTrip) error {

	//customer_id,order_id,driver_id,start_time,end_time,vacant_distance ,engaged_distance,
	//	pickup_longitude ,pickup_latitude  ,dropoff_longitude  ,dropoff_latitude,Status,created_at,updated_at

	result, err := repo.db.Exec(addOrderTrip, orderTrip.CustomerID, orderTrip.OrderID, orderTrip.DriverID, orderTrip.StartTime, orderTrip.EndTime,
		orderTrip.VacantDistance, orderTrip.EngagedDistance, orderTrip.PickupLongitude, orderTrip.PickupLatitude, orderTrip.DropoffLongitude, orderTrip.DropoffLatitude, orderTrip.Status, orderTrip.CreateAt, orderTrip.UpdateAt)
	if err != nil {
		log.Error(err)
		return err
	}

	tripID, _ := result.LastInsertId()
	orderTrip.TripID = tripID

	// write to Cache
	data, err := json.Marshal(orderTrip)
	if err != nil {
		return err
	}
	// use orderID as the key
	repo.redis.setBytes(flags.CacheOrderTripPrefix+strconv.FormatInt(orderTrip.OrderID, 10), []byte(data), flags.CACHETTL)
	return nil
}

func (repo *orderRepository) UpdateOrderDispacherStatus(order models.OrderDispacherStatus) error {
	t := time.Now().String()
	_, err := repo.db.Exec(updateDispacherStatus, order.DriverID, order.Status, t, order.OrderID)
	if err != nil {
		log.Error(err)
		return err
	}

	// write to Cache
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}
	//update cache for order by orderID
	repo.redis.delCache(flags.CacheOrderPrefix + strconv.FormatInt(order.OrderID, 10))
	orderData, _ := repo.GetOrderByID(order.OrderID)
	if orderData.OrderID != 0 {
		repo.redis.setBytes(flags.CacheOrderPrefix+strconv.FormatInt(order.OrderID, 10), []byte(data), flags.CACHETTL)
	}
	return nil
}

func (repo *orderRepository) UpdateOrder(order models.Order) error {
	t := time.Now().String()
	_, err := repo.db.Exec(updateOrder, order.PickupTime, order.CustomerID, order.DriverID, order.PickupLongitude, order.PickupLatitude, order.DropoffLongitude, order.DropoffLatitude, order.Status, t, order.OrderID)
	if err != nil {
		log.Error(err)
		return err
	}

	//delete cache
	repo.redis.delCache(flags.CacheOrderPrefix + strconv.FormatInt(order.OrderID, 10))
	return nil
}

func (repo *orderRepository) WriteOrder(order models.Order) (models.Order, error) {
	t := time.Now().String()
	result, err := repo.db.Exec(addOrder, order.PickupTime, order.CustomerID, order.DriverID, order.PickupLongitude, order.PickupLatitude, order.DropoffLongitude, order.DropoffLatitude, flags.Prebook, t, t)
	if err != nil {
		log.Error(err)
		return order, err
	}
	orderID, _ := result.LastInsertId()
	order.OrderID = orderID

	// write to Cache
	data, err := json.Marshal(order)
	if err != nil {
		return order, err
	}
	repo.redis.setBytes(flags.CacheOrderPrefix+strconv.FormatInt(orderID, 10), []byte(data), flags.CACHETTL)
	return order, nil
}

func (repo *orderRepository) GetOrderByStatus(status int) (models.Order, error) {
	order := models.Order{}
	err := repo.db.QueryRow(getOrderByStatus, status).Scan(&order.OrderID, &order.PickupTime, &order.CustomerName, &order.CustomerID, &order.DriverID,
		&order.PickupLongitude, &order.PickupLatitude, &order.DropoffLongitude, &order.DropoffLatitude, &order.Status, &order.CreateAt, &order.UpdateAt)
	if err != nil {
		log.Error(err)
		return order, err
	}
	return order, nil
}

func (repo *orderRepository) GetOrderByID(orderID int64) (models.Order, error) {
	cacheData, _ := repo.redis.getBytes(flags.CacheOrderPrefix + strconv.FormatInt(orderID, 10))
	order := models.Order{}
	if len(cacheData) < 1 {
		err := repo.db.QueryRow(getOrderByID, orderID).Scan(&order.OrderID, &order.PickupTime, &order.CustomerID, &order.DriverID, &order.PickupLongitude, &order.PickupLatitude, &order.DropoffLongitude, &order.DropoffLatitude, &order.Status, &order.CreateAt, &order.UpdateAt)
		if err != nil {
			log.Error(err)
			return order, err
		}
		data, err := json.Marshal(order)
		if err != nil {
			return order, err
		}
		repo.redis.setBytes(flags.CacheOrderPrefix+strconv.FormatInt(orderID, 10), []byte(data), flags.CACHETTL)
	} else {
		json.Unmarshal(cacheData, &order)
	}

	return order, nil
}
