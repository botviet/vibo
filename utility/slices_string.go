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

// StringContains of slice
func StringContains(element string, data []string) bool {
	return StringIndexOf(element, data) != -1
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

// StringRemoveEmpty .
func StringRemoveEmpty(vs []string) []string {
	return StringFilter(vs, func(s string) bool {
		return s != ""
	})
}

// StringClean .
func StringClean(vs []string) []string {
	vs = StringRemoveEmpty(vs)

	for i := range vs {
		vs[i] = strings.TrimSpace(vs[i])
	}

	return vs
}

// StringLower .
func StringLower(vs []string) []string {
	for i := range vs {
		vs[i] = strings.ToLower(vs[i])
	}

	return vs
}

// StringReplace .
func StringReplace(vs []string, old, new string) []string {

	replace := func(s string, old, new string) string {
		ws := strings.Split(s, " ")
		for i := range ws {
			if ws[i] == old {
				ws[i] = new
			}
		}
		return strings.Join(ws, " ")
	}

	vsnew := make([]string, len(vs))
	for i := range vs {
		vsnew[i] = replace(vs[i], old, new)
	}

	return vsnew
}
