// spelling correction
//	use : Bayes P(c|w) = P(c). P(w|c) / P(w)
//			- P(c)		: xác xuất của từ c
//			- P(w|c)	: (Error Model) xác xuất người dùng nhập nhầm w thay vì c (như mong muốn).
//						=> Tìm từ c có P(c|x) cao nhất trong các từ sai 1 or 2 vị trí so với w (x là từ đứng trước)
//			- P(w)		: (Language Model) xác xuất nhập từ w.
//			Vì P(w) không đổi với các c khác nhau, để tìm P(c|w) lớn nhất, ta cần tìm c sao cho P(c).P(w|c) lớn nhất.

package nlp

import (
	"strings"

	"github.com/botviet/vibo/utility"
)

// Correction word with sideWord
func (dic *Dictionary) Correction(word string, wr []string, wl []string) string {

	// Từ dài nhất trong tiếng Việt gồm 7 chữ cái là từ "nghiêng"
	if len(word) > 7 {
		word = word[:7]
	}

	if len(wr) == 0 && len(wl) == 0 {
		return word
	}

	return correct(&dic.CoOccurrence, strings.ToLower(word), utility.StringLower(wr), utility.StringLower(wl))
}

// editing distance
func edits(word string, alphabet string) (ws []string) {
	type Pair struct{ a, b string }
	var splits []Pair
	for i := 0; i < len(word)+1; i++ {
		splits = append(splits, Pair{word[:i], word[i:]})
	}

	for _, s := range splits {
		if len(s.b) > 0 {
			ws = append(ws, s.a+s.b[1:])
		}
		if len(s.b) > 1 {
			ws = append(ws, s.a+string(s.b[1])+string(s.b[0])+s.b[2:])
		}
		for _, c := range alphabet {
			if len(s.b) > 0 {
				ws = append(ws, s.a+string(c)+s.b[1:])
			}
		}
		for _, c := range alphabet {
			ws = append(ws, s.a+string(c)+s.b)
		}
	}

	return
}

func edits0(word string) []string {
	return edits(word, "`'?~.^(+-")
}

func edits1(word string) []string {
	return edits(word, "abcdefghijklmnopqrstuvwxyz`'?~.^(+-")
}

func edits2(word string) (ws2 []string) {
	ws := edits1(word)

	for _, w := range ws {
		ws2 = append(ws2, edits1(w)...)
	}

	return
}

func best(model *CoOccurrence, edits func(string) []string, word string, wr []string, wl []string) string {
	word = utility.Transform(word, utility.UNICODE, utility.VIQR)

	ws := edits(word)

	maxProb := 0.0
	correction := ""
	for _, word := range ws {
		word = utility.Transform(word, utility.VIQR, utility.UNICODE)
		if prob := (*model).prob(word, wr, wl); prob > maxProb {
			maxProb, correction = prob, word
		}
	}
	return correction
}

func correct(model *CoOccurrence, word string, wr []string, wl []string) string {
	if correction := best(model, edits0, word, wr, wl); correction != "" {
		return correction
	}
	if correction := best(model, edits1, word, wr, wl); correction != "" {
		return correction
	}
	if correction := best(model, edits2, word, wr, wl); correction != "" {
		return correction
	}
	return word
}
