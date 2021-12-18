package main

import (
	"testing"
)

func TestMagnitude(t *testing.T) {
	var tests = []struct {
		input  string
		result int
	}{
		{"[9,1]", 29},
		{"[1,9]", 21},
		{"[[9,1],[1,9]]", 129},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
		{"[[[[0,9],2],3],4]", 548},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			number, _ := NewSmallFishNumber(tt.input, 0)

			result := number.magnitude()

			if result != tt.result {
				t.Errorf("Dóna %d i hauria de donar %d", result, tt.result)
			}

		})
	}
}
