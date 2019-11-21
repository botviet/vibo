package config

const (
	// Uncensored file, not use for words
	Uncensored = ".uncensored"
)

var (
	// Storage .
	Storage string
	// StorageListWords .
	StorageListWords string
	// StorageBigText .
	StorageBigText string
	// StorageModel .
	StorageModel string
)

func init() {
	LoadWithStore("./storage")
}

// LoadWithStore .
// can overwrite config
func LoadWithStore(storage string) {
	Storage = storage
	StorageListWords = Storage + "/listwords"
	StorageModel = Storage + "/model"
	StorageBigText = Storage + "/big-text"
}
