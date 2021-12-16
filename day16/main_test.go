package main

import (
	"fmt"
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestDecode(t *testing.T) {
	var tests = []struct {
		tipe    string
		input   string
		version []int
	}{
		{"trailingzeros", "000", []int{}},
		{"literal", "110100101111111000101000", []int{6}},
		{"operator0", "00111000000000000110111101000101001010010001001000000000", []int{1, 6, 2}},
		{"operator1", "11101110000000001101010000001100100000100011000001100000", []int{7, 2, 4, 1}},
	}

	for _, tt := range tests {
		testname := tt.tipe
		t.Run(testname, func(t *testing.T) {

			version, _, _ := decode(0, tt.input)

			if !Equal(version, tt.version) {
				t.Errorf("La versió %d no és la que esperava %d", version, tt.version)
			}

		})
	}
}

func TestPart1(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"38006F45291200", 9},
		{"EE00D40C823060", 14},
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Packet %s", tt.input)
		t.Run(testname, func(t *testing.T) {

			result := Part1(tt.input)
			if result != tt.expected {
				t.Errorf("Ha donat mal resultat %d no és %d", result, tt.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"9C0141080250320F1802104A08", 1},
		{"9C005AC2F8F0", 0},
		{"F600BC2D8F", 0},
		{"D8005AC2A8F0", 1},
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Part2 %s", tt.input)
		t.Run(testname, func(t *testing.T) {

			result := Part2(tt.input)
			if result != tt.expected {
				t.Errorf("Ha donat mal resultat, %d no és %d", result, tt.expected)
			}
		})
	}
}
