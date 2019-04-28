package util

import (
	"encoding/json"
	"parkingLot/adaptor/mysql"
	"parkingLot/parkingmodel"
	"testing"
)

func CreateParkingZeroInputTest(t *testing.T) {
	var res parkingmodel.Response
	if dbConn, err := mysql.GetMySqlConn(); err == nil { //In Production unitTest will mock the Database coneection response
		res = CreateParking(0, dbConn)
		assertEqual(t, res.Err, nil)
	}
}

func CreateParkingSuccessTest(t *testing.T) {
	var res parkingmodel.Response
	if dbConn, err := mysql.GetMySqlConn(); err == nil {
		res = CreateParking(1, dbConn)
		assertEqual(t, res.Msg, "Created a parking lot with 1 slots")
	}
}

func ParkCarTest(t *testing.T) {
	var (
		res        parkingmodel.Response
		newParking parkingmodel.ParkingLot
	)
	if dbConn, err := mysql.GetMySqlConn(); err == nil {
		res = ParkCar(getDummyParkingLotData(), dbConn)
		resByteData, _ := json.Marshal(res.Data)
		json.Unmarshal(resByteData, &newParking)
		assertEqual(t, newParking.RegistrationNumber, "DL23VK2312")
	}
}

func GetCarWithColorTest(t *testing.T) {
	var (
		res     parkingmodel.Response
		parking parkingmodel.ParkingLot
	)
	if dbConn, err := mysql.GetMySqlConn(); err == nil {
		res = GetCarWithColor("silver", dbConn)
		resByteData, _ := json.Marshal(res.Data)
		json.Unmarshal(resByteData, &parking)
		assertEqual(t, parking.Color, "Silver")
	}
}

func GetCarWithRegistrationNumberTest(t *testing.T) {
	var (
		res     parkingmodel.Response
		parking parkingmodel.ParkingLot
	)
	if dbConn, err := mysql.GetMySqlConn(); err == nil {
		res = GetCarWithColor("silver", dbConn)
		resByteData, _ := json.Marshal(res.Data)
		json.Unmarshal(resByteData, &parking)
		assertEqual(t, parking.RegistrationNumber, "DL23VK2312")
	}
}

func VacateParkingTest(t *testing.T) {
	var res parkingmodel.Response
	if dbConn, err := mysql.GetMySqlConn(); err == nil {
		res = VacateParking(1, dbConn)
		assertEqual(t, res.Msg, "Parking Vacated")
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func getDummyParkingLotData() (newParking parkingmodel.ParkingLot) {
	newParking.RegistrationNumber = "DL23VK2312"
	newParking.Color = "Silver"
	return
}
