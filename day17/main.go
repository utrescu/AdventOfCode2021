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
	MAXINT = int(^uint(0) >> 1)
	MININT = -MAXINT - 1
)

// --- UTILS

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func readLines(path string) (Area, error) {
	file, err := os.Open(path)
	if err != nil {
		return Area{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	line = strings.TrimPrefix(line, "target area: ")

	separation := strings.Split(line, ", ")

	var area Area

	for _, value := range separation {

		values := strings.Split(value, ", ")
		for _, v := range values {
			coord := strings.Split(v, "=")
			numbers := strings.Split(coord[1], "..")

			switch coord[0] {
			case "x":
				area.x1, _ = stringToInt(numbers[0])
				area.x2, _ = stringToInt(numbers[1])
			case "y":
				area.y1, _ = stringToInt(numbers[0])
				area.y2, _ = stringToInt(numbers[1])
			}

		}
	}

	return area, scanner.Err()
}

// --- MAIN

const FILENAME = "input"

func main() {
	mapa, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	if result1, result2, err := Part(mapa); err != nil {
		fmt.Println("No result")
	} else {
		fmt.Println("Part 1:", result1)
		fmt.Println("Part 2:", result2)
	}
}

type Area struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

type Prove struct {
	x  int
	y  int
	vx int
	vy int
}

func (p *Prove) Move() {
	p.x += p.vx
	if p.vx != 0 {
		if p.vx > 0 {
			p.vx -= 1
		} else {
			p.vx += 1
		}
	}

	p.y += p.vy
	p.vy--
}

// Inside comprova si el prove està dins de l'àrea especificada
func (p Prove) Inside(a Area) bool {
	return p.x >= min(a.x1, a.x2) && p.x <= max(a.x1, a.x2) && p.y >= min(a.y1, a.y2) && p.y <= max(a.y1, a.y2)
}

// missed comprova si hem passat de llarg de l'àrea de destí
func missed(x int, y int, prove Prove, area Area) bool {

	if prove.vy < 0 && prove.y < min(area.y1, area.y2) {
		return true
	}

	if prove.x == max(area.x1, area.x2) {
		return false
	}

	origin2prove := abs(prove.x-x) / (prove.x - x)
	area2prove := abs(prove.x-max(area.x1, area.x2)) / (prove.x - max(area.x1, area.x2))

	return area2prove == origin2prove
}

// Launch envia el prove per veure si aconsegueix arribar al àrea de destí
// o es passa de llarg
func Launch(prove Prove, area Area) (int, error) {

	maxy := prove.y
	x := prove.x
	y := prove.y
	isMissed := false

	for !prove.Inside(area) && !isMissed {
		prove.Move()
		isMissed = missed(x, y, prove, area)

		if !isMissed && prove.y > maxy {
			maxy = prove.y
		}
	}

	if isMissed {
		return -1, errors.New("missed")
	}

	return maxy, nil
}

func Part(area Area) (int, int, error) {

	initialX := 0
	initialY := 0
	dirx := abs(area.x1-initialX) / (area.x1 - initialX)

	startX := dirx
	endX := max(area.x1-initialX, area.x2-initialX)
	startY := -max(abs(area.y1), abs(area.y2)) - initialY
	endY := -startY

	possibles := 0

	maxY := MININT

	for ly := startY; ly <= endY; ly++ {
		for lx := startX; lx <= endX; lx += dirx {
			prove := Prove{x: initialX, y: initialY, vx: lx, vy: ly}
			y, err := Launch(prove, area)
			if err == nil {
				if y > maxY {
					maxY = y
				}
				possibles++
			}
		}
	}
	return maxY, possibles, nil
}
