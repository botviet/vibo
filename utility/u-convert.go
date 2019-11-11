package utility

import "strings"

const (
	// VIQR .
	VIQR = "VIQR"
	// UNICODE .
	UNICODE = "UNICODE"
)

var viqr []string
var viqrAccent []string

var unicode []string

func init() {
	unicode = []string{
		"À", "Á", "Â", "Ã", "È", "É", "Ê", "Ì", "Í", "Ò",
		"Ó", "Ô", "Õ", "Ù", "Ú", "Ý", "à", "á", "â", "ã",
		"è", "é", "ê", "ì", "í", "ò", "ó", "ô", "õ", "ù",
		"ú", "ý", "Ă", "ă", "Đ", "đ", "Ĩ", "ĩ", "Ũ", "ũ",
		"Ơ", "ơ", "Ư", "ư", "Ạ", "ạ", "Ả", "ả", "Ấ", "ấ",
		"Ầ", "ầ", "Ẩ", "ẩ", "Ẫ", "ẫ", "Ậ", "ậ", "Ắ", "ắ",
		"Ằ", "ằ", "Ẳ", "ẳ", "Ẵ", "ẵ", "Ặ", "ặ", "Ẹ", "ẹ",
		"Ẻ", "ẻ", "Ẽ", "ẽ", "Ế", "ế", "Ề", "ề", "Ể", "ể",
		"Ễ", "ễ", "Ệ", "ệ", "Ỉ", "ỉ", "Ị", "ị", "Ọ", "ọ",
		"Ỏ", "ỏ", "Ố", "ố", "Ồ", "ồ", "Ổ", "ổ", "Ỗ", "ỗ",
		"Ộ", "ộ", "Ớ", "ớ", "Ờ", "ờ", "Ở", "ở", "Ỡ", "ỡ",
		"Ợ", "ợ", "Ụ", "ụ", "Ủ", "ủ", "Ứ", "ứ", "Ừ", "ừ",
		"Ử", "ử", "Ữ", "ữ", "Ự", "ự", "Ỳ", "ỳ", "Ỵ", "ỵ",
		"Ỷ", "ỷ", "Ỹ", "ỹ",
	}

	viqr = []string{
		"A`", "A'", "A^", "A~", "E`", "E'", "E^", "I`", "I'", "O`",
		"O'", "O^", "O~", "U`", "U'", "Y'", "a`", "a'", "a^", "a~",
		"e`", "e'", "e^", "i`", "i'", "o`", "o'", "o^", "o~", "u`",
		"u'", "y'", "A(", "a(", "D-", "d-", "I~", "i~", "U~", "u~",
		"O+", "o+", "U+", "u+", "A.", "a.", "A?", "a?", "A^'", "a^'",
		"A^`", "a^`", "A^?", "a^?", "A^~", "a^~", "A^.", "a^.", "A('", "a('",
		"A(`", "a(`", "A(?", "a(?", "A(~", "a(~", "A(.", "a(.", "E.", "e.",
		"E?", "e?", "E~", "e~", "E^'", "e^'", "E^`", "e^`", "E^?", "e^?",
		"E^~", "e^~", "E^.", "e^.", "I?", "i?", "I.", "i.", "O.", "o.",
		"O?", "o?", "O^'", "o^'", "O^`", "o^`", "O^?", "o^?", "O^~", "o^~",
		"O^.", "o^.", "O+'", "o+'", "O+`", "o+`", "O+?", "o+?", "O+~", "o+~",
		"O+.", "o+.", "U.", "u.", "U?", "u?", "U+'", "u+'", "U+`", "u+`",
		"U+?", "u+?", "U+~", "u+~", "U+.", "u+.", "Y`", "y`", "Y.", "y.",
		"Y?", "y?", "Y~", "y~",
	}

	viqrAccent = []string{"`", "'", "?", "~", ".", "^", "(", "+", "-"}
}

// Transform .
// from, to use utility.VIQR or utility.UNICODE
func Transform(text, from, to string) string {
	return transform(text, from, to)
}

// from, to use utility.VIQR or utility.UNICODE
func transform(text, from, to string) string {

	if from == UNICODE && to == VIQR {
		return unicodeToViqr(text)
	} else if from == VIQR && to == UNICODE {
		return viqrToUnicode(text)
	}

	return ""
}

func viqrToUnicode(text string) string {
	var ts []string

	// split
	for _, c := range text {
		i := StringIndexOf(string(c), viqrAccent)
		if i == -1 || len(ts) == 0 {
			ts = append(ts, string(c))
		} else {
			ts[len(ts)-1] += string(c)
		}
	}

	// convert
	for index, c := range ts {
		i := StringIndexOf(string(c), viqr)
		if i != -1 {
			ts[index] = unicode[i]
		}
	}

	return strings.Join(ts, "")
}

func unicodeToViqr(text string) string {
	var ts []string

	for _, c := range text {
		i := StringIndexOf(string(c), unicode)
		if i == -1 {
			ts = append(ts, string(c))
			continue
		}
		ts = append(ts, viqr[i])
	}

	return strings.Join(ts, "")
}
