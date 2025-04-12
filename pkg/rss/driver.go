package rss

import "fmt"

type Driver struct {
	ID       int
	Name     string
	Mobile   string
	Location *Location
	Status   Status
}

func NewDriver(d *Driver) *Driver {
	return &Driver{
		ID:       d.ID,
		Name:     d.Name,
		Mobile:   d.Mobile,
		Location: d.Location,
		Status:   d.Status,
	}
}

func (d *Driver) Update(ride *Ride) {
	if d.Status == Available {
		distance := CalculateDistance(d.Location, ride.From)
		fmt.Printf("Driver %s is %f km away from passenger %s\n", d.Name, distance, ride.Passenger.Name)
		if distance < 5.0 {
			fmt.Printf("Driver %s notified of ride request from %s\n", d.Name, ride.Passenger.Name)
			// Additional logic to accept the ride can be added here.
		}
	}
}
