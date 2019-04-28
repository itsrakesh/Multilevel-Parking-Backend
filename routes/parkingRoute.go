package routes

import (
	"net/http"
	"parkingLot/services"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		services.Index,
	},
	Route{
		"CreateParking",
		"GET",
		"/create_parking_lot/{number}",
		services.CreateParking,
	},
	Route{
		"Park",
		"POST",
		"/park",
		services.Park,
	},
	Route{
		"Vacate",
		"GET",
		"/vacate/{parkingNumber}",
		services.Vacate,
	},
	Route{
		"GetCar",
		"GET",
		"/getCar/color/{color}",
		services.GetCarWithColor,
	},
	Route{
		"GetCar",
		"GET",
		"/getCar/registrationNumber/{registrationNumber}",
		services.GetCarWithRegistrationNumber,
	},
}
