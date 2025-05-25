// Package jsonify provides JSONify interface to marshal/unmarshal any structs.
package jsonify

import (
	"fmt"

	"github.com/mailru/easyjson"
)

type JSONify interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

// Easyjson JSON-serializer.
type easyjsonJSONify struct{}

func NewJSONify() JSONify {
	return new(easyjsonJSONify)
}

// Marshal serializes JSON with easyjson.
func (ej easyjsonJSONify) Marshal(v any) ([]byte, error) {
	if m, ok := v.(easyjson.Marshaler); ok {
		return easyjson.Marshal(m)
	}
	return nil, fmt.Errorf("the entity to serialize does not implement easyjson.Marshaler")
}

// Unmarshal deserializes JSON with easyjson.
func (ej easyjsonJSONify) Unmarshal(data []byte, v any) error {
	if um, ok := v.(easyjson.Unmarshaler); ok {
		return easyjson.Unmarshal(data, um)
	}
	return fmt.Errorf("the entity to deserialize does not implement easyjson.Unmarshaler")
}
