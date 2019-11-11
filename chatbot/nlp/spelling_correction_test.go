package nlp

import (
	"testing"

	"github.com/botviet/vibo/config"

	"github.com/stretchr/testify/assert"
)

func TestCorrection(t *testing.T) {
	var dic Dictionary
	config.LoadWithStore("../../storage")
	dic.loadCoOccurrence(config.StorageBigText)

	assert.Equal(t, "làm", dic.Correction("lam", []string{"gì"}, []string{"đang"}))
	assert.Equal(t, "ăn", dic.Correction("an", []string{"cơm"}, []string{"đi"}))
	assert.Equal(t, "ơi", dic.Correction("oi", []string{"bạn"}, []string{}))
	assert.Equal(t, "nghiêng", dic.Correction("nghienga", []string{"ngả"}, []string{"đi"}))
	assert.Equal(t, "chán", dic.Correction("chan", []string{"quá", "đi"}, []string{"haizz"}))
	assert.Equal(t, "đi", dic.Correction("di", []string{"chơi", "không"}, []string{"ê"}))
	assert.Equal(t, "gì", dic.Correction("gì", []string{"là", "tên", "bạn"}, []string{}))
}

func BenchmarkCorrection(b *testing.B) {
	var dic Dictionary
	config.LoadWithStore("../../storage")
	dic.loadCoOccurrence(config.StorageBigText)

	for n := 0; n < b.N; n++ {
		dic.Correction("gì", []string{"là", "tên", "bạn"}, []string{})
		dic.Correction("nghienga", []string{"ngả"}, []string{"đi"})
	}
}
