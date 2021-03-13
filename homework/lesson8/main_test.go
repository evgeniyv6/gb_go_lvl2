package main

import (
	"strings"
	"testing"
)

var (
	ds = []string{"a","b","c","d","e"}
)

func main() {

}

func BenchmarkJoin(b *testing.B) {
	for i:=0; i<b.N;i++ {
		s := strings.Join(ds, ",")
		_ = s
	}
}
func BenchmarkBuilder(b *testing.B) {
	sb:=strings.Builder{}
	sb.Grow(2*len(ds)-1)

	b.ResetTimer()

	for i:=0; i<b.N;i++ {
		sb.WriteString(ds[0])
		sb.WriteString(",")
		sb.WriteString(ds[1])
		sb.WriteString(",")
		sb.WriteString(ds[2])
		sb.WriteString(",")
		sb.WriteString(ds[3])
		sb.WriteString(",")
		sb.WriteString(ds[4])
		s:=sb.String()
		_=s
	}
}
