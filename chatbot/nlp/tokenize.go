package nlp

import (
	"strings"

	"gitlab.com/chatbot.vibo/vibo/frame"
	"gitlab.com/chatbot.vibo/vibo/utility"
)

const (
	tokenizeDate        = "date"
	tokenizeTime        = "time"
	tokenizeURL         = "url"
	tokenizeTemperature = "temperature"
	tokenizeProbability = "probability"
	tokenizeEmail       = "email"
	tokenizePhone       = "phone"
	tokenizeNumber      = "number"
)

/*
WordTokenize from listWords , stopWords
@algorithm: Maximum Matching

	detect date		(ddMMyyy | ddMM | int + ngày)
	detect time		(int + [h | giờ] + int + [' | phút])
	detect number
	detect URL
	detect temperature (°C, ° c, °F, ° f)
	detect probability (60%)
*/
func (dic *Dictionary) WordTokenize(input []string) (wordTokenize map[string][]string, cantReadableWord []string) {

	// extract special tokenize
	// date, time, number
	wordTokenize = SpecialTokenize(input)
	// remove extracted words
	for w := range wordTokenize {
		input = utility.StringReplace(input, w, "")
	}

	for _, sentence := range input {

		// sentence
		ws := SplitWords(sentence)

		size := len(ws)
		index := size
		for size != 0 {
			// up speed
			// độ dài của từ vựng việt nam dài nhất tầm 10 từ đơn,
			// khi index quá lớn sẽ gây phí phạm xử lý.
			// set maximum index là 10
			if index > 10 {
				index = 10
			}

			testWord := strings.Join(ws[:index], " ")

			// not found word, next find word
			if index == 0 {

				if !utility.StringContains(ws[0], cantReadableWord) {
					cantReadableWord = append(cantReadableWord, ws[0])
				}

				ws = ws[1:]
				size = len(ws)
				index = size
				continue
			}

			// found it, words in dictionary
			if _, found := dic.ListWords[testWord]; found {
				wordTokenize[testWord] = dic.TypeString(testWord)
				ws = ws[index:]
				size = len(ws)
				index = size
				continue
			}

			index--
		}
	}

	return
}

// SpecialTokenize .
// extract special tokenize
// date, time, URL, temperature, probability, number
func SpecialTokenize(input []string) (specTokenize map[string][]string) {
	specTokenize = make(map[string][]string, 0)

	doc := strings.Join(input, " ")

	removeWords := func(ws []string) {
		for _, w := range ws {
			doc = strings.Replace(doc, w, "", -1)
		}
	}

	// extract date
	dates := extractDate(doc)
	removeWords(dates)
	for _, date := range dates {
		specTokenize[strings.TrimSpace(date)] = []string{tokenizeDate}
	}

	// extract time
	times := extractTime(doc)
	removeWords(times)
	for _, time := range times {
		specTokenize[strings.TrimSpace(time)] = []string{tokenizeTime}
	}

	emails := extractEmail(doc)
	removeWords(emails)
	for _, email := range emails {
		specTokenize[strings.TrimSpace(email)] = []string{tokenizeEmail}
	}

	// extract URL
	urls := extractURL(doc)
	removeWords(urls)
	for _, url := range urls {
		specTokenize[strings.TrimSpace(url)] = []string{tokenizeURL}
	}

	phones := extractPhone(doc)
	removeWords(phones)
	for _, phone := range phones {
		specTokenize[strings.TrimSpace(phone)] = []string{tokenizePhone}
	}

	tems := extractTemperature(doc)
	removeWords(tems)
	for _, tem := range tems {
		specTokenize[strings.TrimSpace(tem)] = []string{tokenizeTemperature}
	}

	probs := extractProbability(doc)
	removeWords(probs)
	for _, prob := range probs {
		specTokenize[strings.TrimSpace(prob)] = []string{tokenizeProbability}
	}

	annos := extractAnnotation(doc)
	removeWords(annos)
	for _, anno := range annos {
		specTokenize[strings.TrimSpace(anno)] = []string{frame.ANNOTATION}
	}

	// extract number
	number := extractNumber(doc)
	removeWords(number)
	for _, num := range number {
		specTokenize[strings.TrimSpace(num)] = []string{tokenizeNumber}
	}

	return
}
