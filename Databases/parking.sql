-- Create syntax for TABLE 'parking_history'
CREATE TABLE `parking_history` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `idParkingSlot` int(32) NOT NULL,
  `registrationNumber` varchar(50) NOT NULL,
  `color` varchar(50) NOT NULL,
  `startTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `endTime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `VechileType` enum('MoterCycle','car','Truck','bus','Auto') DEFAULT NULL,
  `level` int(2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;

-- Create syntax for TABLE 'parking_lot'
CREATE TABLE `parking_lot` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `idParkingSlot` int(32) NOT NULL,
  `registrationNumber` varchar(50) NOT NULL,
  `color` varchar(50) DEFAULT NULL,
  `startTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=latin1;

-- Create syntax for TABLE 'parking_pricing'
CREATE TABLE `parking_pricing` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `id_parking_slot` int(32) NOT NULL,
  `base_price` decimal(10,2) NOT NULL,
  `base_price_expiry_hours` int(2) DEFAULT NULL,
  `hourly_charges` decimal(10,2) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Create syntax for TABLE 'parking_slot'
CREATE TABLE `parking_slot` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `isAvailable` enum('yes','no') DEFAULT NULL,
  `level` int(2) NOT NULL,
  `vechileType` enum('MoterCycle','car','Truck','bus','Auto') DEFAULT NULL,
  `idParkingPricing` int(32) DEFAULT NULL,
  `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=latin1;
