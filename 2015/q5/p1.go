package main

import (
	"fmt"
	"log"

	"github.com/koushik-shetty/advent-of-code/v2/2015/utils"
)

func q4p1(file string) int {

	ip := utils.NewInput(file, 100)
	goodStrings := 0
	for {
		if strs, err := ip.ReadBatch(); err != nil {
			log.Fatalf("Error: %#v", err.Error())
			break
		} else if len(strs) == 0 {
			fmt.Printf("done")
			break
		} else {
			for _, str := range strs {
				if isGoodString(str) {
					goodStrings++
				}
			}
		}
	}

	return goodStrings
}

var vowels = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
}

var blacklist = map[string]bool{
	"ab": true,
	"cd": true,
	"pq": true,
	"xy": true,
}

func isGoodString(str string) bool {
	if len(str) == 0 {
		return false
	}
	hasTwinChars := false
	vowelCount := 0
	prevChar := ""

	remaining := len(str) % 2

	strLenToProcess := len(str[:len(str)-remaining])
	for i := 0; i < strLenToProcess; i += 2 {
		c1, c2 := string(str[i]), string(str[i+1])
		if blacklist[c1+c2] || blacklist[prevChar+c1] {
			return false
		}

		if ok := vowels[c1]; ok {
			vowelCount++
		}
		if ok := vowels[c2]; ok {
			vowelCount++
		}

		if c1 == c2 || c1 == prevChar {
			hasTwinChars = true
		}

		prevChar = c2
	}
	if remaining > 0 {
		ch := string(str[len(str)-1])
		prev := string(str[len(str)-2])
		if vowels[ch] {
			vowelCount++
		}
		if ch == prev {
			hasTwinChars = true
		}
		if blacklist[prev+ch] {
			return false
		}
	}

	return vowelCount >= 3 && hasTwinChars
}
