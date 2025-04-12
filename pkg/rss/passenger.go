package rss

type Passenger struct {
	ID       int
	Name     string
	Mobile   string
	Location *Location
}

func NewPassenger(p *Passenger) *Passenger {
	return &Passenger{
		ID:       p.ID,
		Name:     p.Name,
		Mobile:   p.Mobile,
		Location: p.Location,
	}
}
