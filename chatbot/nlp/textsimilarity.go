// calc Similarity
//	use : tf-idf + Cosine

package nlp

import (
	"errors"
	"math"

	"github.com/botviet/vibo/utility"
)

// Similarity returns the cosine similarity between two documents using
// Tf-Idf vectorization using the corpus.
func (dic *Dictionary) Similarity(a, b string) (float64, error) {

	tokensA := utility.StringLower(splitLatin(a))
	tokensB := utility.StringLower(splitLatin(b))

	combinedTokens := union(tokensA, tokensB)
	// Populate the vectors using frequency in the corpus.
	n := len(combinedTokens)
	vectorA := make([]float64, n)
	vectorB := make([]float64, n)
	for k, v := range combinedTokens {
		vectorA[k] = tfidf(v, tokensA, n, &dic.Bow)
		vectorB[k] = tfidf(v, tokensB, n, &dic.Bow)
	}

	similarity, err := Cosine(vectorA, vectorB)
	if err != nil {
		return 0.0, err
	}
	return similarity, nil
}

// Cosine returns the Cosine Similarity between two vectors.
func Cosine(a, b []float64) (float64, error) {
	count := 0
	lengthA := len(a)
	lengthB := len(b)
	if lengthA > lengthB {
		count = lengthA
	} else {
		count = lengthB
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= lengthA {
			s2 += math.Pow(b[k], 2)
			continue
		}
		if k >= lengthB {
			s1 += math.Pow(a[k], 2)
			continue
		}
		sumA += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("null vector")
	}
	return sumA / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}

func count(key string, a []string) int {
	count := 0
	for _, s := range a {
		if key == s {
			count = count + 1
		}
	}
	return count
}

func tfidf(v string, tokens []string, n int, Bow *map[string]int) float64 {
	if _, found := (*Bow)[v]; !found {
		return 0.0
	}

	tf := float64(count(v, tokens)) / float64((*Bow)[v])
	idf := math.Log(float64(n) / float64((*Bow)[v]))
	return tf * idf
}

func union(a, b []string) []string {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}
