package main

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	var tests = []struct {
		input    []string
		value    string
		times    int
		expected bool
	}{
		{[]string{"a", "b"}, "a", 1, true},
		{[]string{"a", "b", "d"}, "c", 1, false},
		{[]string{"a", "b"}, "a", 2, false},
		{[]string{"a", "b"}, "b", 2, false},
		{[]string{"a", "b", "a"}, "c", 2, false},
		{[]string{"a", "b", "a"}, "a", 2, true},
		{[]string{"a", "b", "a"}, "b", 2, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("alreadyInserted%d %v %s", tt.times, tt.input, tt.value)
		t.Run(testname, func(t *testing.T) {

			result := alreadyInserted(tt.input, tt.value, tt.times)
			if result != tt.expected {
				t.Errorf("Ha donat mal resultat")
			}
		})
	}
}
