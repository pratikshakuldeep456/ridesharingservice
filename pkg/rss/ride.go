package rss

type Ride struct {
	ID         int
	Passenger  *Passenger
	Driver     *Driver
	From       *Location
	To         *Location
	RideStatus RideStatus
	RideType   RideType
}

func NewRide(r *Ride) *Ride {
	return &Ride{
		ID:         r.ID,
		Passenger:  r.Passenger,
		Driver:     r.Driver,
		From:       r.From,
		To:         r.To,
		RideStatus: r.RideStatus,
	}
}
