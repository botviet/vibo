package nlp

import (
	"regexp"

	"github.com/botviet/vibo/config"
	"github.com/botviet/vibo/utility"
)

// Dictionary .
type Dictionary struct {
	Bow          map[string]int // use for tf-idf + cosine -> similarity sentences
	CoOccurrence CoOccurrence   // window Occurrence of words
}

// Load init Dictionary
// Load words from pathListWords, but not load file name is .uncensored
func (dic *Dictionary) Load(pathBigText string) {
	dic.loadBow(pathBigText)
	dic.loadCoOccurrence(pathBigText)
}

// loadBow creates the document frequency mapping.
// pathCorpus required yml format
func (dic *Dictionary) loadBow(pathText ...string) {

	err := utility.LoadModel(config.StorageModel+"/"+modelBow, &dic.Bow)
	if err != nil {
		// load big-text
		documents := loadText(pathText...)

		// load and build model file
		dic.buildModelBow(documents)
	}
}

func (dic *Dictionary) loadCoOccurrence(pathText ...string) {
	err := utility.LoadModel(config.StorageModel+"/"+modelCoOccurrence, &dic.CoOccurrence)
	if err != nil {
		// load big-text
		documents := loadText(pathText...)

		// load and build model file
		dic.buildModelCoOccurrence(documents)
	}
}

func loadText(pathText ...string) (docs []string) {
	for _, p := range pathText {
		bs, _ := utility.ReadFiles(p)
		for _, b := range bs {
			docs = append(docs, regexp.MustCompile(`[\r\n.,;]`).Split(string(b), -1)...)
		}
	}

	return
}
