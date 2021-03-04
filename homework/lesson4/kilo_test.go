package main

import "testing"

func TestKiloGoroutines2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"simple", 1000},
		{"simple", 1000},
		{"simple", 1000},
		{"simple", 1000},
		{"simple", 1000},
		{"simple", 1000},
		{"simple", 1000},
		{"simple", 1000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := kiloGoroutins2()
			if got != tc.want {
				t.Fatalf("%s: expected %v; got - %v", tc.name, tc.want, got)
			}
		})
	}
}
