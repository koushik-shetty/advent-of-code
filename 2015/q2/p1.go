package main

import (
	"fmt"

	"github.com/koushik-shetty/advent-of-code/v2/2015/utils"
)

func q2p1(file string) int {
	min := utils.Min
	ip := utils.NewInput(file, 100)
	totalPaper := 0
	rl := 0
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
				totalPaper += min(h*l, min(l*w, w*h)) + 2*l*w + 2*w*h + 2*h*l
			}
		}
	}

	return totalPaper
}
