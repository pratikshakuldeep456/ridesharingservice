package rss

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
