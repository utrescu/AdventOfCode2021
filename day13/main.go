package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILENAME = "input"

func getNumber(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		panic("number incorrect")
	}
	return num
}

func readLines(path string) (map[Point]string, []Fold, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	readPoints := true
	points := make(map[Point]string)
	folds := make([]Fold, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		text := scanner.Text()
		if len(text) == 0 {
			readPoints = false
		} else {
			if readPoints {
				textPoints := strings.Split(text, ",")
				punt := Point{getNumber(textPoints[0]), getNumber(textPoints[1])}
				points[punt] = "x"
			} else {
				text := strings.TrimPrefix(text, "fold along ")
				textPoints := strings.Split(text, "=")
				folds = append(folds, Fold{dir: textPoints[0], value: getNumber(textPoints[1])})
			}
		}

	}
	return points, folds, scanner.Err()
}

// --- Point ----
type Point struct {
	x int
	y int
}

type Fold struct {
	dir   string
	value int
}

// ---------------------------------------------------

func main() {
	punts, folds, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result := Part1(punts, folds[0])
	fmt.Println("Part 1:", result)

	fmt.Println()
	fmt.Println("Part 2:")
	Part2(punts, folds)

}

// --- Common

const MAXINT = int(^uint(0) >> 1)

// FoldIn: Plega el paper pel lloc indicat i retorna el nou
func FoldIn(punts map[Point]string, fold Fold) map[Point]string {
	plegay := 1000
	plegax := 1000

	nousPunts := make(map[Point]string)
	if fold.dir == "x" {
		plegax = fold.value
		plegay = MAXINT
	} else {
		plegay = fold.value
		plegax = MAXINT
	}
	for punt := range punts {
		x := punt.x
		y := punt.y
		if punt.x > plegax {
			x = plegax - (x - plegax)
		}
		if punt.y > plegay {
			y = plegay - (y - plegay)
		}

		nouPunt := Point{x, y}
		if _, ok := nousPunts[nouPunt]; !ok {
			nousPunts[nouPunt] = "x"
		}
	}
	return nousPunts
}

// -- Part 1

func Part1(punts map[Point]string, fold Fold) int {

	return len(FoldIn(punts, fold))
}

// --- Part 2

func Part2(punts map[Point]string, folds []Fold) {

	for _, fold := range folds {
		punts = FoldIn(punts, fold)
	}
	Print(punts)

}

// Print pinta el resultat per pantalla
func Print(punts map[Point]string) {
	var maxx, maxy int
	// locate max
	for punt := range punts {
		if punt.x > maxx {
			maxx = punt.x
		}
		if punt.y > maxy {
			maxy = punt.y
		}
	}

	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			punt := Point{x, y}
			if _, ok := punts[punt]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
