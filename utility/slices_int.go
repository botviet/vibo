package utility

// IntIndexOf .
func IntIndexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

// IntContains .
func IntContains(element int, data []int) bool {
	return IntIndexOf(element, data) != -1
}
