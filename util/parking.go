package util

import (
	"parkingLot/constant"
	"parkingLot/parkingmodel"

	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

func buildResponse(data interface{}, err error, msg string) (res parkingmodel.Response) {
	res.Data = data
	res.Err = err
	res.Msg = msg
	return
}

func CreateParking(numberOfCarParking int, dbConn *gorm.DB) (res parkingmodel.Response) {
	oldParkingCount := 0
	numberOfParkingCreated := 0
	getNumberOfOldParkingSQL := `SELECT count(1) as oldParkingCount from parking_slot`
	dbConn.Debug().Table("parking_slot").Select(getNumberOfOldParkingSQL).Count(&oldParkingCount)
	if numberOfCarParking != 0 && numberOfCarParking+oldParkingCount < constant.MAX_PARKING {
		var newparking parkingmodel.ParkingSlot
		for i := 0; i < numberOfCarParking; i++ {
			newparking = getNewCarParkingSlotConfig(numberOfCarParking)
			if err := dbConn.Debug().Table("parking_slot").Create(newparking).Error; err != nil {
				res.Err = err
			} else {
				numberOfParkingCreated++
			}
		}
	} else {
		res.Msg = "Please verify the Input Param. Max Allowed Parking is " + strconv.Itoa(constant.MAX_PARKING) + ". Already Created Parking is " + strconv.Itoa(oldParkingCount)
		return
	}
	res.Msg = "Created a parking lot with " + strconv.Itoa(numberOfParkingCreated) + " slots"
	return
}

func ParkCar(newParking parkingmodel.ParkingLot, dbConn *gorm.DB) (res parkingmodel.Response) {
	var (
		firstAvailableSlot parkingmodel.AvailableParkingSlot
	)
	if err := dbConn.Debug().Table("parking_slot").Where("isAvailable = ?", "yes").Find(&firstAvailableSlot).Error; err != nil || firstAvailableSlot.Id < 1 {
		res.Msg = "Parking is Full"
		return
	}

	newParking.IdParkingSlot = firstAvailableSlot.Id
	newParking.StartTime, newParking.CreatedAt = time.Now(), time.Now()

	if err := dbConn.Debug().Table("parking_lot").Create(newParking).Error; err != nil {
		res.Err = err
	} else {
		updateSlotStatus(firstAvailableSlot.Id, constant.PARKING_NOT_AVAILABLE_STATUS, dbConn)
		res.Data = newParking
	}
	return
}

func VacateParking(parkingSlot int, dbConn *gorm.DB) (res parkingmodel.Response) {
	var parkingHistory parkingmodel.ParkingHistory
	if err := dbConn.Debug().Table("parking_lot").Where(" idParkingSlot= ?", parkingSlot).Find(&parkingHistory).Error; err != nil {
		res.Err = err
		return
	}
	res.Err = logParkingHistory(parkingHistory, dbConn)
	if res.Err != nil {
		res.Msg += res.Err.Error()
	}
	res.Err = deleteCurrentParking(parkingSlot, dbConn)
	if res.Err != nil {
		res.Msg += res.Err.Error()
	}
	res.Err = updateSlotStatus(parkingSlot, constant.PARKING_AVAILABLE_STATUS, dbConn)
	if res.Err != nil {
		res.Msg += res.Err.Error()
	}
	res.Msg += "Parking Vacated"
	return
}

func GetCarWithColor(color string, dbConn *gorm.DB) (res parkingmodel.Response) {
	var parkingLot []parkingmodel.ParkingLot
	if err := dbConn.Debug().Table("parking_lot").Where(" color= ?", color).Find(&parkingLot).Error; err != nil {
		res.Err = err
		return
	} else {
		res.Data = parkingLot
	}
	return
}

func GetCarWithRegistrationNumber(registrationNumber string, dbConn *gorm.DB) (res parkingmodel.Response) {
	var parkingLot []parkingmodel.ParkingLot
	if err := dbConn.Debug().Table("parking_lot").Where(" registrationNumber= ?", registrationNumber).Find(&parkingLot).Error; err != nil {
		res.Err = err
		return
	} else {
		res.Data = parkingLot
	}
	return
}

func getNewCarParkingSlotConfig(numberOfCarParking int) (newparking parkingmodel.ParkingSlot) {
	newparking.IsAvailable = constant.PARKING_AVAILABLE_STATUS
	newparking.VechileType = "car"
	newparking.Level = 0
	newparking.IdParkingPricing = 1
	newparking.CreatedAt = time.Now()
	return
}

func updateSlotStatus(latestAvailableSlotId int, status string, dbConn *gorm.DB) error {
	updateColVal := map[string]interface{}{"isAvailable": status}
	return dbConn.Debug().Table("parking_slot").
		Where(`id = ?`, latestAvailableSlotId).
		Updates(updateColVal).Error
}

func logParkingHistory(parkingHistory parkingmodel.ParkingHistory, dbConn *gorm.DB) error {
	parkingHistory.EndTime = time.Now()
	return dbConn.Debug().Table("parking_history").Create(parkingHistory).Error
}

func deleteCurrentParking(parkingSlot int, dbConn *gorm.DB) error {
	var parkingLot parkingmodel.ParkingLot
	parkingLot.IdParkingSlot = parkingSlot
	return dbConn.Debug().Table("parking_lot").Where("idParkingSlot=?", parkingSlot).Delete(parkingLot).Error
}
