package nlp

import (
	"regexp"
	"strings"

	"gitlab.com/chatbot.vibo/vibo/regex"
	"gitlab.com/chatbot.vibo/vibo/utility"

	"mvdan.cc/xurls"
)

func extractNumber(doc string) []string {
	return utility.StringClean(regexp.MustCompile(`(^|\s)([0-9]+[0-9.,]+|[0-9]{1})(\s|$)`).FindAllString(doc, -1))
}

func extractTime(doc string) (times []string) {
	doc = clean(doc)

	pH := "([0-9]{1,2}|1[0-2])"
	pSep := `([\s]|)`
	pHT := "(h|giờ|g|:|AM|am|PM|pm)"
	pM := "([0-9]{1,2})"
	pMT := "(p|'|AM|am|PM|pm|)"
	pHH := pH + pSep + pHT
	pHM := pH + pSep + pHT + pSep + pM + pMT

	times = regexp.MustCompile(pHM+"|"+pHH).FindAllString(doc, -1)
	return
}

func extractDate(doc string) (dates []string) {
	doc = clean(doc)
	dates = make([]string, 0)

	//************************************
	//  handle detect year->day
	//************************************
	pYYYY := "((1[6-9]|20)[0-9]{2})" // 19xx or 20xx
	pMM := "([0-9]{1,2}|1[0-2])"
	pDD := "([0-9]{1,2})"
	pSep := `([-/\s]+)`

	pYYYYmmDD := pYYYY + pSep + pMM + pSep + pDD
	pDDmmYYYY := pDD + pSep + pMM + pSep + pYYYY

	// Oct 12th, 2018 2:15 PM
	pMMM := "[A-Za-z]{3}"
	pSep2 := `(th, | )`
	pMMMddYYY := pMMM + pSep2 + pDD + pSep2 + pYYYY

	dates = regexp.MustCompile(pYYYYmmDD+"|"+pDDmmYYYY+"|"+pMMMddYYY+"|"+pYYYY).FindAllString(doc, -1)

	// remove text in dates
	for _, date := range dates {
		doc = strings.Replace(doc, date, "", -1)
	}

	//************************************
	//  handle detect month->day
	//************************************
	pDDmm := pDD + pSep + pMM
	dates = append(dates, regexp.MustCompile(pDDmm).FindAllString(doc, -1)...)

	return
}

// refer: https://github.com/mvdan/xurls
func extractURL(doc string) []string {
	return xurls.Relaxed().FindAllString(doc, -1)
}

// temperature (°C, °F)
func extractTemperature(doc string) (tems []string) {
	doc = clean(doc)

	pN := "([0-9.,]+)"
	pSep := `([\s]|)`
	pFC := "(°C|° C|°c|° c|°F|° F|°f|° f)"

	tems = regexp.MustCompile(pN+pSep+pFC).FindAllString(doc, -1)
	return
}

func extractProbability(doc string) (probs []string) {
	doc = clean(doc)

	pN := "([0-9.,]+)"
	pSep := `([\s]|)`
	pP := "([%％])"

	probs = regexp.MustCompile(pN+pSep+pP).FindAllString(doc, -1)
	return
}

func extractEmail(doc string) []string {
	ptn := `([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`
	return regexp.MustCompile(ptn).FindAllString(clean(doc), -1)
}

func extractPhone(doc string) []string {
	ptn := `[\.\-+)(]*([0-9]{3,4})[\.\-)( ]*([0-9]{3,4})[\.\-)( ]*([0-9]{4})`
	return regexp.MustCompile(ptn).FindAllString(clean(doc), -1)
}

func extractAnnotation(doc string) []string {
	return regexp.MustCompile(regex.Annotation).FindAllString(doc, -1)
}

func clean(s string) string {
	re1 := regexp.MustCompile(regex.CleanSentence)
	re2 := regexp.MustCompile(regex.Space)

	return strings.TrimSpace(re2.ReplaceAllString(re1.ReplaceAllString(s, " "), " "))
}

func cleanWord(s string) string {
	// split by CleanWords and Space
	re1 := regexp.MustCompile(regex.CleanWords)
	re2 := regexp.MustCompile(regex.Space)

	return strings.TrimSpace(re2.ReplaceAllString(re1.ReplaceAllString(s, " "), " "))
}

func splitLatin(s string) []string {
	return regexp.MustCompile(regex.Latin).FindAllString(s, -1)
}

// SplitWords .
func SplitWords(sentence string) (ws []string) {
	return utility.StringRemoveEmpty(strings.Split(cleanWord(sentence), " "))
}
