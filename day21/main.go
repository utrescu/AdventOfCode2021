package main

import (
	"bufio"
	"day21/utils"
	"fmt"
	"os"
	"regexp"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]Player, error) {
	var re = regexp.MustCompile(`Player (\d+) starting position: (\d+)`)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Player
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		position, _ := utils.StringToInt(match[2])
		lines = append(lines, Player{position: position})
	}
	return lines, scanner.Err()
}

// -- Part 1

type Player struct {
	position int
	score    int
}

func (p *Player) Move(n int) {
	p.position = (p.position+n-1)%10 + 1
	p.score += p.position
}

func (s Player) Stop(end int) bool {
	return s.score >= end
}

func Part1(players []Player, rollsides, endscore int) (int, int) {
	dice := 1
	rolls := 0
	stop := false
	winner := 0
	for !stop {
		for i := range players {
			rolls += 3
			players[i].Move(3*dice + 3)
			if players[i].Stop(endscore) {
				stop = true
				winner = i
				break
			}
			dice = (dice + 3) % rollsides
		}

	}

	return utils.Min(players[0].score, players[1].score) * rolls, winner
}

// -- PART 2

type Position struct {
	players   []Player
	turn      int
	universes int
}

func (p Position) NewPosition(roll int) Position {

	newPlayers := make([]Player, len(p.players))
	for i, v := range p.players {
		newPlayers[i] = v
	}
	newPlayers[p.turn].Move(roll)
	turn := (p.turn + 1) % len(p.players)

	universes := p.universes * UniversesPerRoll(roll)

	return Position{players: newPlayers, turn: turn, universes: universes}
}

func (p Position) HaveWinner(score int) int {
	for i := range p.players {
		if p.players[i].Stop(score) {
			return i
		}
	}
	return -1
}

func UniversesPerRoll(roll int) int {
	switch roll {
	case 3:
		return 1 // 1,1,1
	case 4:
		return 3 // 2,1,1 1,1,2 1,2,1
	case 5:
		return 6 // 2,2,1 1,2,2 2,1,2 3,1,1 1,1,3 1,3,1
	case 6:
		return 7 // 2,2,2 3,2,1 3,1,2 2,3,1 1,3,2 1,2,3 2,1,3
	case 7:
		return 6 // 3,3,1 1,3,3 3,1,3 3,2,2 2,3,2 2,2,3
	case 8:
		return 3 // 3,3,2 2,3,3 3,2,3
	case 9:
		return 1 // 3,3,3
	}
	panic("not possible result!")
}

func Part2(players []Player, endscore int) int {

	rolls := []int{3, 4, 5, 6, 7, 8, 9}
	winners := make([]int, len(players))

	start := Position{
		players:   players,
		turn:      0,
		universes: 1,
	}

	positions := []Position{start}

	for len(positions) > 0 {
		current := positions[0]
		positions = positions[1:]
		for _, roll := range rolls {

			newPosition := current.NewPosition(roll)
			if winner := newPosition.HaveWinner(endscore); winner != -1 {
				winners[winner] += newPosition.universes
			} else {
				positions = append(positions, newPosition)
			}
		}
	}
	return utils.Max(winners[0], winners[1])
}

const FILENAME = "input"

func main() {

	players, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result1, _ := Part1(players, 100, 1000)
	fmt.Println("Part 1:", result1)

	players, err = readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}
	result2 := Part2(players, 21)
	fmt.Println("Part 2:", result2)

}
