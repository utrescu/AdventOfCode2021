package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const filename = "input"

func readLines(path string) ([]Line, int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()
	maxX := 0
	maxY := 0
	var lines []Line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var first Point
		var last Point
		var linia Line
		text := scanner.Text()
		textPoints := strings.Split(text, " -> ")
		first.createPoint(textPoints[0])
		last.createPoint(textPoints[1])
		linia.first = first
		linia.last = last
		x, y := linia.max()
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		lines = append(lines, linia)
	}
	return lines, maxX, maxY, scanner.Err()
}

// --- Point ----
type Point struct {
	x float64
	y float64
}

func getNumber(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		panic("number incorrect")
	}
	return num
}

func (p *Point) createPoint(data string) {
	points := strings.Split(data, ",")
	p.x = float64(getNumber(points[0]))
	p.y = float64(getNumber(points[1]))
}

func (p Point) IsNot(q Point) bool {
	return p.x != q.x || p.y != q.y
}

func (p Point) MoveTo(q Point) Point {
	var next Point
	if p.x-q.x == 0 {
		next.x = p.x
	} else {
		next.x = p.x + math.Abs(q.x-p.x)/(q.x-p.x)
	}

	if p.y-q.y == 0 {
		next.y = p.y
	} else {
		next.y = p.y + math.Abs(q.y-p.y)/(q.y-p.y)
	}

	return next
}

//--- Line ---
type Line struct {
	first Point
	last  Point
}

func (l Line) max() (int, int) {

	return int(math.Max(l.first.x, l.last.x)), int(math.Max(l.first.y, l.last.y))
}

func (l Line) Path() []Point {
	var result []Point
	origin := l.first
	result = append(result, origin)
	for origin.IsNot(l.last) {
		next := origin.MoveTo(l.last)
		result = append(result, next)
		origin = next
	}

	return result
}

// ---------------------------------------------------

func main() {
	lines, maxX, maxY, err := readLines(filename)
	if err != nil {
		panic("File read failed")
	}

	straightLines := make([]Line, 0)
	for _, line := range lines {
		if line.first.x == line.last.x || line.first.y == line.last.y {
			straightLines = append(straightLines, line)
		}
	}

	result := Part(straightLines, maxX, maxY)
	fmt.Println("Part 1:", result)

	result2 := Part(lines, maxX, maxY)
	fmt.Println("Part 2:", result2)

}

func Part(lines []Line, maxX int, maxY int) int {

	board := make([][]int, maxY+1)
	for i := 0; i < len(board); i++ {
		board[i] = make([]int, maxX+1)
	}

	for _, line := range lines {
		for _, p := range line.Path() {
			row := int(p.y)
			column := int(p.x)
			board[row][column]++
		}
	}

	resultat := 0
	for line := 0; line < len(board); line++ {
		for column := 0; column < len(board[line]); column++ {
			if board[line][column] > 1 {
				resultat++
			}
		}
	}

	return resultat
}
