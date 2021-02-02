package main

import (
	"fmt"
	"strconv"
)

func solution(n int) int {
	count := 0
	max := 0
	binaryNumber := strconv.FormatInt(int64(n), 2)

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

func main() {
	var i int
	fmt.Println("Введите число: ")
	fmt.Scanf("%d", &i)
	fmt.Println(solution(i))
}
