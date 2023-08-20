package main

import (
	"fmt"
	"log"
	"os"

	"github.com/koushik-shetty/advent-of-code/v2/2015/utils"
)

func main() {
	ip := utils.NewInput("q1.txt", 10000)
	text, err := ip.ReadChunk()
	if err != nil {
		log.Fatalf("Err: %#v\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Final Floor: %#v\n", finalFloor(text))
	fmt.Printf("Position that causes santa to go to basement: %#v\n", finalFloorP2(text))
}
