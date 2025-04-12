package rss

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/jftuga/geodist"
)

type Rideservice struct {
	Passengers map[int]*Passenger
	Rides      map[int]*Ride
	Drivers    map[int]*Driver
	Ride       chan *Ride
	mu         sync.Mutex
}

var RideService *Rideservice
var Once sync.Once

func GetRideService() *Rideservice {
	Once.Do(func() {
		RideService = &Rideservice{
			Passengers: make(map[int]*Passenger),
			Rides:      make(map[int]*Ride),
			Drivers:    make(map[int]*Driver),
			Ride:       make(chan *Ride, 100),
			mu:         sync.Mutex{},
		}
		go RideService.processRideRequests()
	})
	return RideService
}

func (r *Rideservice) AddPassenger(p *Passenger) {
	r.Passengers[p.ID] = p
}

func (r *Rideservice) AddDriver(d *Driver) {
	r.Drivers[d.ID] = d
}
func (r *Rideservice) RequestRide(pId int, start, end *Location, status RideStatus, rtype RideType) *Ride {
	r.mu.Lock()
	defer r.mu.Unlock()
	p := r.Passengers[pId]

	ride := &Ride{
		ID:         int(uuid.New().ID()),
		Passenger:  p,
		From:       start,
		To:         end,
		RideStatus: status,
		RideType:   rtype,
	}
	r.Ride <- ride
	// go routine and notify  to driver

	return ride

}

func (r *Rideservice) AcceptRide(dId, rId int) bool {
	driver := r.Drivers[dId]
	ride := r.Rides[rId]
	if driver == nil || ride == nil || ride.RideStatus != Requested {
		return false
	}
	ride.RideStatus = Accepted
	driver.Status = Busy
	ride.Driver = driver
	return true

}

func (r *Rideservice) CompleteRide(rId int) bool {
	ride := r.Rides[rId]
	if ride == nil || ride.RideStatus != Accepted {
		return false
	}

	ride.RideStatus = Completed
	status, fare := r.CalculateFare(rId)
	fmt.Println("Fare calculated:", fare)
	if !status {
		return false
	}
	fmt.Println("Ride completed. Fare:", fare)
	r.ProcessPayment(fare)
	fmt.Println("Payment processed.")

	ride.Driver.Status = Available

	return true
}

func (r *Rideservice) CancelRide(rId int) bool {
	ride := r.Rides[rId]
	if ride == nil || ride.RideStatus != Requested {
		return false
	}
	ride.RideStatus = Cancelled
	return true
}

func (r *Rideservice) processRideRequests() bool {
	for ride := range r.Ride {
		r.mu.Lock()
		r.Rides[ride.ID] = ride
		r.mu.Unlock()
		r.NotifyDriver(ride)
	}
	return true
}
func (r *Rideservice) NotifyDriver(ride *Ride) {

	// Notify the driver about the ride request

	for _, driver := range r.Drivers {
		if driver.Status == Available && r.CalculateDistance(driver.Location, ride.From) < 5.0 {
			// Notify the driver about the ride request
			// This could be a push notification, SMS, etc.
			// For simplicity, we'll just print it to the console
			println("Notifying driver:", driver.Name, "about ride request from", ride.Passenger.Name)
			break
		}
	}
}
func (r *Rideservice) CalculateDistance(loc1, loc2 *Location) float64 {
	loc1Coord := geodist.Coord{Lat: loc1.Lattitude, Lon: loc1.Longitude}
	loc2Coord := geodist.Coord{Lat: loc2.Lattitude, Lon: loc2.Longitude}
	dist, _ := geodist.HaversineDistance(loc1Coord, loc2Coord)
	return dist
}

func (r *Rideservice) CalculateFare(rId int) (bool, float64) {
	ride := r.Rides[rId]
	if ride == nil || ride.RideStatus != Completed {
		return false, 0.0
	}
	fmt.Println("Calculating fare for ride ID:", rId)
	distance := r.CalculateDistance(ride.From, ride.To)
	fmt.Println("Distance calculated:", distance)
	// Assuming a fare rate of $1 per kilometer
	fare := 0.0
	if ride.RideType == Regular {
		fare = distance * 1.0
	} else if ride.RideType == Premium {
		fare = distance * 1.5
	}
	fmt.Println("Distance:", distance, "Fare:", fare)
	return true, fare
}

func (r *Rideservice) ProcessPayment(amount float64) bool {
	println("Processing payment of $", amount)
	return true
}
