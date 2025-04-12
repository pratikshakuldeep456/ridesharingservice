package rss

type Observer interface {
	Update(ride *Ride)
}

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers(ride *Ride)
}
