package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "input"

func convertToInts(data []string) []int {
	var results []int
	for _, value := range data {
		val := strings.Trim(value, " ")
		if len(val) != 0 {
			num, err := strconv.Atoi(value)
			if err != nil {
				panic("number incorrect")
			}
			results = append(results, num)
		}
	}
	return results
}

func readLines(path string) ([]int, []board, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return nil, nil, errors.New("no data on first line")
	}
	firstline := scanner.Text()
	balls := strings.Split(firstline, ",")
	numbers := convertToInts(balls)

	if !scanner.Scan() {
		return nil, nil, errors.New("no data on second line")
	}
	var boards []board
	var butlleta board
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), " ")
		if len(line) == 0 {
			// New board
			butlleta.columnCount = make([]int, len(butlleta.Lines[0]))
			boards = append(boards, butlleta)

			butlleta = board{}
		} else {
			linenums := convertToInts(strings.Split(line, " "))
			butlleta.Lines = append(butlleta.Lines, linenums)
			butlleta.rowCount = append(butlleta.rowCount, 0)
		}

	}

	butlleta.columnCount = make([]int, len(butlleta.Lines[0]))
	boards = append(boards, butlleta)
	return numbers, boards, scanner.Err()
}

func main() {
	numbers, boards, err := readLines(filename)
	if err != nil {
		panic("File read failed")
	}

	part1 := play1(numbers, boards)
	fmt.Println("Part 1:", part1)

	for _, butlleta := range boards {
		butlleta.cleanCounters()
	}

	part2 := play2(numbers, boards)
	fmt.Println("Part 2:", part2)

}

// --- PART 1

func HasValue(values []int, value int) bool {
	for _, count := range values {
		if count == value {
			return true
		}
	}
	return false
}

type board struct {
	Lines       [][]int
	rowCount    []int
	columnCount []int
}

func (b *board) cleanCounters() {
	for i := range b.rowCount {
		b.rowCount[i] = 0

	}

	for i := range b.columnCount {
		b.columnCount[i] = 0
	}
}

func (b board) HaveLineEmpty() bool {
	return HasValue(b.rowCount, len(b.Lines))
}

func (b board) HaveColumnEmpty() bool {
	return HasValue(b.columnCount, len(b.Lines[0]))
}

func (b *board) RemoveNumber(number int) {
	var boardresult [][]int

	for row, line := range b.Lines {
		var newline []int
		for col, value := range line {
			if value != number {
				newline = append(newline, value)
			} else {
				newline = append(newline, -1)
				b.columnCount[col]++
				b.rowCount[row]++
			}
		}
		boardresult = append(boardresult, newline)
	}
	b.Lines = boardresult
}

func (b board) Sum() int {
	sum := 0
	for _, line := range b.Lines {
		for _, number := range line {
			if number > 0 {
				sum += number
			}
		}
	}
	return sum
}

func play1(numbers []int, boards []board) int {

	for _, number := range numbers {
		var newBoard []board
		for _, butlleta := range boards {
			butlleta.RemoveNumber(number)
			if butlleta.HaveLineEmpty() || butlleta.HaveColumnEmpty() {
				return number * butlleta.Sum()
			}
			newBoard = append(newBoard, butlleta)
		}
		boards = newBoard
	}
	return -1
}

// PART 2

func play2(numbers []int, boards []board) int {

	for _, number := range numbers {
		var newBoard []board
		for _, butlleta := range boards {
			butlleta.RemoveNumber(number)
			if butlleta.HaveLineEmpty() || butlleta.HaveColumnEmpty() {
				if (len(boards)) == 1 {
					return butlleta.Sum() * number
				}

			} else {
				newBoard = append(newBoard, butlleta)
			}

		}
		boards = newBoard

	}
	return -1
}
