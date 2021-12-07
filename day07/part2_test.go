package main

import (
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{1, 5, 2}, 7},
		{[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 168},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Part2 %v=%d", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			result := Part2(tt.input)
			if result != tt.expected {
				t.Errorf("rebut %d, esperava %d", result, tt.expected)
			}
		})
	}
}
