package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAXINT   = int(^uint(0) >> 1)
	FILENAME = "input"
)

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func stringArrayToNode(stringArray []string) ([]Node, error) {
	var result []Node
	for _, value := range stringArray {
		numero, err := stringToInt(value)
		if err != nil {
			return nil, err
		}
		result = append(result, Node{cost: numero, distance: MAXINT})
	}
	return result, nil
}

func readLines(path string) (Cave, error) {
	file, err := os.Open(path)
	if err != nil {
		return Cave{}, err
	}
	defer file.Close()

	var lines [][]Node
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linearray, _ := stringArrayToNode(strings.Split(scanner.Text(), ""))
		lines = append(lines, linearray)
	}

	return Cave{lines}, scanner.Err()
}

type Position struct {
	row int
	col int
}

type Node struct {
	cost     int
	distance int
	visited  bool
}

type Cave struct {
	cells [][]Node
}

// Adjacents retorna les posicions adjacents a la especificada
func (h Cave) Adjacents(row int, col int) []Position {
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

// Retorna les posicions adjacents a la posició especificada
func (h Cave) AdjacentsPosition(p Position) []Position {
	return h.Adjacents(p.row, p.col)
}

func (c Cave) Get(pos Position) Node {
	return c.cells[pos.row][pos.col]
}

func (c *Cave) Set(pos Position, distance int) {
	c.cells[pos.row][pos.col].distance = distance
}

func (c *Cave) Visited(pos Position) {
	c.cells[pos.row][pos.col].visited = true
}

func main() {

	cave, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result1 := Part1(cave)
	fmt.Println("Part 1:", result1)

	result2 := Part2(cave, 5)
	fmt.Println("Part 2:", result2)
}

// --- Part 1 Dijkstra

func LocateLowestNotVisited(cave Cave) (Position, error) {
	result := Position{-1, -1}
	minDistance := MAXINT

	for row, line := range cave.cells {
		for col, node := range line {
			if !node.visited && node.distance < minDistance {
				minDistance = node.distance
				result = Position{row, col}
			}
		}
	}

	if minDistance == MAXINT {
		return Position{}, errors.New("no more")
	}
	return result, nil
}

func Part1(cave Cave) int {

	position := Position{0, 0}
	cave.Set(position, 0)

	var err error

	for err == nil {
		node := cave.Get(position)
		cave.Visited(position)

		posveins := cave.AdjacentsPosition(position)
		for _, posvei := range posveins {
			vei := cave.Get(posvei)
			if node.distance+vei.cost < vei.distance {
				cave.Set(posvei, node.distance+vei.cost)
			}
		}

		position, err = LocateLowestNotVisited(cave)

	}

	lastNode := Position{len(cave.cells) - 1, len(cave.cells[0]) - 1}
	return cave.Get(lastNode).distance
}

// -- Part 2

func Part2(cave Cave, expansion int) int {
	maxiCave := ExpandMap(cave, expansion)
	return Part1(maxiCave)

}

func ExpandMap(cave Cave, expansion int) Cave {
	originalWidth := len(cave.cells[0])

	newCells := make([][]Node, 0)

	for row := range cave.cells {
		newrow := make([]Node, originalWidth*expansion)
		for col, value := range cave.cells[row] {
			for i := 0; i < expansion; i++ {
				cost := (value.cost + i)
				if cost > 9 {
					cost = cost - 9
				}
				newrow[col+i*originalWidth] = Node{cost: cost, distance: MAXINT}
			}
		}
		newCells = append(newCells, newrow)
	}

	expectedHeight := len(cave.cells) * expansion

	for len(newCells) < expectedHeight {

		currentCell := len(newCells)
		previousCell := currentCell - originalWidth

		newRow := make([]Node, 0)
		previousrow := newCells[previousCell]
		for _, value := range previousrow {
			cost := (value.cost + 1)
			if cost == 10 {
				cost = 1
			}
			newRow = append(newRow, Node{cost: cost, distance: MAXINT})
		}
		newCells = append(newCells, newRow)

	}

	return Cave{cells: newCells}
}
