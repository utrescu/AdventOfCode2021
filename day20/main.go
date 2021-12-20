package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, [][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	enhace := strings.Split(scanner.Text(), "")
	scanner.Scan() // blank line
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		lines = append(lines, line)
	}
	return enhace, lines, scanner.Err()
}

type Enhace []string

func (e Enhace) Get(i int) string {
	if i > 512 {
		panic("Maximum enhace string reached")
	}
	return string(e[i])
}

type Image [][]string

func (i Image) AddBorder(size int) Image {
	newImage := Image{}
	oldSize := len(i)

	for row := -size; row <= oldSize+size; row++ {
		newRow := make([]string, 0)
		for col := -size; col <= oldSize+size; col++ {

			if col < 0 || row < 0 || col >= oldSize || row >= oldSize {
				newRow = append(newRow, ".")
			} else {
				newRow = append(newRow, i[row][col])
			}
		}
		newImage = append(newImage, newRow)
	}
	return newImage
}

func (i Image) Get(y, x int) int {
	bits := ""
	for row := y - 1; row <= y+1; row++ {
		for col := x - 1; col <= x+1; col++ {
			data := ""
			if col < 0 || row < 0 || col >= len(i[0]) || row >= len(i) {
				data = i[0][0]
			} else {
				data = i[row][col]
			}
			switch data {
			case "#":
				bits += "1"
			case ".":
				bits += "0"
			}
		}
	}
	output, _ := strconv.ParseInt(bits, 2, 64)
	return int(output)
}

func (i Image) Print() {
	for row := range i {
		for _, v := range i[row] {
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func (i Image) Count() int {
	count := 0
	for row := 0; row < len(i)-1; row++ {
		for col := 0; col < len(i)-1; col++ {
			if i[row][col] == "#" {
				count++
			}
		}
	}
	return count
}

func Part1(image Image, enhace Enhace, times int) int {

	// He anat afegint marges fins que el valor es repetia ... XAPUSSA
	image2 := image.AddBorder(100)
	border := 1

	for i := 0; i < times; i++ {
		var newImage Image
		for row := border; row <= len(image2)-border; row++ {
			newRow := make([]string, 0)
			for col := border; col <= len(image2[0])-border; col++ {
				value := image2.Get(row, col)
				newRow = append(newRow, enhace.Get(value))
			}
			newImage = append(newImage, newRow)
		}
		image2 = newImage

	}

	return image2.Count()
}

const FILENAME = "input"

func main() {
	enhace, lines, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result1 := Part1(lines, enhace, 2)
	fmt.Println("Part 1:", result1)
	result2 := Part1(lines, enhace, 50)
	fmt.Println("Part 2:", result2)
}
