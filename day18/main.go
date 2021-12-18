package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// --- UTILS

func StringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

// Parse string (he suat)
func readLines(path string) ([]*SnailNumber, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := make([]*SnailNumber, 0, 1)
	for scanner.Scan() {
		nextNumber := &SnailNumber{}
		root := nextNumber
		line := scanner.Text()
		for _, character := range line {
			switch character {
			case '[':
				nextNumber.child1 = &SnailNumber{parent: nextNumber}
				nextNumber.child2 = &SnailNumber{parent: nextNumber}
				// child1 is ..
				nextNumber = nextNumber.child1
			case ']':
				nextNumber = nextNumber.parent
			case ',':
				// child2 is ..
				nextNumber = nextNumber.parent.child2
			default:
				// només hi ha números d'una xifra
				numero, _ := StringToInt(string(character))
				nextNumber.value = numero
			}
		}
		numbers = append(numbers, root)
	}

	return numbers, scanner.Err()

}

// -- Snail Numbers
type SnailNumber struct {
	value  int
	child1 *SnailNumber
	child2 *SnailNumber
	parent *SnailNumber
}

func (n *SnailNumber) hasChildren() bool {
	return n.child1 != nil && n.child2 != nil
}

func (n *SnailNumber) isLeaf() bool {
	return !n.hasChildren()
}

func (n *SnailNumber) Explode(level int, anterior *SnailNumber, toNext *int) (bool, *SnailNumber, *int) {
	done := false

	if n.isLeaf() && toNext != nil {
		n.value += *toNext
		toNext = nil
		return true, n, nil
	}

	if n.isLeaf() {
		return false, n, toNext
	}

	if level >= 4 && n.hasChildren() && n.child1.isLeaf() && n.child2.isLeaf() && toNext == nil {
		nextValue := &n.child2.value

		if anterior != nil {
			anterior.value += n.child1.value
		}
		n.value = 0
		n.child1 = nil
		n.child2 = nil

		return false, anterior, nextValue
	} else {

		done, anterior, toNext = n.child1.Explode(level+1, anterior, toNext)
		if !done {
			done, anterior, toNext = n.child2.Explode(level+1, anterior, toNext)
		}

	}

	return done, anterior, toNext
}

func (n *SnailNumber) Split() bool {
	done := false
	if n.isLeaf() {
		if n.value >= 10 {
			first := int(math.Floor(float64(n.value) / 2))
			second := int(math.Ceil(float64(n.value) / 2))

			n.child1 = &SnailNumber{value: first, parent: n}
			n.child2 = &SnailNumber{value: second, parent: n}
			n.value = 0
			return true
		}
	} else {
		done = n.child1.Split()
		if !done {
			done = n.child2.Split()
		}
	}
	return done
}

// reduce redueix el número pas a pas, primer explodes i si no van després fa splits
func (n *SnailNumber) reduce() {
	doing := true
	for doing {
		doing, _, _ = n.Explode(0, nil, nil)
		if !doing {
			doing = n.Split()
		}
	}
}

// magnitude calcula la magnitud del número
func (n SnailNumber) magnitude() int {
	if !n.hasChildren() {
		return n.value
	}
	return n.child1.magnitude()*3 + n.child2.magnitude()*2
}

// Print no serveix de res però pinta l'expressió
func (n *SnailNumber) Print() bool {
	if !n.hasChildren() {
		fmt.Print(n.value)
		return true
	}
	fmt.Print("[")
	n.child1.Print()
	fmt.Print(",")
	n.child2.Print()
	fmt.Print("]")
	return true
}

// --- PART 1

func Part1(numbers []*SnailNumber) int {

	// Primer número
	total := numbers[0]
	total.reduce()

	for _, number := range numbers[1:] {

		newTotal := &SnailNumber{child1: total, child2: number}
		number.parent = newTotal
		total.parent = newTotal
		newTotal.reduce()
		total = newTotal
	}

	return total.magnitude()
}

// -- Part 2

func (s *SnailNumber) Clone() *SnailNumber {
	copia := &SnailNumber{}
	if s.child1 != nil {
		copia.child1 = s.child1.Clone()
		copia.child1.parent = copia
	}
	if s.child2 != nil {
		copia.child2 = s.child2.Clone()
		copia.child2.parent = copia
	}
	copia.value = s.value

	return copia
}

func Part2(numbers []*SnailNumber) int {
	max := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			firstNumber := numbers[i].Clone()
			secondNumber := numbers[j].Clone()

			suma := &SnailNumber{child1: firstNumber, child2: secondNumber}
			firstNumber.parent = suma
			secondNumber.parent = suma
			suma.reduce()
			value := suma.magnitude()
			if value > max {
				max = value
			}
		}
	}
	return max
}

const FILENAME = "input"

func main() {
	numbers, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result1 := Part1(numbers)
	fmt.Println("Part 1:", result1)

	numbers, err = readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result2 := Part2(numbers)
	fmt.Println("Part 2: ", result2)
}
