package jsonify

import (
	"fmt"

	"github.com/mailru/easyjson"
)


// интерфейс JSON-сериализатора
type JSONify interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}


// easyjson JSON-сериализатор
type easyjsonJSONify struct {}

// конструктор для типа интерфейса JSONify
func NewJSONify() JSONify {
	return new(easyjsonJSONify)
}

// для сериализации JSON с помощью easyjson
func (this easyjsonJSONify) Marshal(v interface{}) ([]byte, error) {
	if m, ok := v.(easyjson.Marshaler); ok {
		return easyjson.Marshal(m)
	}
	return nil, fmt.Errorf("the entity to serialize does not implement easyjson.Marshaler")
}

// для десериализации JSON с помощью easyjson
func (this easyjsonJSONify) Unmarshal(data []byte, v interface{}) error {
	if um, ok := v.(easyjson.Unmarshaler); ok {
		return easyjson.Unmarshal(data, um)
	}
	return fmt.Errorf("the entity to serialize does not implement easyjson.Unmarshaler")
}
