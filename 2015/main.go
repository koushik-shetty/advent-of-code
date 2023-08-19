package main

import (
	"fmt"
)

func main() {
	// f, err := os.ReadFile("q1.txt")
	// if err != nil {
	// 	fmt.Printf("Error: %#v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("Final Floor: %#v\n", finalFloor(string(f)[:len(f)-1]))
	// fmt.Printf("Final Floor: %#v\n", finalFloorP2(string(f)[:len(f)-1]))
	// fmt.Printf("Total Paper: %#v\n", q2p1("q2p1.txt"))
	fmt.Printf("Total Ribbon: %#v\n", q2p2("q2p1.txt"))
}
