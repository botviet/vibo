package nlp

import (
	"math"
	"regexp"

	"github.com/botviet/vibo/config"
	"github.com/botviet/vibo/utility"

	log "github.com/sirupsen/logrus"
)

const (
	modelIDF          = "IDF"
	modelCoOccurrence = "CoOccurrence"
)

// IDF(game) = 1 + log(Total Number Of Documents / Number Of Documents with term `game` in it)
func (dic *Dictionary) buildModelIDF(documents *[]string) {
	// Total Number Of Documents
	var totalDocument float64
	// Number Of Documents with term
	var NOD = make(map[string]float64)

	for _, doc := range *documents {
		totalDocument++

		ws := utility.StringLower(splitLatin(doc))
		utility.RemoveDuplicates(&ws)
		for _, w := range ws {
			NOD[w]++
		}
	}

	// remove word's frequency < 100
	for word, frequency := range NOD {
		if frequency < 100 {
			delete(NOD, word)
		}
	}

	// calc IDF
	for word := range NOD {
		NOD[word] = 1 + math.Log(totalDocument/NOD[word])
	}

	dic.IDF = NOD
	if err := utility.DumpModel(config.StorageModel+"/"+modelIDF, dic.IDF); err != nil {
		log.Error(err)
	}
}

func (dic *Dictionary) buildModelCoOccurrence(documents *[]string) {

	dic.CoOccurrence.init(3)

	for _, doc := range *documents {
		ws := utility.StringFilter(
			utility.StringLower(splitLatin(doc)), func(w string) bool {
				_, has := dic.IDF[w]
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
