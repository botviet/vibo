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

// StringLower .
func StringLower(vs []string) []string {
	for i := range vs {
		vs[i] = strings.ToLower(vs[i])
	}

	return vs
}
