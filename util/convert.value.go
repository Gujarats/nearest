package util

import "strconv"

// Convert string to float64.
func ConvertToFloat64(inputs ...string) []float64 {
	var result []float64
	for _, value := range inputs {
		parseValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, parseValue)
	}

	return result
}
