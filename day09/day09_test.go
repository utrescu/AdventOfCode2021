package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		input    [][]int
		expected int
	}{
		{[][]int{
			{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
			{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
			{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
			{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
			{9, 8, 9, 9, 9, 6, 5, 6, 7, 8}},
			15},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Part1 %v=%d", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			hotmap := Hotmap{cells: tt.input}
			result := Part1(hotmap)
			if result != tt.expected {
				t.Errorf("rebut %d, esperava %d", result, tt.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		input    [][]int
		expected int
	}{
		{[][]int{
			{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
			{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
			{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
			{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
			{9, 8, 9, 9, 9, 6, 5, 6, 7, 8}},
			1134},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Part2 %v=%d", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			hotmap := Hotmap{cells: tt.input}
			result := Part2(hotmap)
			if result != tt.expected {
				t.Errorf("rebut %d, esperava %d", result, tt.expected)
			}
		})
	}
}
