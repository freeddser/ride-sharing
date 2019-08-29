CREATE TABLE `rs_order` (
  `order_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pickup_time` datetime NOT NULL,
  `customer_id` bigint(20) NOT NULL DEFAULT 0,
  `driver_id` bigint(20) NOT NULL DEFAULT 0,
  `pickup_longitude` double NOT NULL,
  `pickup_latitude` double NOT NULL,
  `dropoff_longitude` double NOT NULL,
  `dropoff_latitude` double NOT NULL,
  `status` char(8) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `rs_driver` (
  `driver_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `driver_name` varchar(255) NOT NULL DEFAULT ' ',
  `driver_mobile` varchar(31) NOT NULL DEFAULT ' ',
  `vehicle_id` bigint(20) NOT NULL DEFAULT '0',
  `vehicle_no` varchar(12) NOT NULL DEFAULT ' ',
  `status` char(8) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`driver_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `rs_passenger` (
  `customer_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `customer_name` varchar(255) NOT NULL DEFAULT ' ',
  `customer_mobile` varchar(31) NOT NULL DEFAULT ' ',
  `status` char(8) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



CREATE TABLE `rs_trip` (
  `trip_id` bigint(20)  NOT NULL AUTO_INCREMENT,
  `customer_id` bigint(20) NOT NULL,
  `order_id` bigint(20) NOT NULL,
  `driver_id` bigint(20) NOT NULL,
  `start_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `vacant_distance` int(11) NOT NULL,
  `engaged_distance` int(11) NOT NULL,
  `pickup_longitude` double NOT NULL,
  `pickup_latitude` double NOT NULL,
  `dropoff_longitude` double NOT NULL,
  `dropoff_latitude` double NOT NULL,
  `Status`           int    ,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`trip_id`,`order_id`),
  UNIQUE KEY `order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `rs_dispacher_logs` (
  `dispacher_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_id` bigint(20) NOT NULL,
  `driver_id` bigint(20) NOT NULL,
  `pick_time` datetime NOT NULL,
  `pickup_longitude` double NOT NULL,
  `pickup_latitude` double NOT NULL,
  `dropoff_longitude` double NOT NULL,
  `dropoff_latitude` double NOT NULL,
  `distance`  double NOT NULL,
   `status` char(8) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`dispacher_id`,`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

