package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func stringArrayToInt(stringArray []string) ([]int, error) {
	var result []int
	for _, value := range stringArray {
		numero, err := stringToInt(value)
		if err != nil {
			return nil, err
		}
		result = append(result, numero)
	}
	return result, nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	argLength := len(os.Args[1:])
	if argLength != 1 {
		fmt.Println("Has de passar el nom del fitxer com a paràmetre")
		os.Exit(1)
	}

	filestrings, err := readLines(os.Args[1])
	if err != nil {
		panic("File read failed")
	}

	numbers, err := stringArrayToInt(filestrings)

	result1 := Increased1(numbers[0], numbers[1:])

	fmt.Println("Part 1:", result1)

	result2 := Increased2(numbers)
	fmt.Println("Part 2:", result2)

}

func Increased1(first int, rest []int) int {
	previous := first
	sum := 0

	for _, value := range rest {
		if value > previous {
			sum++
		}
		previous = value
	}
	return sum
}

func suma(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func Increased2(numbers []int) int {

	previous := suma(numbers[0:3])
	count := 0

	for first, _ := range numbers[1:] {
		if first+3 > len(numbers) {
			break
		}
		value := suma(numbers[first : first+3])
		if value > previous {
			count++
		}
		previous = value
	}
	return count
}
