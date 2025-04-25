package jsonify

import (
	"fmt"

	"github.com/mailru/easyjson"
)

// JSON-serializer interface
type JSONify interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// easyjson JSON-serializer
type easyjsonJSONify struct{}

// JSONify constructor
func NewJSONify() JSONify {
	return new(easyjsonJSONify)
}

// serialize JSON with easyjson
func (ej easyjsonJSONify) Marshal(v interface{}) ([]byte, error) {
	if m, ok := v.(easyjson.Marshaler); ok {
		return easyjson.Marshal(m)
	}
	return nil, fmt.Errorf("the entity to serialize does not implement easyjson.Marshaler")
}

// deserialize JSON with easyjson
func (ej easyjsonJSONify) Unmarshal(data []byte, v interface{}) error {
	if um, ok := v.(easyjson.Unmarshaler); ok {
		return easyjson.Unmarshal(data, um)
	}
	return fmt.Errorf("the entity to deserialize does not implement easyjson.Unmarshaler")
}
