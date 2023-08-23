package main

func finalFloor(instructions string) int {
	finalFloor := 0
	for _, c := range instructions {
		if c == '(' {
			finalFloor++
		} else if c == ')' {
			finalFloor--
		}
	}
	return finalFloor
}
