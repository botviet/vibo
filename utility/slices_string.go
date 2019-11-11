package utility

// StringIndexOf index string of slice
func StringIndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
