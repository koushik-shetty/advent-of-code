package main

import (
	"fmt"
	"log"

	"github.com/koushik-shetty/advent-of-code/v2/2015/utils"
)

func main() {
	file := "test.txt"
	ip := utils.NewInput(file, 10000)
	secrets, err := ip.ReadChunk()
	if err != nil {
		log.Fatalf("Error: %#v", err.Error())
	}
	fmt.Printf("Smalles positive integer to have 5 zero md5 prefix: %#v\n", q4p1(secrets))
	fmt.Printf("Smalles positive integer to have 6 zero md5 prefix: %#v\n", q4p2(secrets))

}
