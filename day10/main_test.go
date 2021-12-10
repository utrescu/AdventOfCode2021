package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestProcessChunk(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"{([(<{}[<>[]}>{[]{[(<()>", 1197},
		{"[[<[([]))<([[{}[[()]]]", 3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Part1 %v=%d", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			line := strings.Split(tt.input, "")
			result := ProcessChunk(line)
			if result != tt.expected {
				t.Errorf("rebut %d, esperava %d", result, tt.expected)
			}
		})
	}
}

func TestCompleteChunk(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"[({(<(())[]>[[{[]{<()<>>", "}}]])})]"},
		{"<{([{{}}[<[[[<>{}]]]>[]]", "])}>"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Complete %s=%s", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			line := strings.Split(tt.input, "")
			result := CompleteChunk(line)

			resultSlice := strings.Split(tt.expected, "")
			if len(result) != len(resultSlice) {
				t.Errorf("Mida dels resultats diferents %s != %s", result, resultSlice)
			}

			for i := 0; i < len(result); i++ {
				if result[i] != resultSlice[i] {
					t.Errorf("Error a posicio %d => rebut %s, esperava %s", i, result, tt.expected)
				}
			}
		})
	}
}

func TestCompleteScore(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"}}]])})]", 288957},
		{")}>]})", 5566},
		{"}}>}>))))", 1480781},
		{"]]}}]}]}>", 995444},
		{"])}>", 294},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("CompleteScore %s=%d", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			line := strings.Split(tt.input, "")
			result := GetCompleteScore(line)

			if result != tt.expected {
				t.Errorf("rebut %d, esperava %d", result, tt.expected)
			}
		})
	}
}
