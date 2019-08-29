package models

type GeneraResponse struct {
	Response string `json:"Response"`
}

type Order struct {
	OrderID          int64   `json:"order_id"`
	CustomerID       int64   `json:"customer_id"`
	CustomerName     string  `json:"customer_name"`
	CustomerMobile   string  `json:"customer_mobile"`
	DriverID         int64   `json:"driver_id"`
	DriverName       string  `json:"driver_name"`
	DriverMobile     string  `json:"driver_mobile"`
	VehicleID        string  `json:"vehicle_id"`
	VehicleNo        string  `json:"vehicle_no"`
	PickupTime       string  `json:"pick_time"`
	PickupLatitude   float64 `json:"pickup_latitude"`
	PickupLongitude  float64 `json:"pickup_longitude"`
	DropoffLatitude  float64 `json:"dropoff_latitude"`
	DropoffLongitude float64 `json:"dropoff_longitude"`
	Status           int     `json:"status"`
	CreateAt         string  `json:"create_at"`
	UpdateAt         string  `json:"update_at"`
}

type OrderTrip struct {
	TripID           int64   `json:"trip_id"`
	OrderID          int64   `json:"order_id"`
	CustomerID       int64   `json:"customer_id"`
	DriverID         int64   `json:"driver_id"`
	StartTime        string  `json:"start_time"`
	EndTime          string  `json:"end_time"`
	VacantDistance   float64 `json:"vacant_distance"`
	EngagedDistance  float64 `json:"engaged_distance"`
	PickupLatitude   float64 `json:"pickup_latitude"`
	PickupLongitude  float64 `json:"pickup_longitude"`
	DropoffLatitude  float64 `json:"dropoff_latitude"`
	DropoffLongitude float64 `json:"dropoff_longitude"`
	Status           int     `json:"status"`
	CreateAt         string  `json:"create_at"`
	UpdateAt         string  `json:"update_at"`
}

type OrderDispacherStatus struct {
	OrderID         int64   `json:"order_id"`
	DriverID        int64   `json:"driver_id"`
	Status          int     `json:"status"`
	DriverLatitude  float64 `json:"driver_latitude"`
	DriverLongitude float64 `json:"driver_longitude"`
}
