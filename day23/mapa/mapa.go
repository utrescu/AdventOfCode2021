package mapa

import (
	"day23/utils"
	"fmt"
)

func RoomHome(letter string) (int, int) {

	// ...........
	//   A B C D

	switch letter {
	case "A":
		return 2, 0
	case "B":
		return 4, 1
	case "C":
		return 6, 2
	case "D":
		return 8, 3
	default:
		panic("Ops")
	}

}

// TODO: Generalitzar
func HalfwayPositions() []int {
	return []int{0, 1, 3, 5, 7, 9, 10}
}

func CostToMove(letter string) int {

	// ...........
	//   A B C D

	switch letter {
	case "A":
		return 1
	case "B":
		return 10
	case "C":
		return 100
	case "D":
		return 1000
	default:
		return 0
	}

}

type Mapa struct {
	Half  []string
	Rooms [][]string
	Cost  int
}

func NewMap(h []string, r [][]string) Mapa {
	return Mapa{h, r, 0}
}

func (m Mapa) ToHalfMoves() []Mapa {
	maps := make([]Mapa, 0)

	for roomNumber, currentRoom := range m.Rooms {
		if !isEmpty(currentRoom) {
			amphi, movesToExit, rooms := m.exitFromRoom(roomNumber)
			roomPosition := roomNumber*2 + 2
			for _, possible := range m.movesToHalf(roomPosition) {
				halfResult := make([]string, len(m.Half))
				copy(halfResult, m.Half)
				halfResult[possible] = amphi

				costHalf := utils.Abs(possible - roomPosition)

				cost := (movesToExit + costHalf) * CostToMove(amphi)

				maps = append(maps, Mapa{halfResult, rooms, cost + m.Cost})
			}
		}
	}
	return maps
}

func (m Mapa) ToHomeMoves() []Mapa {
	maps := make([]Mapa, 0)

	for _, possible := range m.movesToHome() {
		halfResult := make([]string, len(m.Half))
		copy(halfResult, m.Half)
		amphi := m.Half[possible]
		halfResult[possible] = "."
		home, _ := RoomHome(amphi)
		costHalf := utils.Abs(possible - home)

		newRooms, movesToEnter := m.enterToRoom(amphi)
		cost := (movesToEnter + costHalf) * CostToMove(amphi)

		newMapa := Mapa{halfResult, newRooms, cost + m.Cost}

		maps = append(maps, newMapa)
	}
	return maps
}

func (m Mapa) RomsEquals(solution [][]string) bool {

	for i, rooms := range m.Rooms {
		for j, amphi := range rooms {
			if amphi != solution[i][j] {
				return false
			}
		}
	}
	return true
}

func isEmpty(data []string) bool {
	for _, v := range data {
		if v != "." {
			return false
		}
	}
	return true
}

func (m Mapa) canGoToHalf(from, to int) bool {
	inici := utils.Min(from, to)
	final := utils.Max(from, to)

	for i := inici; i <= final; i++ {
		if m.Half[i] != "." {
			return false
		}
	}
	return true
}

func (m Mapa) homeIsLocked(amphi string) bool {
	_, room := RoomHome(amphi)
	for _, value := range m.Rooms[room] {
		if value != "." && value != amphi {
			return true
		}
	}
	return false
}

func (m Mapa) canGoHome(from int) bool {
	amphi := m.Half[from]
	if m.homeIsLocked(amphi) {
		return false
	}
	// skip current position
	step := 1

	to, _ := RoomHome(amphi)
	if to < from {
		step = -1
	}
	return m.canGoToHalf(from+step, to)
}

func (m Mapa) movesToHalf(origin int) []int {

	possibles := make([]int, 0)
	for _, position := range HalfwayPositions() {
		if m.Half[position] == "." {
			// Hi puc arribar?
			if m.canGoToHalf(origin, position) {
				possibles = append(possibles, position)
			}
		}
	}
	return possibles
}

func (m Mapa) movesToHome() []int {

	possibles := make([]int, 0)
	for _, position := range HalfwayPositions() {
		if m.Half[position] != "." {
			// Hi puc arribar?
			if m.canGoHome(position) {
				possibles = append(possibles, position)
			}
		}
	}

	return possibles
}

func (m Mapa) exitFromRoom(num int) (string, int, [][]string) {

	result := ""
	moves := 99

	resultRooms := make([][]string, 0)
	for roomNumber, room := range m.Rooms {
		newRoom := make([]string, len(room))
		copy(newRoom, room)
		if roomNumber == num {
			for i, v := range room {
				if v != "." {
					result = v
					moves = i + 1
					newRoom[i] = "."
					break
				}
			}
		}

		resultRooms = append(resultRooms, newRoom)
	}

	return result, moves, resultRooms
}

func (m Mapa) enterToRoom(amphi string) ([][]string, int) {
	resultRooms := make([][]string, 0)
	_, num := RoomHome(amphi)
	moves := 0

	for roomNumber, room := range m.Rooms {
		newRoom := make([]string, len(room))

		copy(newRoom, room)
		if roomNumber == num {

			for i := len(room) - 1; i >= 0; i-- {
				if room[i] == "." {
					newRoom[i] = amphi
					moves = i + 1
					break
				}
			}
		}

		resultRooms = append(resultRooms, newRoom)
	}

	return resultRooms, moves
}

func (m Mapa) String() string {
	return fmt.Sprintf("%v %v", m.Half, m.Rooms)
}

// -- Priority queue
type Item struct {
	Value    Mapa
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Retornar la prioritat més baixa
	return pq[i].Priority < pq[j].Priority
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
