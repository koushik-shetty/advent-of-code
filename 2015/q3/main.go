package main

import (
	"fmt"
	"log"

	"github.com/koushik-shetty/advent-of-code/v2/2015/utils"
)

func main() {
	file := "test.txt"
	ip := utils.NewInput(file, 10000)
	dirs, err := ip.ReadChunk()
	if err != nil {
		log.Fatalf("Error: %#v", err.Error())
	}
	fmt.Printf("No of house receiving presents: %#v\n", q3p1(dirs))
	fmt.Printf("No of house receiving presents with santa and robo alternating: %#v\n", q3p2(dirs))
}
