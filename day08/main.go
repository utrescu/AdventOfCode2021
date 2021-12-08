package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func stringArrayToInt(stringArray []string) ([]int, error) {
	var result []int
	for _, value := range stringArray {
		numero, err := stringToInt(value)
		if err != nil {
			return nil, err
		}
		result = append(result, numero)
	}
	return result, nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]Segments, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Segments
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " | ")
		lines = append(lines, Segments{wires: strings.Split(data[0], " "), outputs: strings.Split(data[1], " ")})
	}
	return lines, scanner.Err()
}

func main() {

	digits := []Digit{
		{number: 0, segments: []string{"a", "b", "c", "e", "f", "g"}},
		{number: 1, segments: []string{"c", "f"}},
		{number: 2, segments: []string{"a", "c", "d", "e", "g"}},
		{number: 3, segments: []string{"a", "c", "d", "f", "g"}},
		{number: 4, segments: []string{"b", "c", "d", "f"}},
		{number: 5, segments: []string{"a", "b", "d", "f", "g"}},
		{number: 6, segments: []string{"a", "b", "d", "e", "f", "g"}},
		{number: 6, segments: []string{"a", "c", "f"}},
		{number: 7, segments: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{number: 8, segments: []string{"a", "b", "c", "e", "f", "g"}},
		{number: 9, segments: []string{"a", "b", "c", "d", "f", "g"}},
	}

	data, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	result1 := Part1(data, digits)

	fmt.Println("Part 1:", result1)

}

func Part1(segments []Segments, digits []Digit) int {

	sum := 0
	for _, segment := range segments {
		for _, output := range segment.outputs {
			switch len(output) {
			case digits[1].numSegments():
				sum++
			case digits[4].numSegments():
				sum++
			case digits[7].numSegments():
				sum++
			case digits[8].numSegments():
				sum++
			}
		}
	}
	return sum
}

type Digit struct {
	number   int
	segments []string
}

func (d Digit) numSegments() int {
	return len(d.segments)
}

type Segments struct {
	wires   []string
	outputs []string
}
