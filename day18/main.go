package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func StringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func readLines(path string) ([]*SnailNumber, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := make([]*SnailNumber, 0, 1)
	for scanner.Scan() {
		number := &SnailNumber{}
		start := number
		line := scanner.Text()
		for _, symbol := range line {
			switch symbol {
			case '[':
				number.child1 = &SnailNumber{parent: number}
				number.child2 = &SnailNumber{parent: number}
				number = number.child1
			case ',':
				number = number.parent.child2
			case ']':
				number = number.parent
			default:
				parsed, err := strconv.Atoi(string(symbol))
				if err != nil {
					log.Fatalf("wrong input %s", line)
				}
				number.value = parsed
			}
		}
		numbers = append(numbers, start)
	}

	return numbers, scanner.Err()

}

const FILENAME = "input"

func main() {
	numbers, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result1 := Part1(numbers)
	fmt.Println("Part 1:", result1)
}

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

func NewSmallFishNumber(numbertext string, pc int) (*SnailNumber, int) {

	if numbertext[pc] == '[' {

		l, npc := NewSmallFishNumber(numbertext, pc+1)

		r, npc := NewSmallFishNumber(numbertext, npc+1)
		return &SnailNumber{child1: l, child2: r}, npc + 1

	}

	a, _ := StringToInt(string(numbertext[pc]))

	return &SnailNumber{value: a}, pc + 1

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

func (n *SnailNumber) reduce() {
	doing := true
	for doing {
		doing, _, _ = n.Explode(0, nil, nil)
		if !doing {
			doing = n.Split()
		}
	}
}

func (n SnailNumber) magnitude() int {
	if !n.hasChildren() {
		return n.value
	}
	return n.child1.magnitude()*3 + n.child2.magnitude()*2
}

func (n *SnailNumber) Print(level int) bool {
	if !n.hasChildren() {
		fmt.Print(n.value)
		return true
	}
	fmt.Print("[")
	n.child1.Print(level + 1)
	fmt.Print(",")
	n.child2.Print(level + 1)
	fmt.Print("]")
	return true
}

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
