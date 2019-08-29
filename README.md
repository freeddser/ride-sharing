# 这是一道面试题,在两天内,做一个最小化的打车服务

 1.用户可以下单

 2.司机可以接单

 3.系统可以调度司机,并且只有一定距离内的司机可以接单

 4.当司机接单,到达,行程开始,结束时用户能够知道.

 5.当司机快到达时,给用户提醒

 6.行程结束后,可以查看行驶距离,时间等信息


# 运行服务的预先要求
*  1. Redis          (v. 4.0.9+)
*  2. Mysql       (v. 5.6+)

go run main.go -c config.toml


#Test Data
    passenger ID:1001:
    pickup_location:120.428701,36.06874
    dropoff_lication:120.362873,36.106803

    driver1 Bob 2001 location: 120.430102,36.069236
    driver2 Ken 2002 location: 116.39737,40.047623

#测试流程

1.Passenger创建订单

POST：OrderService/orders/create

    {
    "customer_id":1001,
	"pick_time":"2019-08-22 15:04:05",
	"pickup_latitude":120.428701,
	"pickup_longitude":36.06874,
	"dropoff_latitude":120.362873,
	"dropoff_longitude":36.106803

    }
2.Passenger可以查看当前订单状态通过接口

GET:OrderService/orders/?id=order_id

    // Passenger send the request, waiting driver accept
    Prebook = 0

    // Driver accept the order
    Booked = 1

    // Driver Arrived to Passenger
    Arrived = 2

    // Driver take the passenger and start the trip
    StartTrip = 3

    // Driver take the passenger and end the trip
    EndTrip = 4

    {
    "success": true,
    "error": {
        "error_id": 0,
        "message": {
            "cn": "成功",
            "en": "Success"
        }
    },
    "results": {
        "order_id": 11,
        "customer_id": 1001,
        "customer_name": "Alice",
        "customer_mobile": "65321111",
        "driver_id": 0,
        "driver_name": "",
        "driver_mobile": "",
        "vehicle_id": "",
        "vehicle_no": "",
        "pick_time": "2019-08-20 15:04:05",
        "pickup_latitude": 120.429545,
        "pickup_longitude": 36.069554,
        "dropoff_latitude": 120.431989,
        "dropoff_longitude": 36.080039,
        "status": 0,
        "create_at": "2019-08-22 09:45:42",
        "update_at": "2019-08-22 09:45:42"
    }
    }

3.Driver上传自己的信息到RS系统.

POST:trackerService/tracker/driver
    {
	"driver_id":2001,
	"latitude":120.430102,
	"longitude":36.069236,
	"status":1,

    }

4.系统可以通过Dispacher接口调度订单，距离Passenger 5KM的Driver可以获得订单消息

GET：dispacherService/dispacher?id=12

    {
    "success": true,
    "error": {
        "error_id": 0,
        "message": {
            "cn": "成功",
            "en": "Success"
        }
    },
    "results": [
        {
            "driver_id": 2001,
            "driver_name": "",
            "driver_phone": "",
            "vehicle_id": "",
            "vehicle_no": "",
            "latitude": 120.430102,
            "longitude": 36.069236,
            "status": 0,
            "distance": 158.2686166924352
        }
    ]
    }

5.Driver可以通过Dispacher接口拉取,自己能够接的订单

GET:dispacherService/dispacher/getorders?id=driver_id

    {
    "success": true,
    "error": {
        "error_id": 0,
        "message": {
            "cn": "成功",
            "en": "Success"
        }
    },
    "results": [
        {
            "dispacher_id": 14,
            "order_id": 12,
            "driver_id": 2001,
            "distance": 158.2686166924352,
            "pick_time": "2019-08-22 15:04:05",
            "pickup_latitude": 120.428701,
            "pickup_longitude": 36.06874,
            "dropoff_latitude": 120.362873,
            "dropoff_longitude": 36.106803,
            "status": 0,
            "create_at": "2019-08-22 12:04:38",
            "update_at": "2019-08-22 12:04:38"
        }
    ]
    }

6.Driver修改订单状态,如接单,接到passenger,开始行程,结束行程等,修改dispacher order/order rs_orser,rs_trip

POST:dispacherService/dispacher/updateorder

    {
	"order_id":12,
	"driver_id":2001,
	"driver_latitude":120.430808,
    "driver_longitude":36.069094,
	"status":1  //booked
    }

7.Driver可以时时上传自己的位置,和订单

POST: TrackerService/tracker/order/driver

    {
	  "driver_id":106,
		"latitude":120.430808,
		"longitude":36.069094,
		"order_id":6,
		"status":1  //order status booked

    }

8.可以查看Diver绑定订单的最后位置

GET:TrackerService/tracker/order?id=6

    {
    "success": true,
    "error": {
        "error_id": 0,
        "message": {
            "cn": "成功",
            "en": "Success"
        }
    },
    "results": {
        "driver_id": 106,
        "order_id": 6,
        "driver_name": "",
        "driver_mobile": "",
        "vehicle_id": "",
        "vehicle_no": "",
        "latitude": 120.430808,
        "longitude": 36.069094,
        "status": 0,
        "create_at": "",
        "update_at": ""
    }
    }

9.Driver到达后,点到达,将空驶距离,记录进rs_trip

POST:dispacherService/dispacher/updateorder

    {
	"order_id":6,
	"driver_id":2001,
	"driver_latitude":120.430808,
    "driver_longitude":36.069094,
	"status":2  //Arrived
    }

10.Driver到达后,StartTrip

POST:dispacherService/dispacher/updateorder

    {
	"order_id":6,
	"driver_id":2001,
	"driver_latitude":120.430808,
    "driver_longitude":36.069094,
	"status":3  //StartTrip
    }

11.Driver到达后,点到达,将行驶距离,记录进rs_trip

POST:dispacherService/dispacher/updateorder

    {
	"order_id":6,
	"driver_id":2001,
	"driver_latitude":120.430808,
    "driver_longitude":36.069094,
	"status":4  //StopTrip
    }

12. 可以查看整个行程信息

GET: OrderService:/orders/gettripbyorderid?id=6

    {
    "success": true,
    "error": {
        "error_id": 0,
        "message": {
            "cn": "成功",
            "en": "Success"
        }
    },
    "results": {
        "trip_id": 1,
        "order_id": 6,
        "customer_id": 1001,
        "driver_id": 106,
        "start_time": "2019-08-22 11:48:34",
        "end_time": "2019-08-22 11:50:36",
        "vacant_distance": 12694848,
        "engaged_distance": 98,
        "pickup_latitude": 120.430808,
        "pickup_longitude": 36.069094,
        "dropoff_latitude": 120.430129,
        "dropoff_longitude": 36.06978,
        "status": 4,
        "create_at": "2019-08-22 09:06:33",
        "update_at": "2019-08-22 11:50:36"
    }
    }
