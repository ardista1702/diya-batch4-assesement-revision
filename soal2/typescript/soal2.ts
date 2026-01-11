function buildGrid(str: string, totalCol: number, totalRow: number): string[][] {
	let grid:string[][] = []
	for (let i = 0; i < totalRow; i++) {
		grid[i] = []
		for (let j = 0; j < totalCol; j++) {
			grid[i][j] = str[i*totalCol+j]
		}
	}
	return grid
}
function search(strs: string[], target: string): boolean {
	let window = ""
	for (let str of strs) {
		window += str
	}
	if (window == target) {
		return true
	}
	for (let j = target.length; j < window.length; j++) {
        window += strs[j]
        window = window.slice(1)
        if (window == target) {
			return true
		}
	}
	return false
}
function reverse(strs: string[]) {
	let l = 0, r = strs.length-1
	while (l < r) {
		strs[l], strs[r] = strs[r], strs[l]
		l++
		r--
	}
}
function searchRow(grid: string[][], target: string): boolean {
	for (let row of grid) {
		if (search(row, target)) {
			return true
		}
		reverse(row)
		if (search(row, target)) {
			return true
		}
	}
	return false
}
function searchCol(grid :string[][], target :string): boolean {
	const m = grid.length
	const n = grid[0].length
	for (let i = 0; i < m; i++) {
		let col :string[] = []
		for (let j = 0; j < n; j++) {
			col[j] = grid[j][i]
		}
		if (search(col, target)) {
			return true
		}
		reverse(col)
		if (search(col, target)) {
			return true
		}
	}
	return false
}
function isTargetExist(testCase: string):string {
	const splittedTc = testCase.split(" ")
	const target = splittedTc[0]
	const str = splittedTc[1]
	const grid = buildGrid(str, 6, 6)
	if (searchRow(grid, target)) {
		return "YA"
	}
	if (searchCol(grid, target)) {
		return "YA"
	}
	return "TIDAK"
}
