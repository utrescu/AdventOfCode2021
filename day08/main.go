package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// --- Utils

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

// Matches retorna les lletres iguals d'un string
func matches(a string, b string) int {
	sum := 0
	for _, c := range strings.Split(b, "") {
		if strings.Contains(a, c) {
			sum++
		}
	}
	return sum
}

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
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

// ---- Program

func digits() []Digit {
	return []Digit{
		{number: 0, segments: "abcefg"},
		{number: 1, segments: "cf"},
		{number: 2, segments: "acdeg"},
		{number: 3, segments: "acdfg"},
		{number: 4, segments: "bcdf"},
		{number: 5, segments: "abdfg"},
		{number: 6, segments: "abdefg"},
		{number: 7, segments: "acf"},
		{number: 8, segments: "abcdefg"},
		{number: 9, segments: "abcdfg"},
	}
}

func main() {

	data, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	result1 := Part1(data)
	fmt.Println("Part 1:", result1)

	result2 := Part2(data)
	fmt.Println("Part 2:", result2)

}

type Digit struct {
	number   int
	segments string
}

func (d Digit) numSegments() int {
	return len(d.segments)
}

type Segments struct {
	wires   []string
	outputs []string
}

func Part1(segments []Segments) int {

	sum := 0
	digits := digits()
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

func Part2(segments []Segments) int {

	sum := 0
	for _, segment := range segments {
		number, _ := stringToInt(DeduceLine(segment))
		sum += number
	}
	return sum
}

func DeduceLine(segment Segments) string {
	digits := digits()

	candidates := make(map[int][]string)

	// Obtenir candidats per cada número
	for _, data := range segment.wires {
		for _, digit := range digits {
			if digit.numSegments() == len(data) {
				candidates[digit.number] = append(candidates[digit.number], data)
			}

		}

	}

	// Deduir els números
	found := make(map[string]int)
	for len(candidates) != 0 {
		newCandidates := make(map[int][]string)
		for number, candidate := range candidates {
			// Si només hi ha un candidat ja tenim la solució
			if len(candidate) == 1 {
				found[candidate[0]] = number
			} else {
				possibles := make([]string, 0)
				for _, possible := range candidate {
					// Els que ja han estat trobats ja no poden ser candidats
					_, ok := found[possible]
					if !ok {
						count := 0
						// Si no té el mateix número de lletres iguals que el correcte no és un candidat
						for value, numberfound := range found {
							if matches(possible, value) == matches(digits[number].segments, digits[numberfound].segments) {
								count++
							}
						}
						if count == len(found) {
							possibles = append(possibles, possible)
						}
					}

				}
				newCandidates[number] = possibles
			}

		}
		candidates = newCandidates

	}

	// Com que la clau pot estar desordenada, les ordeno per obtenir el dígit
	result := ""
	for _, value := range segment.outputs {
		var digit int
		valueSorted := SortStringByCharacter(value)
		for k, value2 := range found {
			keySorted := SortStringByCharacter(k)
			if keySorted == valueSorted {
				digit = value2
				break
			}
		}
		result = result + fmt.Sprintf("%d", digit)
	}

	return result
}
