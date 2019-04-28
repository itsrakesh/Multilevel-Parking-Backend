package parkingmodel

import "time"

type ParkingSlot struct {
	IsAvailable      string    `json:"isAvailable" gorm:"column:isAvailable"`
	VechileType      string    `gorm:"column:vechileType"`
	Level            int       `gorm:"column:level"`
	IdParkingPricing int       `gorm:"column:idParkingPricing"`
	CreatedAt        time.Time `gorm:"column:createdAt"`
}

type ParkingLot struct {
	IdParkingSlot      int       `json:"IdParkingSlot" gorm:"column:idParkingSlot"`
	RegistrationNumber string    `json:"RegistrationNumber" gorm:"column:registrationNumber"`
	Color              string    `json:"Color" gorm:"column:color"`
	StartTime          time.Time `json:"StartTime" gorm:"column:startTime"`
	CreatedAt          time.Time `json:"CreatedAt" gorm:"column:createdAt"`
}

type ParkingHistory struct {
	IdParkingSlot      int       `gorm:"column:idParkingSlot"`
	RegistrationNumber string    `gorm:"column:registrationNumber"`
	Color              string    `gorm:"column:color"`
	StartTime          time.Time `gorm:"column:startTime"`
	EndTime            time.Time `gorm:"column:endTime"`
	VechileType        string    `gorm:"column:vechileType"`
	Level              int       `gorm:"column:level"`
	UpdatedAt          time.Time `gorm:"column:updatedAt"`
}

type ParkingPricing struct {
	isAvailable      int       `gorm:"column:isAvailable"`
	vechileType      string    `gorm:"column:vechileType"`
	level            int       `gorm:"column:level"`
	idParkingPricing int       `gorm:"column:idParkingPricing"`
	createdAt        time.Time `gorm:"column:createdAt"`
}

type AvailableParkingSlot struct {
	Id int `gorm:"column:id"`
}

type Response struct {
	Data interface{}
	Err  error
	Msg  string
}
