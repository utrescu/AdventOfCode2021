package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), ",")
		for _, x := range lines {
			n, _ := stringToInt(x)
			numbers = append(numbers, n)
		}

	}
	return numbers, scanner.Err()
}

func main() {
	crabs, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	resultat := Part1(crabs)
	fmt.Println("Part 1:", resultat)

	resultat2 := Part2(crabs)
	fmt.Println("Part 2:", resultat2)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}

func Part1(crabs []int) int {
	sort.Ints(crabs)
	median := len(crabs) / 2
	candidate := crabs[median]

	sum := 0
	for _, crab := range crabs {
		sum += abs(crab - candidate)
	}

	return sum
}

func Mitjana(crabs []int) int {
	suma := 0
	for _, s := range crabs {
		suma += s
	}
	return int(float64(suma) / float64(len(crabs)))
}

func calculateMovesCost(crabs []int, mitjana int) int {
	sum := 0
	for _, crab := range crabs {
		moves := abs(crab - mitjana)
		for i := 1; i <= moves; i++ {
			sum += i
		}
	}
	return sum
}

func Part2(crabs []int) int {
	mitjana := Mitjana(crabs)
	resultat := calculateMovesCost(crabs, mitjana)
	return resultat
}
