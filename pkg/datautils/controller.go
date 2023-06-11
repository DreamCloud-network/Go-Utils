package datautils

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func Serialize(data any) ([]byte, error) {
	var buffer bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&buffer) // Will write to network.

	err := enc.Encode(data)
	if err != nil {
		return nil, fmt.Errorf("datautils.Serialize - %w", err)
	}

	return buffer.Bytes(), nil
}

func Deserialize(data []byte, target any) error {

	var buffer bytes.Buffer
	_, err := buffer.Write(data)
	if err != nil {
		return fmt.Errorf("datautils.Deserialize - Error reading bytes to buffer: %w", err)
	}

	dec := gob.NewDecoder(&buffer)
	err = dec.Decode(target)
	if err != nil {
		return fmt.Errorf("datautils.Deserialize - Error decoding bytes: %w", err)
	}

	return nil
}
