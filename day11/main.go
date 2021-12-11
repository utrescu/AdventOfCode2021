package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// -- Utils

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

func readLines(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linearray, _ := stringArrayToInt(strings.Split(scanner.Text(), ""))
		lines = append(lines, linearray)
	}

	return lines, scanner.Err()
}

// --- Mapa

type Position struct {
	row int
	col int
}

type Mapa [][]int

func (m Mapa) Adjacents(row int, col int) []Position {
	result := make([]Position, 0)

	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {

			if r < 0 || c < 0 || r > len(m)-1 || c > len(m[0])-1 || (r == row && c == col) {

			} else {
				result = append(result, Position{row: r, col: c})
			}

		}
	}
	return result
}

func (m Mapa) Get(pos Position) int {
	return m[pos.row][pos.col]
}

func (m Mapa) Increase(pos Position) {
	m[pos.row][pos.col] += 1
}

func (m Mapa) ShowMap() {
	for _, line := range m {
		fmt.Println(line)
	}
}

func (m Mapa) Reset() Mapa {
	for row := range m {
		for col := range m[row] {
			if m[row][col] > 9 {
				m[row][col] = 0
			}
		}
	}
	return m
}

// ---

func main() {
	mapa, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	steps := 100
	result1 := Part1(mapa, steps)
	fmt.Println("Part 1:", result1)

	// Per no carregar el mapa de nou
	// segueixo des del que ja està fet
	// ja tinc `steps` passes fetes

	result2 := Part2(mapa) + steps
	fmt.Println("Part 2:", result2)
}

// -- Part 1

func Part1(mapa Mapa, max int) int {
	result := 0
	for i := 0; i < max; i++ {
		explodes := 0
		mapa, explodes = Step(mapa)
		result += explodes
	}
	return result
}

// -- Part 2

func Part2(mapa Mapa) int {
	result := 0
	max := len(mapa) * len(mapa[0])
	count := 0
	for result != max {
		mapa, result = Step(mapa)
		count++
	}
	return count
}

// -- Global

func Step(mapa Mapa) (Mapa, int) {
	explodes := 0

	for row := range mapa {
		for col := range mapa[row] {
			currentPos := Position{row, col}
			mapa.Increase(currentPos)
			if mapa.Get(currentPos) == 10 {
				// Flash
				explodes += Explode(mapa, currentPos)
			}
		}
	}

	return mapa.Reset(), explodes
}

func Explode(mapa Mapa, pos Position) int {
	count := 1

	for _, pos := range mapa.Adjacents(pos.row, pos.col) {
		mapa.Increase(pos)
		if mapa.Get(pos) == 10 {
			count += Explode(mapa, pos)
		}
	}
	return count
}
