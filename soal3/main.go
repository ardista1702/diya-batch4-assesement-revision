package main

import (
	"fmt"
	"strings"
)

type coordinate struct {
	X int
	Y int
}

func isIntersection(coordiate1, coordinate2 coordinate) bool {
	return coordiate1.X == coordinate2.X && coordiate1.Y == coordinate2.Y
}

func move(direction string, c coordinate) coordinate {
	switch direction {
	case "U":
		c.Y++
	case "S":
		c.Y--
	case "B":
		c.X--
	case "T":
		c.X++
	}

	return c
}

func countTotalIntersection(testcase string) int {
	totalIntersection := 1
	splittedTc := strings.Split(testcase, " ")
	path1 := splittedTc[0]
	path2 := splittedTc[1]
	p1 := 0
	p2 := 0
	c1 := coordinate{0, 0}
	c2 := coordinate{0, 0}
	for p1 < len(path1) && p2 < len(path2) {
		c1 = move(string(path1[p1]), c1)
		c2 = move(string(path2[p2]), c2)
		if isIntersection(c1, c2) {
			totalIntersection++
		}
		p1++
		p2++
	}

	return totalIntersection
}

func main() {
	testCases := []string{
		"UUU TTT",
		"BBUU BBUU",
		"TTSSSSS SSTTTTT",
	}
	for _, tc := range testCases {
		fmt.Println(countTotalIntersection(tc))
	}
}
