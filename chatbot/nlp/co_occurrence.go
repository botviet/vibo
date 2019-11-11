package nlp

import (
	"math"

	log "github.com/sirupsen/logrus"
)

// CoOccurrence .
type CoOccurrence struct {
	Size   int // size window should be between 2 - 5
	Header map[string]int
	Window [][]int
}

func (co *CoOccurrence) init(size int) {
	if size < 2 {
		size = 2
		log.Warn("CoOccurrence window size minimun is 2, setting size default is 2")
	}
	co.Size = size
	co.Header = make(map[string]int)
	co.Window = make([][]int, 0)
}

func (co *CoOccurrence) updateWords(ws []string) {

	size := len(ws)
	// fix size on text small
	coSize := co.Size
	if size < coSize {
		coSize = size
	}

	for i := 0; i < size; i++ {
		if i+coSize > size {
			return
		}
		co.update(ws[i : i+coSize]...)
	}
}

func (co *CoOccurrence) update(ws ...string) {
	var size = len(ws)
	var iheader = make([]int, size)

	// update header
	for i, w := range ws {
		index, has := co.Header[w]
		if !has {
			index = len(co.Header)
			co.addOneSize()
			co.Header[w] = index
		}

		iheader[i] = index
	}

	// update window
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if x == y {
				continue
			}
			co.Window[iheader[x]][iheader[y]]++
		}
	}
}

func (co *CoOccurrence) addOneSize() {
	y := len(co.Window)
	co.Window = append(co.Window, make([]int, y))

	for x := 0; x <= y; x++ {
		// add column
		co.Window[x] = append(co.Window[x], 0)
		// add row
		co.Window[y][x] = 0
	}

}

func (co *CoOccurrence) occurrence(w1, w2 string) int {
	i1, has := co.Header[w1]
	if !has {
		return 0
	}
	i2, has := co.Header[w2]
	if i2 == -1 {
		return 0
	}

	return co.Window[i1][i2]
}

func (co *CoOccurrence) prob(w1 string, wr []string, wl []string) (occ float64) {

	for i, w := range wr {
		occ += float64(co.occurrence(w1, w)) / math.Pow10(i)
	}

	for i, w := range wl {
		occ += float64(co.occurrence(w1, w)) / math.Pow10(i)
	}

	return
}
