package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const FILENAME = "input"

// readLines carrega els camins del fitxer
func readLines(path string) (map[string][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "-")
		origin := data[0]
		destination := data[1]

		if value, ok := lines[origin]; ok {
			value = append(value, destination)
			lines[origin] = value
		} else {
			lines[origin] = []string{destination}
		}

		if value, ok := lines[destination]; ok {
			value = append(value, origin)
			lines[destination] = value
		} else {
			lines[destination] = []string{origin}
		}
	}
	return lines, scanner.Err()
}

func main() {
	filestrings, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result1 := Part(1, filestrings)
	fmt.Println("Part 1:", result1)

	result2 := Part(2, filestrings)
	fmt.Println("Part 2:", result2)
}

// -- Part 1 i 2

func Part(part int, data map[string][]string) int {
	paths := GeneratePath([]string{"start"}, "start", part, data)
	return len(paths)
}

func IsLower(s string) bool {
	return unicode.IsLower(rune(s[0]))
}

// alreadyInserted mira si el valor en minúscules pot ser solució
//    per la part 1 en tenia prou amb mirar si hi era però per la
//    part 2 he hagut de canviar-lo
func alreadyInserted(list []string, value string, times int) bool {
	index := make(map[string]int)
	alreadyOneRepeated := 0

	for _, v := range list {
		if IsLower(v) {
			if count, ok := index[v]; ok {
				if count+1 >= times {
					alreadyOneRepeated = -1
				}
				index[v] = count + 1
			} else {
				index[v] = 1
			}
		}
	}

	count, ok := index[value]
	if count < times+alreadyOneRepeated {
		return false
	}
	return ok
}

// GeneratePath localitza tots els resultats
//    current: resultat que estem calculant
//    actual: valor actual (es podria obtenir de current[len(current)-1] però em fa mandra)
//    smalltimes: quantes vegades pot sortir un resultat minúscula (és per la diferència de la part 2)
//    paths: camins possibles
func GeneratePath(current []string, actual string, smalltimes int, paths map[string][]string) [][]string {
	if actual == "end" {
		return [][]string{current}
	}

	newResults := make([][]string, 0)
	for _, item := range paths[actual] {

		if item == "start" {
			continue
		}
		if IsLower(item) {
			if alreadyInserted(current, item, smalltimes) {
				continue
			}
		}
		newCurrent := append(current, item)
		data := GeneratePath(newCurrent, item, smalltimes, paths)
		if len(data) != 0 {
			newResults = append(newResults, data...)
		}
	}
	return newResults
}
