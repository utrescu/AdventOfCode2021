package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type location struct {
	position int
	depth    int
	aim      int
}

type move struct {
	moviment string
	units    int
}

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, " ")
	return strconv.Atoi(nonFractionalPart[0])
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]move, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []move
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		number, _ := stringToInt(parts[1])

		lines = append(lines, move{parts[0], number})
	}
	return lines, scanner.Err()
}

func main() {

	moves, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	part1 := moveit(moves)
	fmt.Println("Part 1:", part1.depth*part1.position, "(", part1, ")")

	part2 := moveit2(moves)
	fmt.Println("Part 2:", part2.depth*part2.position, "(", part2, ")")

}

func moveit(moviments []move) location {
	posicio := location{position: 0, depth: 0}

	for _, moviment := range moviments {
		switch moviment.moviment {
		case "forward":
			posicio.position += moviment.units
		case "down":
			posicio.depth += moviment.units
		case "up":
			posicio.depth -= moviment.units
		}
	}

	return posicio
}

func moveit2(moviments []move) location {
	posicio := location{position: 0, depth: 0, aim: 0}

	for _, moviment := range moviments {
		fmt.Println("....", moviment)
		switch moviment.moviment {
		case "forward":
			posicio.position += moviment.units
			posicio.depth += posicio.aim * moviment.units
		case "down":
			// posicio.depth += moviment.units
			posicio.aim += moviment.units
		case "up":
			// posicio.depth -= moviment.units
			posicio.aim -= moviment.units
		}
		fmt.Println(posicio)
	}

	return posicio
}
