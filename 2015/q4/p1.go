package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

func q4p1(secret string) int {
	for smallestNum := 0; smallestNum <= math.MaxInt; smallestNum++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%v", secret, smallestNum)))

		hexCode := hex.EncodeToString(hash[:])

		if strings.HasPrefix(hexCode, "00000") {
			return smallestNum
		}
	}
	return 0
}
