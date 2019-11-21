package utility

// MapFindKeyIndex .
func MapFindKeyIndex(m map[int]string, val string) int {

	for i, v := range m {
		if v == val {
			return i
		}
	}

	return -1
}

// MapFindKeyByValues .
func MapFindKeyByValues(m map[string][]string, vals ...string) string {

	for key, vs := range m {
		for _, val := range vals {
			if StringContains(val, vs) {
				return key
			}
		}
	}

	return ""
}
