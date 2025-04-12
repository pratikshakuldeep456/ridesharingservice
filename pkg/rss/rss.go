package rss

import (
	"sync"

	"github.com/google/uuid"
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
	})
	return RideService
}

func (r *Rideservice) AddPassenger(p *Passenger) {
	r.Passengers[p.ID] = p
}

func (r *Rideservice) AddDriver(d *Driver) {
	r.Drivers[d.ID] = d
}
func (r *Rideservice) RequestRide(pId int, start, end *Location, status RideStatus) *Ride {
	r.mu.Lock()
	defer r.mu.Unlock()
	p := r.Passengers[pId]

	ride := &Ride{
		ID:         int(uuid.New().ID()),
		Passenger:  p,
		From:       start,
		To:         end,
		RideStatus: status,
	}
	r.Ride <- ride
	r.Rides[ride.ID] = ride
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
