package main

func q3p2(dirs string) int {
	sx, sy, rx, ry := 0, 0, 0, 0
	visited := make(map[string]bool)

	visited[toKey(sx, sy)] = true

	giftedHouses := 1
	roboTurn := true

	for _, c := range dirs {
		if roboTurn {
			rx, ry = nextLoc(rx, ry, string(c))
			if present, ok := visited[toKey(rx, ry)]; !ok || !present {
				visited[toKey(rx, ry)] = true
				giftedHouses++
			}
		} else {
			sx, sy = nextLoc(sx, sy, string(c))
			if present, ok := visited[toKey(sx, sy)]; !ok || !present {
				visited[toKey(sx, sy)] = true
				giftedHouses++
			}
		}
		roboTurn = !roboTurn
	}

	return giftedHouses
}
