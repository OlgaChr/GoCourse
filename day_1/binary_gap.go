package main

import (
	"strconv"
)

func Solution(N int) int {
	count := 0
	max := 0
	binaryNumber := strconv.FormatInt(int64(N), 2)

	for _, letter := range binaryNumber {
		if string(letter) == "0" {
			count++
		} else {
			if count > max {
				max = count
				count = 0
			}
		}
	}
	if count > max {
		max = count
	}
	return max
}
