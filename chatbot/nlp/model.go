package nlp

import (
	"regexp"

	"github.com/botviet/vibo/config"
	"github.com/botviet/vibo/utility"

	log "github.com/sirupsen/logrus"
)

const (
	modelBow          = "Bow"
	modelCoOccurrence = "CoOccurrence"
)

// loadBow creates the document frequency mapping.
// pathCorpus required yml format
func (dic *Dictionary) buildModelBow(documents []string) {
	dic.Bow = make(map[string]int, 0)

	var allTokens []string

	for _, doc := range documents {
		allTokens = append(allTokens, utility.StringLower(splitLatin(doc))...)
	}

	for _, t := range allTokens {
		if dic.Bow[t] == 0 {
			dic.Bow[t] = 1
		} else {
			dic.Bow[t] = dic.Bow[t] + 1
		}
	}

	// remove word's frequency < 50
	for word, frequency := range dic.Bow {
		if frequency < 50 {
			delete(dic.Bow, word)
		}
	}

	if err := utility.DumpModel(config.StorageModel+"/"+modelBow, dic.Bow); err != nil {
		log.Error(err)
	}
}

func (dic *Dictionary) buildModelCoOccurrence(documents []string) {

	dic.CoOccurrence.init(3)

	for _, doc := range documents {
		ws := utility.StringFilter(
			utility.StringLower(splitLatin(doc)), func(w string) bool {
				_, has := dic.Bow[w]
				return has
			})

		dic.CoOccurrence.updateWords(ws)
	}

	if err := utility.DumpModel(config.StorageModel+"/"+modelCoOccurrence, dic.CoOccurrence); err != nil {
		log.Error(err)
	}
}

func splitLatin(s string) []string {
	return regexp.MustCompile(`[\p{L}]+`).FindAllString(s, -1)
}
