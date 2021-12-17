package main

import (
	"testing"
)

func TestMissed(t *testing.T) {
	var tests = []struct {
		name   string
		prove  Prove
		area   Area
		result bool
	}{
		{"1,1 dy-1", Prove{x: 1, y: 1, vx: 1, vy: 1}, Area{x1: 2, y1: 2, x2: 5, y2: 5}, false},
		{"1,1 dy+1", Prove{x: 1, y: 1, vx: 1, vy: -1}, Area{x1: 2, y1: 2, x2: 5, y2: 5}, true},
		{"1,1 dy-1", Prove{x: 1, y: 1, vx: 1, vy: -1}, Area{x1: 2, y1: -2, x2: 5, y2: -5}, false},
		{"1,1 dy+1", Prove{x: 1, y: 1, vx: 1, vy: -1}, Area{x1: 2, y1: -2, x2: 5, y2: -5}, false},
		{"ax", Prove{x: 1, y: 1}, Area{x1: 2, y1: 2, x2: 5, y2: 5}, false},
		{"ax", Prove{x: 6, y: 1}, Area{x1: 2, y1: 2, x2: 5, y2: 5}, true},
		{"ax", Prove{x: 6, y: 1}, Area{x1: 2, y1: -2, x2: 5, y2: -5}, true},
	}

	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {

			result := missed(0, 0, tt.prove, tt.area)

			if result != tt.result {
				t.Errorf("misset falla")
			}

		})
	}
}

func TestLaunch(t *testing.T) {
	var tests = []struct {
		name   string
		prove  Prove
		area   Area
		result int
	}{
		{"7,2", Prove{x: 0, y: 0, vx: 7, vy: 2}, Area{x1: 20, x2: 30, y1: -10, y2: -5}, 3},
		{"6,3", Prove{x: 0, y: 0, vx: 6, vy: 3}, Area{x1: 20, x2: 30, y1: -10, y2: -5}, 6},
		{"6,3", Prove{x: 0, y: 0, vx: 9, vy: 0}, Area{x1: 20, x2: 30, y1: -10, y2: -5}, 0},
		{"6,3", Prove{x: 0, y: 0, vx: 17, vy: -4}, Area{x1: 20, x2: 30, y1: -10, y2: -5}, -1},
		{"6,9", Prove{x: 0, y: 0, vx: 6, vy: 9}, Area{x1: 20, x2: 30, y1: -10, y2: -5}, 45},
		{"30,-10", Prove{x: 0, y: 0, vx: 30, vy: -10}, Area{x1: 20, x2: 30, y1: -10, y2: -5}, 0},
	}

	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {

			result, _ := Launch(tt.prove, tt.area)

			if result != tt.result {
				t.Errorf("Dóna %d i hauria de donar %d", result, tt.result)
			}

		})
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		name   string
		area   Area
		result int
	}{
		{"sample", Area{x1: 20, x2: 30, y1: -10, y2: -5}, 112},
	}

	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {

			_, result, _ := Part(tt.area)

			if result != tt.result {
				t.Errorf("Dóna %d i hauria de donar %d", result, tt.result)
			}

		})
	}
}
