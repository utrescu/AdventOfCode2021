package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	maxInt   = int(^uint(0) >> 1)
	FILENAME = "input"
)

// --- UTILS

func MaxAndMin(numbers map[rune]int) (int, int) {
	maxNumber := 0
	minNumber := maxInt

	for _, n := range numbers {
		if n > maxNumber {
			maxNumber = n
		}
		if n < minNumber {
			minNumber = n
		}
	}
	return maxNumber, minNumber
}

func readLines(path string) (map[string]string, string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	pairs := make(map[string]string)
	start := ""

	hasStart := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		text := scanner.Text()
		if len(text) == 0 {
			hasStart = true
		} else {
			if !hasStart {
				start = text

			} else {
				data := strings.Split(text, " -> ")
				pairs[data[0]] = data[1]
			}
		}

	}
	return pairs, start, scanner.Err()
}

// ---------------------------------------------------

func main() {
	pairs, start, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result := Part2(start, pairs, 10)
	fmt.Println("Part 1:", result)

	result = Part2(start, pairs, 40)
	fmt.Println("Part 2:", result)

}

// --- Common

// Step: Avança la conversió
func Step(start string, pairs map[string]string) string {
	var newValue string

	for i := 0; i < len(start)-1; i += 1 {
		newValue += start[i : i+1]
		pair := start[i : i+2]
		if value, ok := pairs[pair]; ok {
			newValue += value
		}
	}
	newValue += string(start[len(start)-1])
	return newValue
}

func Step2(start map[string]int, pairs map[string]string) map[string]int {

	newStart := make(map[string]int)
	for pair, count := range start {
		if result, ok := pairs[pair]; ok {
			one := string(pair[0]) + result
			two := result + string(pair[1])
			if v, is := newStart[one]; is {
				newStart[one] = v + count
			} else {
				newStart[one] = count
			}
			if v, is := newStart[two]; is {
				newStart[two] = v + count
			} else {
				newStart[two] = count
			}
		}
	}
	return newStart

}

// -- Part 1 : He tornat a picar i ho he fet bruteforce, però la segona part no acaba mai ...

// func Part1(start string, pairs map[string]string, times int) int {

// 	for i := 0; i < times; i++ {
// 		start = Step(start, pairs)
// 	}

// 	counts := make(map[rune]int)
// 	for _, letter := range start {
// 		if value, ok := counts[letter]; ok {
// 			counts[letter] = value + 1
// 		} else {
// 			counts[letter] = 1
// 		}
// 	}

// 	max, min := MaxAndMin(counts)
// 	return max - min
// }

// Part2 Agafa tots els parells de la cadena i els va fent en paquets.
//   S'assembla molt a un d'anterior
func Part2(start string, pairs map[string]string, times int) int {
	startpairs := make(map[string]int)

	for i := 0; i < len(start)-1; i++ {
		pair := start[i : i+2]
		if value, ok := startpairs[pair]; ok {
			startpairs[pair] = value + 1
		} else {
			startpairs[pair] = 1
		}
	}

	for i := 0; i < times; i++ {
		startpairs = Step2(startpairs, pairs)
	}

	// Mira quines lletres estan més i menys vegades en el resultat
	// com que són parells només importa una de les dues però s'hi
	// ha de sumar 1 al darrer caràcter d'entrada
	last := rune(start[len(start)-1])
	max, min := MaxAndMinPair(startpairs, last)
	return max - min
}

// MaxAndMinPair suma els caràcters dels parells.
//    Com que estan superposats només agafo el primer.
//    Però li falta el caràcter final de la cadena original que
//    sempre estarà en segona posició
func MaxAndMinPair(pairs map[string]int, last rune) (int, int) {

	runesCount := make(map[rune]int)

	for pair, count := range pairs {
		single := rune(pair[0])
		if v, ok := runesCount[single]; ok {
			runesCount[single] = v + count
		} else {
			runesCount[single] = count
		}

	}

	runesCount[last] = runesCount[last] + 1

	return MaxAndMin(runesCount)
}
