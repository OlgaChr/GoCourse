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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i < n; i++ {
		a = append(a, r.Intn(50))
	}
	fmt.Println("Массив:")
	fmt.Println(a)
	var k int
	fmt.Println("Введите количество сдвигов: ")
	fmt.Scanf("%d", &k)
	a = append(a[n-k:], a[:n-k]...)
	fmt.Println("Сдвинутый массив:")
	fmt.Println(a)
}