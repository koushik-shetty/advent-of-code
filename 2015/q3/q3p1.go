package main

import "fmt"

func q3p1(dirs string) int {
	x, y := 0, 0
	visited := make(map[string]bool)
	visited[toKey(x, y)] = true

	giftedHouses := 1
	for _, c := range dirs {
		x, y = nextLoc(x, y, string(c))
		if present, ok := visited[toKey(x, y)]; !ok || !present {
			visited[toKey(x, y)] = true
			giftedHouses++
		}
	}

	return giftedHouses
}

func toKey(x, y int) string { return fmt.Sprintf("%v_%v", x, y) }

func nextLoc(x, y int, dir string) (int, int) {
	switch dir {
	case "^":
		return x + 1, y
	case "v":
		return x - 1, y
	case "<":
		return x, y - 1
	case ">":
		return x, y + 1
	}
	return x, y
}
