package utility

import "strings"

// StringFilter .
func StringFilter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// StringIndexOf index string of slice
func StringIndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

// RemoveDuplicates .
func RemoveDuplicates(elements *[]string) {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range *elements {
		if encountered[(*elements)[v]] == false {
			// Record this element as an encountered element.
			encountered[(*elements)[v]] = true
			// Append to result slice.
			result = append(result, (*elements)[v])
		}
	}

	*elements = result
}

// StringLower .
func StringLower(vs []string) []string {
	for i := range vs {
		vs[i] = strings.ToLower(vs[i])
	}

	return vs
}
