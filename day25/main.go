package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}
	return lines, scanner.Err()
}

func EmptyMapa(mapa [][]string) [][]string {
	result := make([][]string, 0)
	for _, line := range mapa {
		newLine := []string{}
		for range line {
			newLine = append(newLine, ".")
		}
		result = append(result, newLine)
	}
	return result
}

func print(mapa [][]string) {
	for _, line := range mapa {
		fmt.Println(line)
	}
}

func Part1(mapa [][]string) int {
	steps := 0
	height := len(mapa)
	width := len(mapa[0])

	moves := -1
	for moves != 0 {
		newmapa := EmptyMapa(mapa)
		moves = 0
		// Move left
		for row, line := range mapa {
			for col, v := range line {
				if v == ">" {
					newcol := (col + 1) % width
					if mapa[row][newcol] == "." {
						newmapa[row][newcol] = ">"
						moves++
					} else {
						newmapa[row][col] = ">"
					}
				} else {
					if newmapa[row][col] == "." {
						newmapa[row][col] = v
					}
				}
			}
		}
		// move down
		mapa = newmapa
		newmapa = EmptyMapa(mapa)
		for row, line := range mapa {
			for col, v := range line {
				if v == "v" {
					newrow := (row + 1) % height
					if mapa[newrow][col] == "." {
						newmapa[newrow][col] = "v"
						moves++
					} else {
						newmapa[row][col] = "v"
					}
				} else {
					if newmapa[row][col] == "." {
						newmapa[row][col] = v
					}
				}
			}
		}

		mapa = newmapa
		steps++

	}

	return steps
}

const FILENAME = "input"

func main() {

	mapa, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result1 := Part1(mapa)
	fmt.Println("Part 1:", result1)

}
