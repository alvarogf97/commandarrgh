package utils

// Search index of the given value in the given slice
func IndexOf(array []string, value string) (int, bool) {
	for i, v := range array {
		if v == value {
			return i, true
		}
	}
	return -1, false
}
