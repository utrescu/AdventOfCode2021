package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// --- Implemento una Pila per no fer servir cap llibreria extra ----
type Stack struct {
	values []*string
	count  int
}

func (s *Stack) Push(n *string) {
	s.values = append(s.values[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *string {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.values[s.count]
}

func NewStack() *Stack {
	return &Stack{}
}

// --- Utils
func readLines(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}
	return lines, scanner.Err()
}

// --- main

func main() {

	chunks, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	resultat1 := Part1(chunks)
	fmt.Println("Part 1: ", resultat1)

	chunks2 := make([][]string, 0)
	for _, chunk := range chunks {
		if ProcessChunk(chunk) == 0 {
			chunks2 = append(chunks2, chunk)
		}
	}

	resultat2 := Part2(chunks2)
	fmt.Println("Part 2: ", resultat2)
}

// getCompleteSimbols retorna el simbol que tanca
func getCompleteSimbols() map[string]string {
	return map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">"}
}

// --- Part 1

// Punts que s'obtenen segons el símbol erròni
func getPart1ValueScores() map[string]int {
	return map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
}

func Part1(chunks [][]string) int {
	totalScore := 0

	for _, chunk := range chunks {
		totalScore += ProcessChunk(chunk)
	}

	return totalScore
}

// ProcessChunk retorna els punts obtinguts per detectar l'error
//   els punts depenen del simbol que dóna error
func ProcessChunk(chunk []string) int {
	score := 0
	scores := getPart1ValueScores()
	simbols := getCompleteSimbols()
	stack := NewStack()
	for _, value := range chunk {
		if _, ok := simbols[value]; ok {
			data := value
			stack.Push(&data)
		} else {
			v := stack.Pop()
			if simbols[*v] != value {
				// Error!
				score = scores[value]
				break
			}
		}
	}
	return score
}

// --- Part 2

func Part2(chunks [][]string) int {
	scores := make([]int, 0)
	for _, chunk := range chunks {
		complete := CompleteChunk(chunk)
		scores = append(scores, GetCompleteScore(complete))
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

// GetCompleteScore calcula els punts obtinguts per tancar l'expressió
func GetCompleteScore(complete []string) int {
	valueScores := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	score := 0

	for _, v := range complete {
		score *= 5
		score += valueScores[v]
	}
	return score
}

// CompleteChunk retorna els símbols que falten per tancar l'expressió
func CompleteChunk(chunk []string) []string {
	simbols := getCompleteSimbols()
	stack := NewStack()
	for _, value := range chunk {
		if _, ok := simbols[value]; ok {
			data := value
			stack.Push(&data)
		} else {
			stack.Pop()
		}
	}
	complete := make([]string, 0)
	for stack.count != 0 {
		v := stack.Pop()
		complete = append(complete, simbols[*v])
	}
	return complete
}
