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

func readLines(path string) ([]Line, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

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
		lines = append(lines, linia)
	}
	return lines, scanner.Err()
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
	x := l.first.x
	if l.last.x > x {
		x = l.last.x
	}
	y := l.last.y
	if l.last.y > y {
		y = l.last.y
	}
	return int(x), int(y)
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
	lines, err := readLines(filename)
	if err != nil {
		panic("File read failed")
	}

	straightLines := make([]Line, 0)
	for _, line := range lines {
		if line.first.x == line.last.x || line.first.y == line.last.y {
			straightLines = append(straightLines, line)
		}
	}

	result := Part(straightLines)
	fmt.Println("Part 1:", result)

	result2 := Part(lines)
	fmt.Println("Part 2:", result2)

}

func Part(lines []Line) int {
	var board [1000][1000]int
	maxX := 0
	maxY := 0
	for _, line := range lines {
		x, y := line.max()
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	for _, line := range lines {
		for _, p := range line.Path() {
			row := int(p.y)
			column := int(p.x)
			board[row][column]++
		}
	}

	resultat := 0
	for line := 0; line <= maxY; line++ {
		for column := 0; column <= maxX; column++ {
			if board[line][column] > 1 {
				resultat++
			}
		}
	}

	return resultat
}
