package main

func finalFloorP2(instructions string) int {
	finalFloor := 0
	for k, c := range instructions {
		if c == '(' {
			finalFloor++
		} else if c == ')' {
			finalFloor--
		}
		if finalFloor == -1 {
			return k + 1
		}
	}
	return finalFloor
}
