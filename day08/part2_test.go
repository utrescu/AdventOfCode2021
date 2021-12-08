package main

import (
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {
	var tests = []struct {
		input    Segments
		expected string
	}{
		{Segments{wires: []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}, outputs: []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"}}, "5353"},
		{Segments{wires: []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}, outputs: []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}}, "8394"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Part2 %v=%s", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			result := DeduceLine(tt.input)
			if result != tt.expected {
				t.Errorf("rebut %s, esperava %s", result, tt.expected)
			}
		})
	}
}
