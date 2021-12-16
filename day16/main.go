package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Hex2Bits(code string) string {

	h2b := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	result := ""
	for _, digit := range code {
		result += h2b[string(digit)]
	}
	return result
}

func readLines(path string) (string, error) {
	file, err := os.Open(path)

	var line string

	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}
	return line, scanner.Err()
}

func main() {

	data, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result := Part1(data)
	fmt.Println("Part 1", result)

	result = Part2(data)
	fmt.Println("Part 2", result)
}

func ConvertToNumber(bits string) int {
	output, err := strconv.ParseInt(bits, 2, 64)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return int(output)
}

func sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func product(numbers []int) int {
	result := 1
	for _, v := range numbers {
		result *= v
	}
	return result
}

func minimum(numbers []int) int {

	lowest := numbers[0]
	for _, v := range numbers {
		if v < lowest {
			lowest = v
		}
	}
	return lowest
}

func maximum(numbers []int) int {

	greatest := numbers[0]
	for _, v := range numbers {
		if v > greatest {
			greatest = v
		}
	}
	return greatest
}

func greaterthan(numbers []int) int {

	if numbers[0] > numbers[1] {
		return 1
	}
	return 0
}

func lowerthan(numbers []int) int {

	if numbers[0] < numbers[1] {
		return 1
	}
	return 0
}

func equal(numbers []int) int {

	if numbers[0] == numbers[1] {
		return 1
	}
	return 0
}

func decode(num int, data string) ([]int, int, string) {

	isValid := strings.Replace(data, "0", "", -1)
	if len(isValid) == 0 {
		return []int{}, -1, ""
	}
	version := ConvertToNumber(data[0:3])
	versions := []int{
		version,
	}

	typeid := ConvertToNumber(data[3:6])

	undecoded := data
	result := 0
	numbers := []int{}

	switch typeid {
	case 4: // Literal
		number := ""
		start := 6
		for {
			number += data[start+1 : start+5]
			if data[start:start+1] == "0" {
				break
			}
			start += 5
		}
		undecoded = data[start+5:]

		result = ConvertToNumber(number)

	default: // Operator
		tipus := data[6:7]
		switch tipus {
		case "0":
			bits := ConvertToNumber(data[7 : 7+15])

			toDecode := data[22 : 22+bits]
			for len(toDecode) != 0 {
				version, number, data := decode(num+1, toDecode)
				if len(version) != 0 {
					numbers = append(numbers, number)

					versions = append(versions, version...)
				}
				toDecode = data
			}
			undecoded = data[22+bits:]

		case "1":
			packets := ConvertToNumber(data[7 : 7+11])
			undecoded = data[18:]

			for packets > 0 {
				version, number, data := decode(num+1, undecoded)
				if len(version) != 0 {
					numbers = append(numbers, number)
					versions = append(versions, version...)
				}
				undecoded = data
				packets--
			}
		}
		switch typeid {
		case 0: // suma
			result = sum(numbers)
		case 1: // product
			result = product(numbers)
		case 2: // minimum
			result = minimum(numbers)
		case 3: // maximum
			result = maximum(numbers)
		case 5: // greather than
			result = greaterthan(numbers)
		case 6: // less than
			result = lowerthan(numbers)
		case 7: // equal
			result = equal(numbers)
		}

	}

	return versions, result, undecoded
}

func Part1(data string) int {

	toDecode := Hex2Bits(data)

	versions, _, _ := decode(0, toDecode)

	return sum(versions)

}

const FILENAME = "input"

func Part2(data string) int {
	toDecode := Hex2Bits(data)

	_, values, _ := decode(0, toDecode)

	return values
}
