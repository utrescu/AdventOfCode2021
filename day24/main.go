package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

type Alu struct {
	variables map[string]int
	Result    string
}

func NewAlu(values []string) Alu {
	alu := Alu{}
	alu.variables = make(map[string]int)
	for _, value := range values {
		alu.variables[value] = 0
	}

	alu.Result = "00000000000000"
	return alu
}

func (a Alu) Clone() Alu {
	alu := Alu{}
	alu.variables = make(map[string]int)
	for k, v := range a.variables {
		alu.variables[k] = v
	}
	alu.Result = a.Result
	return alu
}

func (a Alu) String() string {
	return fmt.Sprintf("%d,%d,%d", a.variables["x"], a.variables["y"], a.variables["z"])
}

type Instruction struct {
	name           string
	value1, value2 string
	number         int
}

func readLines(path string) ([][]Instruction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	groups := make([][]Instruction, 0)
	scanner := bufio.NewScanner(file)
	group := make([]Instruction, 0)
	first := true
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if parts[0] == "inp" {

			if !first {
				groups = append(groups, group)

				group = make([]Instruction, 0)
			}
			first = false
			group = append(group, Instruction{parts[0], parts[1], "", 0})
		} else {
			secondArgument := parts[2]
			number, err := StringToInt(secondArgument)
			if err == nil {
				secondArgument = ""
			}

			group = append(group, Instruction{parts[0], parts[1], secondArgument, number})
		}
	}
	groups = append(groups, group)
	return groups, scanner.Err()
}

func process(alu Alu, commands []Instruction, imp int, round int) (Alu, error) {

	for _, command := range commands {
		secondArgument := command.number
		if command.value2 != "" {
			secondArgument = alu.variables[command.value2]
		}
		switch command.name {
		case "inp":
			alu.variables[command.value1] = imp
		case "add":
			alu.variables[command.value1] = alu.variables[command.value1] + secondArgument
		case "mul":
			alu.variables[command.value1] = alu.variables[command.value1] * secondArgument
		case "div":
			if secondArgument == 0 {
				return alu, errors.New("division by zero")
			}
			alu.variables[command.value1] = alu.variables[command.value1] / secondArgument
		case "mod":
			if alu.variables[command.value1] < 0 || secondArgument <= 0 {
				return alu, errors.New("incorrect Mod")
			}
			alu.variables[command.value1] = alu.variables[command.value1] % secondArgument
		case "eql":
			value := 0
			if alu.variables[command.value1] == secondArgument {
				value = 1
			}
			alu.variables[command.value1] = value
		default:
			panic("Instruction not recognised")
		}
	}

	alu.Result = alu.Result[0:round] + fmt.Sprintf("%d", imp) + alu.Result[round+1:]

	return alu, nil

}

func Part2(grups [][]Instruction) string {

	alu := NewAlu([]string{"x", "y", "z", "w"})

	toDo := make(map[string]Alu)
	toDo[alu.String()] = alu
	round := 0
	for _, grup := range grups {
		fmt.Println("round", round)
		newtoDo := make(map[string]Alu)
		for _, alu := range toDo {
			for i := 1; i < 10; i++ {
				oldAlu := alu.Clone()
				newalu, err := process(oldAlu, grup, i, round)
				if err == nil {
					if v, ok := newtoDo[newalu.String()]; ok {
						if v.Result > newalu.Result {
							newtoDo[newalu.String()] = newalu
						}
					} else {
						newtoDo[newalu.String()] = newalu
					}
				}
			}
		}
		toDo = newtoDo
		round++
	}

	var min string = "99999999999999"
	for _, v := range toDo {
		if v.variables["z"] == 0 {
			fmt.Println(v.String())
			if v.Result < min {
				min = v.Result
			}
		}
	}

	return min
}

func Part1(grups [][]Instruction) string {

	alu := NewAlu([]string{"x", "y", "z", "w"})

	toDo := make(map[string]Alu)
	toDo[alu.String()] = alu
	round := 0
	for _, grup := range grups {
		fmt.Println("round", round)
		newtoDo := make(map[string]Alu)
		for _, alu := range toDo {
			for i := 9; i > 0; i-- {
				oldAlu := alu.Clone()
				newalu, err := process(oldAlu, grup, i, round)
				if err == nil {
					if v, ok := newtoDo[newalu.String()]; ok {
						if v.Result < newalu.Result {
							newtoDo[newalu.String()] = newalu
						}
					} else {
						newtoDo[newalu.String()] = newalu
					}
				}
			}
		}
		toDo = newtoDo
		round++
	}

	var max string = ""
	for _, v := range toDo {
		if v.variables["z"] == 0 {
			fmt.Println(v.String())
			if v.Result > max {
				max = v.Result
			}
		}
	}

	return max
}

const FILENAME = "input"

func main() {

	instructs, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result1 := Part1(instructs)
	fmt.Println("Part 1:", result1)
	result2 := Part2(instructs)
	fmt.Println("Part 2:", result2)
}
