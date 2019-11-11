package config

var (
	// Storage .
	Storage string
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
	StorageModel = Storage + "/model"
	StorageBigText = Storage + "/big-text"
}
