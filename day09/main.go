package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func readLines(path string) (Hotmap, error) {
	file, err := os.Open(path)
	if err != nil {
		return Hotmap{}, err
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linearray, _ := stringArrayToInt(strings.Split(scanner.Text(), ""))
		lines = append(lines, linearray)
	}

	return Hotmap{cells: lines}, scanner.Err()
}

type Position struct {
	row int
	col int
}

func contains(positions []Position, position Position) bool {
	for _, pos := range positions {
		if pos.col == position.col && pos.row == position.row {
			return true
		}
	}
	return false
}

type Hotmap struct {
	cells [][]int
}

// Adjacents retorna les posicions adjacents a la especificada
func (h Hotmap) Adjacents(row int, col int) []Position {
	result := make([]Position, 0)

	if row > 0 {
		result = append(result, Position{row: row - 1, col: col})
	}
	if col > 0 {
		result = append(result, Position{row: row, col: col - 1})
	}

	if row < len(h.cells)-1 {
		result = append(result, Position{row: row + 1, col: col})
	}

	if col < len(h.cells[0])-1 {
		result = append(result, Position{row: row, col: col + 1})
	}
	return result
}

// AdjacentValues retorna els valors adjacents a la cel·la
func (h Hotmap) AdjacentValues(row int, col int) []int {
	result := make([]int, 0)

	adjacents := h.Adjacents(row, col)

	for _, adjacent := range adjacents {
		result = append(result, h.cells[adjacent.row][adjacent.col])
	}

	return result
}

// IsLowPoint diu si el punt és un LowPoint
func (h Hotmap) IsLowPoint(row int, col int) bool {
	data := h.AdjacentValues(row, col)
	me := h.cells[row][col]
	for _, d := range data {
		if d <= me {
			return false
		}
	}
	return true
}

// Get retorna el valor de la posició actual
func (h Hotmap) Get(p Position) int {
	return h.cells[p.row][p.col]
}

// Retorna les posicions adjacents a la posició especificada
func (h Hotmap) AdjacentsPosition(p Position) []Position {
	return h.Adjacents(p.row, p.col)
}

// Bassins Calcula el valor del "bassin" del punt
func (h Hotmap) Basins(row int, col int) int {
	if h.cells[row][col] == 9 {
		return 0
	}
	pendingNodes := map[Position]bool{{row, col}: true}
	done := make([]Position, 0)

	for len(pendingNodes) > 0 {
		newNodes := make(map[Position]bool)
		for pending := range pendingNodes {
			done = append(done, pending)
			possibleNewNodes := h.AdjacentsPosition(pending)
			for _, newadjacent := range possibleNewNodes {
				adjacentValue := h.Get(newadjacent)
				if adjacentValue < 9 && !contains(done, newadjacent) {
					newNodes[newadjacent] = true
				}
			}
		}
		pendingNodes = newNodes
	}
	return len(done)
}

func main() {

	mapaCalor, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	result1, result2 := Part(mapaCalor)
	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}

func Part(hotmap Hotmap) (int, int) {
	results := make([]int, 0)
	sum := 0
	for row := 0; row < len(hotmap.cells); row++ {
		for col := 0; col < len(hotmap.cells[0]); col++ {
			if hotmap.IsLowPoint(row, col) {
				// Part 1
				sum += hotmap.cells[row][col] + 1
				// Part 2
				results = append(results, hotmap.Basins(row, col))
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(results)))
	return sum, results[0] * results[1] * results[2]
}
