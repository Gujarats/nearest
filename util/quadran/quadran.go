package quadran

import "errors"

type Location struct {
	Lat float64
	Lon float64
}

// asssuming that the base location value are positive.
func getQuadranPosition(baseLocation Location, inputLocation Location) (string, error) {
	if baseLocation.Lat == inputLocation.Lat && baseLocation.Lon == inputLocation.Lon {
		return "center", nil
	}

	if baseLocation.Lat <= inputLocation.Lat && baseLocation.Lon <= inputLocation.Lon {
		return "q1", nil
	} else if baseLocation.Lat >= inputLocation.Lat && baseLocation.Lon <= inputLocation.Lon {
		return "q2", nil
	} else if baseLocation.Lat >= inputLocation.Lat && baseLocation.Lon >= inputLocation.Lon {
		return "q3", nil
	} else if baseLocation.Lat <= inputLocation.Lat && baseLocation.Lon >= inputLocation.Lon {
		return "q4", nil
	}

	return "", errors.New("No quadran found")
}
