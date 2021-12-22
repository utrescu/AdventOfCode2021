package main

import (
	"bufio"
	"day22/utils"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func readLines(path string) ([]Area, error) {
	var re = regexp.MustCompile(`(?m)(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Area
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		enabled := true
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		if match[1] == "off" {
			enabled = false
		}
		data := Area{
			utils.StringToInt(match[2]),
			utils.StringToInt(match[3]),
			utils.StringToInt(match[4]),
			utils.StringToInt(match[5]),
			utils.StringToInt(match[6]),
			utils.StringToInt(match[7]),
			enabled,
		}
		lines = append(lines, data)

	}
	return lines, scanner.Err()
}

type Area struct {
	x1, x2  int
	y1, y2  int
	z1, z2  int
	enabled bool
}

func (c Area) Intersect(other Area) (Area, error) {
	if other.x1 > c.x2 || c.x1 > other.x2 {
		return Area{}, errors.New("No")
	}

	if other.y1 > c.y2 || c.y1 > other.y2 {
		return Area{}, errors.New("No")
	}

	if other.z1 > c.z2 || c.z1 > other.z2 {
		return Area{}, errors.New("No")
	}

	xMin := utils.Max(c.x1, other.x1)
	xMax := utils.Min(c.x2, other.x2)
	yMin := utils.Max(c.y1, other.y1)
	yMax := utils.Min(c.y2, other.y2)
	zMin := utils.Max(c.z1, other.z1)
	zMax := utils.Min(c.z2, other.z2)

	// Enabled? Si xoca no ho ha d'estar

	enabled := !c.enabled
	// Si són del mateix tipus la intersecció no s'ha de comptar
	if c.enabled == other.enabled {
		enabled = !other.enabled
	}

	return Area{x1: xMin, x2: xMax, y1: yMin, y2: yMax, z1: zMin, z2: zMax, enabled: enabled}, nil
}

func (c Area) Size() uint {
	return uint((utils.Abs(c.x1-c.x2) + 1) * (utils.Abs(c.y1-c.y2) + 1) * (utils.Abs(c.z1-c.z2) + 1))
}

func (c Area) Inside(low int, upper int) bool {
	if c.x1 < low || c.y1 < low || c.z1 < low {
		return false
	}
	if c.x2 > upper || c.y2 > upper || c.z2 > upper {
		return false
	}

	return true
}

func (c Area) String() string {
	retorn := ""
	for x := c.x1; x <= c.x2; x++ {
		for y := c.y1; y <= c.y2; y++ {
			for z := c.z1; z <= c.z2; z++ {
				retorn += fmt.Sprintf("%d, %d, %d, %t\n", x, y, z, c.enabled)
			}
		}
	}
	return retorn
}

// -- Part 1

// Part 1 i 2: Es van afegint els cubs un a un i en comptes de treure afegeixo
//             les àrees d'interseccions com a resta (excepte amb els que ja eren
// 				interseccions on s'ha de tornar a sumar)
func Part1(data []Area, low, up int) uint64 {

	cubesAdded := make([]Area, 0)

	for _, cube := range data {
		if !cube.Inside(low, up) {
			continue
		}

		cubesProcessed := make([]Area, 0)
		for _, oldCube := range cubesAdded {
			// Poso les interseccions a restar (o sumar si eren desactivacions)
			if nou, err := oldCube.Intersect(cube); err == nil {
				cubesProcessed = append(cubesProcessed, nou)

				// fmt.Println(nou.String())
			}

		}
		// Els cubs que desactiven no cal afegir-los
		if cube.enabled {
			cubesProcessed = append(cubesProcessed, cube)
			// fmt.Println(cube.String())
		}
		cubesAdded = append(cubesAdded, cubesProcessed...)
	}

	var sum uint64 = 0
	for _, v := range cubesAdded {
		size := v.Size()
		if v.enabled {
			sum += uint64(size)
		} else {
			sum -= uint64(size)
		}
	}
	return sum
}

const FILENAME = "input"

func main() {

	data, err := readLines(FILENAME)
	if err != nil {
		panic("File read failed")
	}

	result1 := Part1(data, -50, 50)
	fmt.Println("Part 1:", result1)

	result2 := Part1(data, utils.MININT, utils.MAXINT)
	fmt.Println("Part 2:", result2)

}
