package flags

const ACTIVE = 0
const DELETE = -1

// orders status
// Passenger send the request, waiting driver accept
const Prebook = 0

// Driver accept the order
const Booked = 1

// Driver Arrived to Passenger
const Arrived = 2

// Driver take the passenger and start the trip
const StartTrip = 3

// Driver take the passenger and end the trip
const EndTrip = 4

// Approaching Distance default is 2km
const ApproachingDistance = 2000

// if Driver around passenger 5km, will recive the order
const DriverPassengerDistance = 5000

// cache
const CacheDriverList = "CacheDriverList"
const CacheTrackerOrderPrefix = "CacheTrackerOrder_"
const CacheOrderPrefix = "ORDER_"
const CacheOrderTripPrefix = "ORDER_TRIP_"
const CachePassengerPrefix = "PASSENGER_"
const CacheDriverPrefix = "DRIVER_"
const CACHETTL = 86400

// driver
const DriverPrebook = 0
const DriverBooked = 1

// dispacher
const DispacherPrebook = 0
const DispacherOffer = 1
