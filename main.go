package main

import (
	"fmt"
	"pratikshakuldeep456/ridesharingservice/pkg/rss"
	"time"
)

func main() {
	loc1 := &rss.Location{Latitude: 12.9716, Longitude: 77.5946}
	loc2 := &rss.Location{Latitude: 12.9716, Longitude: 77.5946}

	dist := rss.CalculateDistance(loc1, loc2)
	fmt.Printf("Test Distance: %.4f km\n", dist) // Should print 0.0000

	Service := rss.GetRideService()
	driver := &rss.Driver{
		ID:       1,
		Name:     "John Doe",
		Mobile:   "1234567890",
		Location: &rss.Location{Longitude: 77.5946, Latitude: 12.9716},
		Status:   rss.Available,
	}

	passenger := &rss.Passenger{
		ID:       1,
		Name:     "Jane Doe",
		Mobile:   "0987654321",
		Location: &rss.Location{Longitude: 77.5946, Latitude: 12.9716}, // Bangalore (same location)
	}
	Service.AddDriver(driver)

	fmt.Println("Driver Service: ", Service)
	fmt.Println("Driver Location:", driver.Location)
	fmt.Println("Passenger Location:", passenger.Location)
	Service.AddPassenger(passenger)
	start := &rss.Location{Latitude: 12.9716, Longitude: 77.5946} // Pickup in Bangalore
	end := &rss.Location{Latitude: 12.2958, Longitude: 76.6396}   // Dropoff in Mysore
	ride := Service.RequestRide(passenger.ID, end, start, rss.Requested, rss.Regular)
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
