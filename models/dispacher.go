package models

//order_id,driver_id,pick_time,pickup_longitude,pickup_latitude,dropoff_longitude,dropoff_latitude,distance,created_at,updated_at
type Dispacher struct {
	DispacherID      int64   `json:"dispacher_id"`
	OrderID          int64   `json:"order_id"`
	DriverID         int64   `json:"driver_id"`
	Distance         float64 `json:"distance"`
	PickupTime       string  `json:"pick_time"`
	PickupLatitude   float64 `json:"pickup_latitude"`
	PickupLongitude  float64 `json:"pickup_longitude"`
	DropoffLatitude  float64 `json:"dropoff_latitude"`
	DropoffLongitude float64 `json:"dropoff_longitude"`
	Status           int     `json:"status"`
	CreateAt         string  `json:"create_at"`
	UpdateAt         string  `json:"update_at"`
}

type DispacherDriver struct {
	DriverID    int64   `json:"driver_id"`
	DriverName  string  `json:"driver_name"`
	DriverPhone string  `json:"driver_phone"`
	VehicleID   string  `json:"vehicle_id"`
	VehicleNo   string  `json:"vehicle_no"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Status      int     `json:"status"`
	Distance    float64 `json:"distance"`
	//CreateAt    string  `json:"create_at"`
	//UpdateAt    string  `json:"update_at"`
}

type DispacherOrderStatus struct {
	OrderID         int64   `json:"order_id"`
	DriverID        int64   `json:"driver_id"`
	Status          int     `json:"status"`
	DriverLatitude  float64 `json:"driver_latitude"`
	DriverLongitude float64 `json:"driver_longitude"`
}
