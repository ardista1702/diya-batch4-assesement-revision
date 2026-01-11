package main

import "strings"

func buildGrid(str string, totalCol, totalRow int) [][]string {
	grid := make([][]string, totalRow)
	for i := 0; i < totalRow; i += totalCol {
		grid[i] = make([]string, totalCol)
		for j := 0; j < totalCol; j++ {
			grid[i][j] = string(str[i*totalCol+j])
		}
	}
	return grid
}
func search(strs []string, target string) bool {
	window := ""
	for _, str := range strs {
		window += str
	}
	if window == target {
		return true
	}
	for j := len(target); j < len(window); j++ {
		if window[j-len(target):j] == target {
			return true
		}
	}
	return false
}
func reverse(strs []string) {
	l, r := 0, len(strs)-1
	for l < r {
		strs[l], strs[r] = strs[r], strs[l]
		l++
		r--
	}
}
func searchRow(grid [][]string, target string) bool {
	for _, row := range grid {
		if search(row, target) {
			return true
		}
		reverse(row)
		if search(row, target) {
			return true
		}
	}
	return false
}
func searchCol(grid [][]string, target string) bool {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		col := make([]string, n)
		for j := 0; j < n; j++ {
			col[j] = grid[j][i]
		}
		if search(col, target) {
			return true
		}
		reverse(col)
		if search(col, target) {
			return true
		}
	}
	return false
}
func isTargetExist(testCase string) string {
	splittedTc := strings.Split(testCase, " ")
	target := splittedTc[0]
	str := splittedTc[1]
	grid := buildGrid(str, 6, 6)
	if searchRow(grid, target) {
		return "YA"
	}
	if searchCol(grid, target) {
		return "YA"
	}
	return "TIDAK"
}
func main() {
	testCase := []string{
		"MELATI MELATIABCDEFGHIJKLMNOPQRSTUVWXYZABCD",
		"MAWAR MELATIABCDEFGHIJKLMNOPQRSTUVWXYZABCD",
		// "MATAHARI MELATIABCDEFGHIJKLMNOPQRSTUVWXYZABCD",
	}
	for _, tc := range testCase {
		println(isTargetExist(tc))
	}
}
