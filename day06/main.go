package main

import (
	"fmt"
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

func main() {
	//value := "3,4,3,1,2"
	value := "1,4,1,1,1,1,1,1,1,4,3,1,1,3,5,1,5,3,2,1,1,2,3,1,1,5,3,1,5,1,1,2,1,2,1,1,3,1,5,1,1,1,3,1,1,1,1,1,1,4,5,3,1,1,1,1,1,1,2,1,1,1,1,4,4,4,1,1,1,1,5,1,2,4,1,1,4,1,2,1,1,1,2,1,5,1,1,1,3,4,1,1,1,3,2,1,1,1,4,1,1,1,5,1,1,4,1,1,2,1,4,1,1,1,3,1,1,1,1,1,3,1,3,1,1,2,1,4,1,1,1,1,3,1,1,1,1,1,1,2,1,3,1,1,1,1,4,1,1,1,1,1,1,1,1,1,1,1,1,2,1,1,5,1,1,1,2,2,1,1,3,5,1,1,1,1,3,1,3,3,1,1,1,1,3,5,2,1,1,1,1,5,1,1,1,1,1,1,1,2,1,2,1,1,1,2,1,1,1,1,1,2,1,1,1,1,1,5,1,4,3,3,1,3,4,1,1,1,1,1,1,1,1,1,1,4,3,5,1,1,1,1,1,1,1,1,1,1,1,1,1,5,2,1,4,1,1,1,1,1,1,1,1,1,1,1,1,1,5,1,1,1,1,1,1,1,1,2,1,4,4,1,1,1,1,1,1,1,5,1,1,2,5,1,1,4,1,3,1,1"

	llista, _ := stringArrayToInt(strings.Split(value, ","))

	result := Part2(llista, 80)
	fmt.Println("Part 1:", result)

	result2 := Part2(llista, 256)
	fmt.Println("Part 2:", result2)
}

// func Part1(fish []int, days int) int {
// 	for i := 0; i < days; i++ {
// 		var children []int
// 		for position, value := range fish {
// 			if value == 0 {
// 				fish[position] = 6
// 				children = append(children, 8)
// 				// new fish
// 			} else {
// 				value = value - 1
// 				fish[position] = value
// 			}

// 		}
// 		fish = append(fish, children...)
// 	}
// 	return len(fish)
// }

func Part2(fish []int, days int) int {
	groupedFish := make(map[int]int)

	for _, days := range fish {
		qty, exists := groupedFish[days]
		if exists {
			groupedFish[days] = qty + 1
		} else {
			groupedFish[days] = 1
		}
	}

	for i := 0; i < days; i++ {
		reagroup := make(map[int]int)
		for value, qty := range groupedFish {

			if value == 0 {
				// criar
				reagroup[8] = qty
				old, exists := reagroup[6]
				if exists {
					reagroup[6] = old + qty
				} else {
					reagroup[6] = qty
				}
			} else {
				old, exists := reagroup[value-1]
				if exists {
					reagroup[value-1] = old + qty
				} else {
					reagroup[value-1] = qty
				}
			}
		}
		groupedFish = reagroup
	}

	sum := 0
	for _, v := range groupedFish {
		sum += v
	}

	return sum
}
