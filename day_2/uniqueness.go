package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Println("Введите длину массива:")
	fmt.Scanf("%d", &n)
	
	a := make([]int, 0, n)
	m := make(map[int]bool, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i < n; i++ {
		number := r.Intn(10)
		a = append(a, number)
		m[number] = true
	}
	fmt.Println("Массив:")
	fmt.Println(a)
	fmt.Println("Количество уникальных чисел:")
	fmt.Println(len(m))
}