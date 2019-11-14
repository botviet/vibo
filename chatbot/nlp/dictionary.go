package nlp

import (
	"regexp"

	"github.com/botviet/vibo/config"
	"github.com/botviet/vibo/utility"
)

// Dictionary .
type Dictionary struct {
	IDF          map[string]float64 // use for tf-idf + cosine -> similarity sentences
	CoOccurrence CoOccurrence       // window Occurrence of words
}

// Load words from pathListWords, but not load file name is .uncensored
func (dic *Dictionary) Load(pathBigText string) {
	dic.loadIDF(pathBigText)
	dic.loadCoOccurrence(pathBigText)
}

// loadIDF load or creates the word's idf mapping.
// pathCorpus required yml format
func (dic *Dictionary) loadIDF(pathText ...string) {

	err := utility.LoadModel(config.StorageModel+"/"+modelIDF, &dic.IDF)
	if err != nil {
		// load big-text
		documents := loadText(pathText...)

		// load and build model file
		dic.buildModelIDF(&documents)
	}
}

func (dic *Dictionary) loadCoOccurrence(pathText ...string) {
	err := utility.LoadModel(config.StorageModel+"/"+modelCoOccurrence, &dic.CoOccurrence)
	if err != nil {
		// load big-text
		documents := loadText(pathText...)

		// load and build model file
		dic.buildModelCoOccurrence(&documents)
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
