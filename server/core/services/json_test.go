package services

import (
	"github.com/mailru/easyjson"
	"testing"
)


// !!! HINT !!!
// before running benchmarks you have to add to json.go file struct MyStruct{A int B string} and generate easyjson for it


func BenchmarkDirectEasyjsonMarshal(b *testing.B) {
	obj := MyStruct{A: 10, B: "hello"}
	for i := 0; i < b.N; i++ {
		_, _ = easyjson.Marshal(obj)
	}
}

func BenchmarkCustomInterfaceMarshal(b *testing.B) {
	var obj interface{} = MyStruct{A: 10, B: "hello"}
	for i := 0; i < b.N; i++ {
		_, _ = EasyjsonEncoder(obj)
	}
}

func BenchmarkDirectEasyjsonUnmarshal(b *testing.B) {
	data := []byte(`{"A": 10, "B": "hello"}`)
	var obj MyStruct
	for i := 0; i < b.N; i++ {
		_ = easyjson.Unmarshal(data, &obj)
	}
}

func BenchmarkCustomInterfaceUnmarshal(b *testing.B) {
	data := []byte(`{"A": 10, "B": "hello"}`)
	var obj interface{}
	for i := 0; i < b.N; i++ {
		_ = EasyjsonDecoder(data, obj)
	}
}
