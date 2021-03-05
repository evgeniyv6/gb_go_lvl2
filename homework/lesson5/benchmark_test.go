package main

import (
	"testing"
)

const (
	num int = 10000000
	num10pc = (int)(float32(num) * 0.1)
	num90pc = (int)(float32(num) * 0.9)
	num50pc = (int)(float32(num) * 0.5)
)



func BenchmarkPerf1090(b *testing.B) {
	var testSl []int
	for i:=0; i<num ; i++ {
		testSl = append(testSl,i)
	}
	set := NewSet()
	b.ResetTimer()
	for i:=0; i< b.N;i++ {
		for _, el := range testSl[:num10pc] {
			set.Add(el)
		}
		for _, el := range testSl[num10pc:] {
			set.Has(el)
		}
	}

}

func BenchmarkPerf5050(b *testing.B) {
	var testSl []int
	for i:=0; i<num ; i++ {
		testSl = append(testSl,i)
	}
	set := NewSet()
	b.ResetTimer()
	for i:=0; i< b.N;i++ {
		for _, el := range testSl[:num50pc] {
			set.Add(el)
		}
		for _, el := range testSl[num50pc:] {
			set.Has(el)
		}
	}

}

func BenchmarkPerf9010(b *testing.B) {
	var testSl []int
	for i:=0; i<num ; i++ {
		testSl = append(testSl,i)
	}
	set := NewSet()
	b.ResetTimer()
	for i:=0; i< b.N;i++ {
		for _, el := range testSl[:num90pc] {
			set.Add(el)
		}
		for _, el := range testSl[num90pc:] {
			set.Has(el)
		}
	}

}

// RW

func BenchmarkPerfRW1090(b *testing.B) {
	var testSl []int
	for i:=0; i< num; i++ {
		testSl = append(testSl,i)
	}
	set := NewSetRW()
	b.ResetTimer()
	for i:=0; i< b.N;i++ {
		for _, el := range testSl[:num10pc] {
			set.AddRW(el)
		}
		for _, el := range testSl[num10pc:] {
			set.HasRW(el)
		}
	}

}

func BenchmarkPerfRW5050(b *testing.B) {
	var testSl []int
	for i:=0; i< num; i++ {
		testSl = append(testSl,i)
	}
	set := NewSetRW()
	b.ResetTimer()
	for i:=0; i< b.N;i++ {
		for _, el := range testSl[:num50pc] {
			set.AddRW(el)
		}
		for _, el := range testSl[num50pc:] {
			set.HasRW(el)
		}
	}

}

func BenchmarkPerfRW9010(b *testing.B) {
	var testSl []int
	for i:=0; i< num; i++ {
		testSl = append(testSl,i)
	}
	set := NewSetRW()
	b.ResetTimer()
	for i:=0; i< b.N;i++ {
		for _, el := range testSl[:num90pc] {
			set.AddRW(el)
		}
		for _, el := range testSl[num90pc:] {
			set.HasRW(el)
		}
	}
}