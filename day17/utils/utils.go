package utils

import (
	"strconv"
	"strings"
)

const (
	MAXINT = int(^uint(0) >> 1)
	MININT = -MAXINT - 1
)

// --- UTILS

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func StringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}
