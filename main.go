package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Stop struct {
	StopID   string  `json:"stop_id"`
	Distance float64 `json:"distance_km"`
}

type Route struct {
	RouteID       string `json:"route_id"`
	Stops         []Stop `json:"stops"`
	TotalDistance float64
}

type Vehicle struct {
	ID        string `json:"id"`
	Capacity  int    `json:"capacity_kwh"`
	KwhPer100 int    `json:"kwh_per_100_km"`
	Range     int
}

type RouteData struct {
	Routes []*Route `json:"routes"`
}

type VehicleData struct {
	Vehicles []*Vehicle `json:"vehicles"`
}

func main() {
	routeData := &RouteData{}
	if rf, err := os.Open("./routes.json"); err != nil {
		log.Fatal(err)
	} else {
		defer rf.Close()
		err = json.NewDecoder(rf).Decode(routeData)
		if err != nil {
			log.Fatal(err)
		}
	}

	vehicleData := &VehicleData{}
	if vf, err := os.Open("./vehicles.json"); err != nil {
		log.Fatal(err)
	} else {
		defer vf.Close()
		err = json.NewDecoder(vf).Decode(vehicleData)
		if err != nil {
			log.Fatal(err)
		}
	}

	totalRoutesDistance := 0.0

	// Calculate total distance for each route.
	for _, route := range routeData.Routes {
		totalDistance := 0.0
		for _, stop := range route.Stops {
			totalDistance += stop.Distance
		}
		route.TotalDistance = totalDistance
		totalRoutesDistance += totalDistance
	}

	leastEfficientVehicle := vehicleData.Vehicles[0]
	for _, vehicle := range vehicleData.Vehicles {
		if vehicle.KwhPer100 > leastEfficientVehicle.KwhPer100 {
			leastEfficientVehicle = vehicle
		}
	}

	// Calculate least efficient vehicle energy spent.
	energy := (totalRoutesDistance * float64(leastEfficientVehicle.KwhPer100)) / 100
	fmt.Printf("Least efficent energy: %f\n", energy)

	// Calculate range for each vehicle
	for _, vehicle := range vehicleData.Vehicles {
		vehicle.Range = (vehicle.Capacity / vehicle.KwhPer100) * 100
	}

	// pairs := make(map[string]Vehicle)
}
