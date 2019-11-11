package utility

import (
	"bytes"
	"encoding/gob"
)

// DumpModel .
func DumpModel(path string, model interface{}) error {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.

	err := enc.Encode(model)
	if err != nil {
		return err
	}

	WriteFile(path, network.Bytes())
	return nil
}

// LoadModel .
func LoadModel(path string, model interface{}) error {
	b, err := ReadFile(path)
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(model)
	if err != nil {
		return err
	}

	return nil
}
