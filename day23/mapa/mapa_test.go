package mapa

import (
	"testing"
)

func TestCanGoToHalf(t *testing.T) {
	var tests = []struct {
		testname string
		input1   int
		input2   int
		half     []string
		result   bool
	}{

		{"one", 0, 2, []string{"A", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."}, false},
		{"one", 1, 2, []string{"A", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."}, true},
		{"one", 3, 2, []string{".", ".", ".", "A", ".", ".", ".", ".", ".", ".", "."}, false},
		{"one", 2, 2, []string{".", ".", ".", "A", ".", ".", ".", ".", ".", ".", "."}, true},
	}

	for _, tt := range tests {
		rooms := make([][]string, 0)

		t.Run(tt.testname, func(t *testing.T) {

			mapa := Mapa{tt.half, rooms, 1}
			result := mapa.canGoToHalf(tt.input1, tt.input2)

			if result != tt.result {
				t.Errorf("Dóna %t i hauria de donar %t", result, tt.result)
			}

		})
	}
}

func TestIsEmpty(t *testing.T) {
	var tests = []struct {
		testname string

		input  []string
		result bool
	}{

		{"one", []string{".", "."}, true},
		{"one", []string{".", "Q"}, false},
		{"one", []string{"A", "."}, false},
	}

	for _, tt := range tests {

		t.Run(tt.testname, func(t *testing.T) {

			result := isEmpty(tt.input)

			if result != tt.result {
				t.Errorf("Dóna %t i hauria de donar %t", result, tt.result)
			}

		})
	}
}
