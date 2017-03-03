package util

import "strconv"

// Convert string to float64.
func ConvertToFloat64(inputs ...string) ([]float64, error) {
	var result []float64
	for _, value := range inputs {
		parseValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, parseValue)
	}

	return result, nil
}

// convert string to int
func ConvertToInt64(inputs ...string) ([]int64, error) {
	var result []int64
	for _, value := range inputs {
		parseValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, parseValue)
	}

	return result, nil
}
