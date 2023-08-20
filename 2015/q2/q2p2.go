package main

import (
	"fmt"
	"sort"

	"github.com/koushik-shetty/advent-of-code/v2/2015/utils"
)

func q2p2(file string) int {
	ip := utils.NewInput(file, 100)
	totalRibbon := 0
	rl := 0
	s := [3]int{}
	for {
		if lines, err := ip.ReadBatch(); err != nil {
			panic(err)
		} else if len(lines) == 0 {
			break
		} else {
			rl += len(lines)
			for _, line := range lines {
				l, w, h := 0, 0, 0
				fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
				s[0] = l
				s[1] = w
				s[2] = h
				sort.Ints(s[:])
				totalRibbon += 2*s[0] + 2*s[1] + l*w*h
			}
		}
	}

	return totalRibbon
}
