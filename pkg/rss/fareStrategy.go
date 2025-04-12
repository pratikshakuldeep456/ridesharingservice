package rss

type FareStrategy interface {
	CalculateFare(dis float64) float64
}

type RegularFare struct{}

func (n *RegularFare) CalculateFare(dis float64) float64 {
	return dis * 10
}

type PremiumFare struct{}

func (p *PremiumFare) CalculateFare(dis float64) float64 {
	return dis * 15
}

func NewFareStrategy(rideType RideType) FareStrategy {
	switch rideType {
	case Regular:
		return &RegularFare{}
	case Premium:
		return &PremiumFare{}
	default:
		return &RegularFare{}
	}
}
