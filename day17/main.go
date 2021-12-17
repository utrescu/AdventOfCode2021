package main

import (
	"bufio"
	"day17/utils"
	"errors"
	"fmt"
	"os"
	"strings"
)

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
				area.x1, _ = utils.StringToInt(numbers[0])
				area.x2, _ = utils.StringToInt(numbers[1])
			case "y":
				area.y1, _ = utils.StringToInt(numbers[0])
				area.y2, _ = utils.StringToInt(numbers[1])
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
	return p.x >= utils.Min(a.x1, a.x2) && p.x <= utils.Max(a.x1, a.x2) && p.y >= utils.Min(a.y1, a.y2) && p.y <= utils.Max(a.y1, a.y2)
}

// missed comprova si hem passat de llarg de l'àrea de destí
func missed(x int, y int, prove Prove, area Area) bool {

	if prove.vy < 0 && prove.y < utils.Min(area.y1, area.y2) {
		return true
	}

	if prove.x == utils.Max(area.x1, area.x2) {
		return false
	}

	origin2prove := utils.Abs(prove.x-x) / (prove.x - x)
	area2prove := utils.Abs(prove.x-utils.Max(area.x1, area.x2)) / (prove.x - utils.Max(area.x1, area.x2))

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
	dirx := utils.Abs(area.x1-initialX) / (area.x1 - initialX)

	startX := dirx
	endX := utils.Max(area.x1-initialX, area.x2-initialX)
	startY := -utils.Max(utils.Abs(area.y1), utils.Abs(area.y2)) - initialY
	endY := -startY

	possibles := 0

	maxY := utils.MININT

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
