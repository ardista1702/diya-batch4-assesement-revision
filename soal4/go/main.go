package main

import (
	"fmt"
	"strconv"
)

func countTotalClockWise(current, target int) int {
	total := 0
	for current != target {
		current--
		if current < 0 {
			current = 9
		}
		total++
	}
	return total
}

func countTotalCounterClockWise(current, target int) int {
	total := 0
	for current != target {
		current++
		if current > 9 {
			current = 0
		}
		total++
	}
	return total
}

func countTotalDirectionChange(tc string) int {
	patterns := make([]int, 0)
	for i := 0; i < len(tc); i++ {
		num, _ := strconv.Atoi(string(tc[i]))
		patterns = append(patterns, num)
	}

	changes := 0
	current := 0
	
	// Arah -1 = Kiri (CCW), 1 = Kanan (CW)
	lastDir := -1 

	for i := 0; i < len(patterns); i++ {
		target := patterns[i]
		var moveDir int

		if i == 0 {
			moveDir = -1
		} else {
			cw := countTotalClockWise(current, target)       // Kanan/Mundur
			ccw := countTotalCounterClockWise(current, target) // Kiri/Maju

			if cw < ccw {
				moveDir = 1 // Kanan
			} else if ccw < cw {
				moveDir = -1 // Kiri
			} else {
				moveDir = lastDir // Jarak sama, ikut arah sebelumnya
			}
		}

		fmt.Printf("Step %d | %d -> %d | Move: %v (Last: %v) | Change? ", i+1, current, target, moveDir, lastDir)

		if i > 0 && moveDir != lastDir {
			changes++
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		lastDir = moveDir
		current = target
	}

	return changes
}

func main() {
	testCases := []string{
		"981", 
		"4350",
		"12345",
		"2121",
	}
	for _, tc := range testCases {
		fmt.Printf("Input: %s, Total Changes: %d\n\n", tc, countTotalDirectionChange(tc))
	}
}