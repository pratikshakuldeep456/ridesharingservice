package rss

type Status string

const (
	Available Status = "available"
	Busy      Status = "busy"
)

type RideStatus string

const (
	Requested RideStatus = "requested"
	Accepted  RideStatus = "accepted"
	Completed RideStatus = "completed"
	Cancelled RideStatus = "cancelled"
	// Ongoing   RideStatus = "ongoing"
)

type RideType string

const (
	Regular RideType = "regular"
	Premium RideType = "premium"
)
