package main

import (
	"reflect"
	"testing"
)

// go test -v stack_test.go myinter_stack.go main.go

func TestNewMyInterStack(t *testing.T) {
	// LIFO
	var tests = []struct {
		name  string
		input []MyInter
		want  []MyInter
	}{
		{"mix vars", []MyInter{1, "a", 3}, []MyInter{1, "a"}},
		{"one var", []MyInter{1}, []MyInter{}},
		{"ten vars", []MyInter{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []MyInter{0, 1, 2, 3, 4, 5, 6, 7, 8}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := NewMyInterStack()
			for _, el := range tc.input {
				got.Push(el)
			}
			got.Pop()
			if !reflect.DeepEqual(tc.want, got.Items) {
				t.Fatalf("%s: expected %v; got - %v", tc.name, tc.want, got.Items)
			}
		})
	}
}
