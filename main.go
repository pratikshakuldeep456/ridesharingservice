package main

import (
	"fmt"
	"pratikshakuldeep456/ridesharingservice/pkg/rss"
	"time"
)

func main() {

	Service := rss.GetRideService()
	driver := &rss.Driver{
		ID:       1,
		Name:     "John Doe",
		Mobile:   "1234567890",
		Location: &rss.Location{Lattitude: 12.9716, Longitude: 77.5946},
		Status:   rss.Available,
	}
	Service.AddDriver(driver)

	fmt.Println("Driver Service: ", Service)

	passenger := &rss.Passenger{
		ID:       1,
		Name:     "Jane Doe",
		Mobile:   "0987654321",
		Location: &rss.Location{Lattitude: 12.9716, Longitude: 77.5946},
	}
	Service.AddPassenger(passenger)
	start := &rss.Location{Lattitude: 12.9716, Longitude: 77.5946} // Bangalore
	end := &rss.Location{Lattitude: 12.2958, Longitude: 76.6396}   // Mysore

	ride := Service.RequestRide(passenger.ID, start, end, rss.Requested, rss.Regular)
	time.Sleep(1 * time.Second)

	// Simulate ride acceptance
	accepted := Service.AcceptRide(driver.ID, ride.ID)
	fmt.Println("Ride acceptance status:", accepted)
	if accepted {
		fmt.Println("Ride accepted by driver:", driver.Name)
	} else {
		fmt.Println("Ride acceptance failed")
	}

	// Simulate ride completion
	completed := Service.CompleteRide(ride.ID)
	if completed {
		fmt.Println("Ride completed successfully.")
	} else {
		fmt.Println("Ride completion failed.")
	}
}
