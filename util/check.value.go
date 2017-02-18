package util

func CheckValue(values ...string) bool {
	for _, value := range values {
		if value == "" {
			return false
		}
	}

	return true
}
