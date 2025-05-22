package jsonify

import (
	"testing"
)


func TestJSONifyInterface(t *testing.T) {
	t.Log("Try to init JSONify")

	jsonify := NewJSONify()
	t.Logf("JSONify type: %T", jsonify)

	_, err := jsonify.Marshal(struct{
		Name string
		Age int
	}{
		Name: "Bob",
		Age: 25,
	})
	t.Logf("Must be easyjson error: %v", err)
}
