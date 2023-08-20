package main

import (
	"log"

	"github.com/koushik-shetty/advent-of-code/v2/2015/utils"
)

func q3p1(file string) int {
	ip := utils.NewInput(file, 10000)
	dirs, err := ip.ReadChunk()
	if err != nil {
		log.Fatalf("Error: %#v", err.Error())
	}
	curr := []int{0, 0}
	visited := make(map[int]map[int]bool)
	visited[curr[0]] = make(map[int]bool)
	visited[curr[0]][curr[1]] = true

	giftedHouses := 1
	for _, c := range dirs {
		nextLoc(curr, string(c))
		x, y := curr[0], curr[1]
		if xaxis, ok := visited[x]; !ok {
			visited[x] = make(map[int]bool)
			visited[x][y] = true
			giftedHouses++
		} else if yaxis, k := xaxis[y]; !yaxis || !k {
			visited[x][y] = true
			giftedHouses++
		}
	}

	return giftedHouses
}

func nextLoc(start []int, dir string) {
	switch dir {
	case "^":
		start[0] += 1
	case "v":
		start[0] -= 1
	case "<":
		start[1] -= 1
	case ">":
		start[1] += 1
	}
}
