package main

import (
	"fmt"
	"time"
)

func count(ch <-chan int) {
	for num := range ch {
		// Произвольное действие: возведение числа в квадрат
		fmt.Println(num * num)
	}
}

func main() {
	ch := make(chan int)

	go count(ch)

	// Отправляем несколько чисел в канал
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

	// Добавляем паузу, чтобы горутина успела вывести результаты
	time.Sleep(time.Second)
}
