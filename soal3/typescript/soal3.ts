
type coordinate = {
	X: number,
	Y: number
}

function isIntersection(coordiate1: coordinate, coordinate2: coordinate): boolean {
	return coordiate1.X == coordinate2.X && coordiate1.Y == coordinate2.Y
}

function move(direction: string, c: coordinate): coordinate {
	switch (direction) {
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

function countTotalIntersection(testcase: string): number {
	let totalIntersection = 1
	let splittedTc = testcase.split(" ")
	let path1 = splittedTc[0]
	let path2 = splittedTc[1]
	let p1 = 0
	let p2 = 0
	let c1 = {X: 0, Y: 0}
	let c2 = {X: 0, Y: 0}
	while (p1 < path1.length && p2 < path2.length) {
		c1 = move(path1[p1], c1)
		c2 = move(path2[p2], c2)
		if (isIntersection(c1, c2)) {
			totalIntersection++
		}
		p1++
		p2++
	}

	return totalIntersection
}

function solve():void {
	const testCases = [
		"UUU TTT",
		"BBUU BBUU",
		"TTSSSSS SSTTTTT",
	]
	for (let tc of testCases) {
		console.log(countTotalIntersection(tc))
	}
}
solve()
