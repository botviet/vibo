package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViqrToUnicode(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{
			"ngu+o+`i",
			"người",
		},
		{
			"cha'n",
			"chán",
		},
		{
			"d-i",
			"đi",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, viqrToUnicode(test.input))
	}
}

func TestUnicodeToViqr(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{
			"người",
			"ngu+o+`i",
		},
		{
			"chán",
			"cha'n",
		},
		{
			"đi",
			"d-i",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, unicodeToViqr(test.input))
	}
}
