package util

import "math"

// Calculate two point latitude and longitude using Harvesine formula
// this will return the distance in meters
func distance(lat1, lon1 float64, lat2, lon2 float64) float64 {
	var R float64 = 6371000.0
	var φ1 = toRadians(lat1)
	var φ2 = toRadians(lat2)
	var Δφ = toRadians(lat2 - lat1)
	var Δλ = toRadians(lon2 - lon1)

	var a = math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c

}

func toRadians(input float64) float64 {
	return input * (math.Pi / 180.0)
}
