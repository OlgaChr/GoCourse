package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(count int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		number := r.Intn(50)
		fmt.Print(number*number, " ")
	}
	fmt.Println()
}

func main() {
	var i int
	fmt.Println("Введите количество генерируемых квадратов: ")
	fmt.Scanf("%d", &i)
	generator(i)
}
