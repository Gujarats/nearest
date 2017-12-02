package util

import (
	"math"
	"sync"
)

var wg sync.WaitGroup
var mutex = &sync.Mutex{}

// Calculate two point latitude and longitude using Harvesine formula
// this will return the distance in meters
func Distance(lat1, lon1 float64, lat2, lon2 float64) float64 {
	var R float64 = 6371000.0

	var φ1 = toRadians(lat1)
	var φ2 = toRadians(lat2)

	var Δφ = toRadians(lat2 - lat1)
	var Δλ = toRadians(lon2 - lon1)

	var a = math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*math.Sin(Δλ/2)*math.Sin(Δλ/2)

	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c

}

func toRadians(input float64) float64 {
	return input * (math.Pi / 180.0)
}

type Location struct {
	Lat float64
	Lon float64
}

type ResultLocation struct {
	RLocation Location
	Distance  float64
}

// Return nearest location with the distance in meters
func GetNearestLocation(myLoc Location, othersLoc []Location) ResultLocation {
	var finalLoc ResultLocation
	finalLoc.Distance = math.MaxFloat64

	rLocations := make(chan ResultLocation, len(othersLoc))

	for _, otherLoc := range othersLoc {
		wg.Add(1)

		// return the location with the calculated distance
		// add the argument to anonymous function to avoid race condition
		go func(loc1 Location, loc2 Location) {
			calDistance := Distance(loc1.Lat, loc1.Lon, loc2.Lat, loc2.Lon)
			rLocation := ResultLocation{
				Distance:  calDistance,
				RLocation: loc2,
			}
			rLocations <- rLocation
			wg.Done()
		}(myLoc, otherLoc)

	}

	wg.Wait()
	close(rLocations)

	for rLocation := range rLocations {
		if finalLoc.Distance > rLocation.Distance {
			finalLoc = rLocation
		}
	}

	return finalLoc
}
