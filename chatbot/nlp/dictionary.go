package nlp

import (
	"log"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/botviet/vibo/config"
	"github.com/botviet/vibo/utility"
)

// Dictionary .
type Dictionary struct {
	ListWords    map[string][]int   // [word] []index of type_word
	ListType     map[int]string     // save map[index]type_word is the file name in ./storage/listwords
	IDF          map[string]float64 // use for tf-idf + cosine -> similarity sentences
	CoOccurrence CoOccurrence       // window Occurrence of words
}

func (dic *Dictionary) init() {
	dic.ListWords = make(map[string][]int, 0)
	dic.ListType = make(map[int]string, 0)
}

// DefaultLoad init Dictionary
// 	defautl call func Load
func (dic *Dictionary) DefaultLoad() {
	dic.Load(
		config.StorageListWords,
		config.StorageBigText,
	)
}

// Load words from pathListWords, but not load file name is .uncensored
func (dic *Dictionary) Load(pathListWords, pathBigText string) {
	dic.init()
	dic.loadListWords(pathListWords)
	dic.loadIDF(pathBigText)
	dic.loadCoOccurrence(pathBigText)
}

// load diction words from pathListWords
// 	normal ner will get word and ToLower word
// 	ner's first letter is uppercase will not has ToLower word (ex: Name_girl, Name_boy)
func (dic *Dictionary) loadListWords(pathListWords string) {

	bs, err := utility.ReadFiles(pathListWords)
	if err != nil {
		log.Fatal(err)
	}

	// sort paths (because map not order)
	var paths []string
	for path := range bs {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	for _, path := range paths {
		b := bs[path]
		fileName := utility.BasePath(path)
		// not use for words
		if config.Uncensored == fileName {
			continue
		}

		// dic.ListType save map index - type
		Type := strings.Trim(fileName, ".")
		// is normal ner
		isNNer := unicode.IsLower(rune(Type[0]))

		index := utility.MapFindKeyIndex(dic.ListType, Type)
		if index == -1 {
			index = len(dic.ListType)
			dic.ListType[index] = Type
		}

		words := strings.Split(strings.Replace(string(b), "\r\n", "\n", -1), "\n")
		// dic.ListWords save map word - type (index)
		for _, w := range words {
			if w == "" {
				continue
			}

			if !utility.IntContains(index, dic.ListWords[w]) {
				dic.ListWords[w] = append(dic.ListWords[w], index)
			}

			if isNNer {
				wlower := strings.ToLower(w)
				if w != wlower && !utility.IntContains(index, dic.ListWords[wlower]) {
					dic.ListWords[wlower] = append(dic.ListWords[wlower], index)
				}
			}
		}
	}
}

// TypeString word from type []int -> type []string
func (dic *Dictionary) TypeString(word string) (ts []string) {
	if _, found := dic.ListWords[word]; !found {
		return
	}

	for _, typ := range dic.ListWords[word] {
		ts = append(ts, dic.ListType[typ])
	}

	return
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
