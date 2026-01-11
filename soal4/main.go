package main

import (
	"fmt"
	"strconv"
)

func countTotalClockWise(current, target int) int {
	return (target - current + 10) % 10
}

func countTotalCounterClockWise(current, target int) int {
	return (current - target + 10) % 10
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countTotalDirectionChange(tc string) int {
	patterns := make([]int, 0)
	for i := 0; i < len(tc); i++ {
		num, _ := strconv.Atoi(string(tc[i]))
		patterns = append(patterns, num)
	}
	changes := 0
	current := 0
	lastDir := "left"

	fmt.Println(patterns)
	for i := 0; i < len(patterns); i++ {
		target := patterns[i]
		var moveDir string
		cw := (target - current + 10) % 10
		ccw := (current - target + 10) % 10
		if cw < ccw {
			moveDir = "right"
		} else if ccw < cw {
			moveDir = "left"
		} else {
			moveDir = lastDir
		}

		if moveDir != lastDir && i > 0 {
			changes++
		}
		fmt.Printf("current: %d, target: %d,clockwise: %d, counterClockwise: %d, moveDir: %s, lastDir: %s, changes: %d\n", current, target, cw, ccw, moveDir, lastDir, changes)
		lastDir = moveDir
		current = target
	}

	return changes
}

func main() {
	testCases := []string{
		// "12345",
		// "2121",
		"981",
		// "4350",
	}
	for _, tc := range testCases {
		fmt.Println(countTotalDirectionChange(tc))
	}
}
