package utils

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MAXINT = int(^uint(0) >> 1)
	MININT = -MAXINT - 1
)

func StringToInt(str string) int {
	nonFractionalPart := strings.Split(str, ".")
	v, err := strconv.Atoi(nonFractionalPart[0])
	if err != nil {
		panic(fmt.Sprintf("'%s' is not a number", nonFractionalPart[0]))
	}
	return v
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
