
function countTotalClockWise(current: number, target: number): number {
	let total = 0
	while (current != target) {
		current--
		if (current < 0) {
			current = 9
		}
		total++
	}
	return total
}

function countTotalCounterClockWise(current: number, target: number): number {
	let total = 0
	while (current != target) {
		current++
		if (current > 9) {
			current = 0
		}
		total++
	}
	return total
}

function countTotalDirectionChange(tc: string): number {
	let patterns: number[] = []
	for (let i = 0; i < tc.length; i++) {
		patterns.push(parseInt(tc[i]))
	}

	let changes = 0
	let current = 0
	
	// Arah -1 = Kiri (CCW), 1 = Kanan (CW)
	let lastDir = -1 

	for (let i = 0; i < patterns.length; i++) {
		let target = patterns[i]
		let moveDir = 0

		if (i == 0) {
			moveDir = -1
		} else {
			let cw = countTotalClockWise(current, target)       // Kanan/Mundur
			let ccw = countTotalCounterClockWise(current, target) // Kiri/Maju

			if (cw < ccw) {
				moveDir = 1 // Kanan
			} else if (ccw < cw) {
				moveDir = -1 // Kiri
			} else {
				moveDir = lastDir // Jarak sama, ikut arah sebelumnya
			}
		}

		if (i > 0 && moveDir != lastDir) {
			changes++
		}

		lastDir = moveDir
		current = target
	}

	return changes
}

function solve() {
	const testCases : string[] = [
		"981", 
		"4350",
		"12345",
		"2121",
	]
	for (let tc of testCases) {
		console.log(`Input: ${tc}, Total Changes: ${countTotalDirectionChange(tc)}\n\n`)
	}
}

solve()