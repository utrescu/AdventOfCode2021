package main

import (
	"container/heap"
	"day23/mapa"
	"fmt"
)

func Part1(state mapa.Mapa, solution [][]string) int {

	// toProcess := []mapa.Mapa{state}
	positionsDone := make(map[string]bool)

	toProcess := make(mapa.PriorityQueue, 1)
	toProcess[0] = &mapa.Item{
		Value:    state,
		Priority: 0,
		Index:    len(toProcess),
	}
	heap.Init(&toProcess)

	for len(toProcess) > 0 {

		// current := toProcess[0]
		// toProcess = toProces[1:]
		item := heap.Pop(&toProcess).(*mapa.Item)
		current := item.Value

		if current.RomsEquals(solution) {
			return current.Cost
		}

		// Si ja està fet, descarta
		if _, ok := positionsDone[current.String()]; ok {
			continue
		}

		positionsDone[current.String()] = true

		// calcula nous moviments i afegeix-los a ToProcess
		// 1. Moviment des de les rooms
		newMoves := current.ToHalfMoves()
		// 2. Return home
		newMoves = append(newMoves, current.ToHomeMoves()...)
		for _, move := range newMoves {
			index := len(toProcess)
			if _, ok := positionsDone[move.String()]; !ok {
				item := &mapa.Item{
					Value:    move,
					Priority: move.Cost,
					Index:    index,
				}
				toProcess.Push(item)
				index++
			}
		}

		heap.Init(&toProcess)

	}

	return -1
}

func main() {

	solution1 := [][]string{
		{"A", "A"},
		{"B", "B"},
		{"C", "C"},
		{"D", "D"},
	}

	rooms := [][]string{
		{"B", "C"}, // As room
		{"C", "D"}, // Bs room
		{"A", "D"}, // Cs room
		{"B", "A"}, // Ds room
	}

	half := []string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."}

	result1 := Part1(mapa.NewMap(half, rooms), solution1)
	fmt.Println("Part 1:", result1)

	solution2 := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "B", "B"},
		{"C", "C", "C", "C"},
		{"D", "D", "D", "D"},
	}

	rooms2 := [][]string{
		{"B", "D", "D", "C"}, // As room
		{"C", "C", "B", "D"}, // Bs room
		{"A", "B", "A", "D"}, // Cs room
		{"B", "A", "C", "A"}, // Ds room
	}

	result2 := Part1(mapa.NewMap(half, rooms2), solution2)
	fmt.Println("Part 2:", result2)
}
