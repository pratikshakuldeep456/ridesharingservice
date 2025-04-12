package rss

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Rideservice struct {
	Passengers map[int]*Passenger
	Rides      map[int]*Ride
	Drivers    map[int]*Driver
	Ride       chan *Ride
	FS         FareStrategy
	observers  []Observer
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
	r.Rides[ride.ID] = ride

	// go routine and notify  to driver
	r.Ride <- ride
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
	r.NotifyObservers(ride)
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
	r.NotifyObservers(ride)
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
		r.NotifyObservers(ride)
	}
	return true
}

func (r *Rideservice) CalculateFare(rId int) (bool, float64) {
	ride := r.Rides[rId]
	if ride == nil || ride.RideStatus != Completed {
		return false, 0.0
	}
	fmt.Println("Calculating fare for ride ID:", rId)
	distance := CalculateDistance(ride.From, ride.To)
	fmt.Println("Distance calculated:", distance)
	stg := NewFareStrategy(ride.RideType)
	fare := stg.CalculateFare(distance)
	fmt.Println("Fare calculated:", fare)
	fmt.Println("Distance:", distance, "Fare:", fare)
	return true, fare
}

func (r *Rideservice) ProcessPayment(amount float64) bool {
	println("Processing payment of $", amount)
	return true
}

func (r *Rideservice) RegisterObserver(observer Observer) {
	r.observers = append(r.observers, observer)
}
func (r *Rideservice) RemoveObserver(observer Observer) {
	for i, o := range r.observers {
		if o == observer {
			r.observers = append(r.observers[:i], r.observers[i+1:]...)
			break
		}
	}
}
func (r *Rideservice) NotifyObservers(ride *Ride) {
	fmt.Println("Notify")
	for _, observer := range r.observers {
		fmt.Printf("Notifying observer: %T\n", observer)
		observer.Update(ride)
	}
}
