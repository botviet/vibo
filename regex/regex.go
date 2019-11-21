package regex

const (
	// CleanSentence .
	CleanSentence = `[　<>\|]`

	// CleanWords .
	CleanWords = `(\s|^)[.,:?!　()<>\|"]+|[.,:?!　()<>\|"]+(\s|$)`

	// Latin .
	Latin = `[\p{L}]+`

	// Space .
	Space = `\s+`

	// Annotation .
	// {{.image@img_1234}}
	Annotation = `{{.(\w+)[@|\w]+}}`
)
