package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"parkingLot/adaptor/mysql"
	"parkingLot/parkingmodel"
	"parkingLot/util"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//testing
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

//CreateParking spaces in the Database
func CreateParking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var res parkingmodel.Response
	if dbConn, err := mysql.GetMySqlConn(); err == nil {
		slots, _ := strconv.Atoi(vars["number"])
		res = util.CreateParking(slots, dbConn)
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

//Park Car in the parking lot
func Park(w http.ResponseWriter, r *http.Request) {
	var (
		newParking parkingmodel.ParkingLot
		res        parkingmodel.Response
	)
	if dbConn, err := mysql.GetMySqlConn(); err == nil {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			if err = json.Unmarshal(body, &newParking); err == nil {
				res = util.ParkCar(newParking, dbConn)
			} else {
				res.Err = err
			}
		} else {
			res.Err = err
		}
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
	return
}

func Vacate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parkingSlot, _ := strconv.Atoi(vars["parkingNumber"])
	var (
		dbConn *gorm.DB
		err    error
		res    parkingmodel.Response
	)
	if dbConn, err = mysql.GetMySqlConn(); err != nil {
		res.Err = err
		return
	}
	res = util.VacateParking(parkingSlot, dbConn)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func GetCarWithColor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	color := vars["color"]
	var (
		dbConn *gorm.DB
		err    error
		res    parkingmodel.Response
	)
	if dbConn, err = mysql.GetMySqlConn(); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	res = util.GetCarWithColor(color, dbConn)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func GetCarWithRegistrationNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	registrationNumber := vars["registrationNumber"]
	var (
		dbConn *gorm.DB
		err    error
		res    parkingmodel.Response
	)
	if dbConn, err = mysql.GetMySqlConn(); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	res = util.GetCarWithRegistrationNumber(registrationNumber, dbConn)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
