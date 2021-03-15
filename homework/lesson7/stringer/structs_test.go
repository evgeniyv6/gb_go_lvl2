package stringer

import (
	"encoding/json"
	"testing"
)

var (
	data = []byte(`{"RealName":"Sanka", "Login":"girl.sascha", "Status":2, "Flags": 1}`)
	u = User{}
	c = Client{}
)

// go test -v -bench=. -benchmem *.go
// go test -v -bench=. *.go

func BenchmarkDecodeStandart(b *testing.B) {
	for i:=0; i<b.N; i++ {
		_ = json.Unmarshal(data, &c)
	}
}

func BenchmarkEasyJson(b *testing.B) {
	for i:=0; i<b.N; i++ {
		_ = u.UnmarshalJSON(data)
	}
}


func BenchmarkEncodeStandart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&c)
	}
}
func BenchmarkEncodeEasyjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = u.MarshalJSON()
	}
}