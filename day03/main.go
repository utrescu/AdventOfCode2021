package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		caracters := strings.Split(scanner.Text(), "")
		lines = append(lines, caracters)
	}
	return lines, scanner.Err()
}

const oxygen = 0
const co2 = 1

func main() {
	linies, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	part1 := calculatePower(linies)
	fmt.Println("Part 1:", part1.epsilon*part1.gama)

	part2 := calculateLifeSuport(linies, oxygen) * calculateLifeSuport(linies, co2)
	fmt.Println("Part 2:", part2)
}

// ---- PART 1

type power struct {
	gama    int
	epsilon int
}

type counts [2]int

func calculatePower(linies [][]string) power {

	dataResults := make([]counts, len(linies[0]))

	for _, line := range linies {
		for index, column := range line {
			if column == "0" {
				dataResults[index][0]++
			} else {
				dataResults[index][1]++
			}
		}
	}

	result := power{gama: 0, epsilon: 0}
	mida := len(dataResults) - 1
	increment := 1
	for i := mida; i >= 0; i-- {
		if dataResults[i][0] > dataResults[i][1] {
			result.epsilon += increment
		} else {
			result.gama += increment
		}
		increment *= 2
	}
	return result

}

// -- PART 2

func array2Binary(data []string) int {
	result := 0
	increment := 1
	mida := len(data) - 1
	for i := mida; i >= 0; i-- {
		if data[i] == "1" {
			result += increment
		}
		increment *= 2
	}
	return result
}

// Retorna més corrent, menys corrent
func mostCommonFirst(data [][]string, column int) [2]string {
	result := [2]int{0, 0}
	for _, value := range data {
		if value[column] == "0" {
			result[0]++
		} else {
			result[1]++
		}
	}

	if result[0] > result[1] {
		return [2]string{"0", "1"}
	}
	return [2]string{"1", "0"}
}

func removeIncorrectNumbers(linies [][]string, column int, correctValue string) [][]string {
	nextValue := make([][]string, 0)

	for _, linia := range linies {
		if linia[column] == correctValue {
			nextValue = append(nextValue, linia)
		}
	}
	return nextValue
}

func calculateLifeSuport(linies [][]string, quin int) int {
	column := 0
	for len(linies) != 1 {
		correctValue := mostCommonFirst(linies, column)
		linies = removeIncorrectNumbers(linies, column, correctValue[quin])
		column = (column + 1) % len(linies[0])
	}
	return array2Binary(linies[0])
}
