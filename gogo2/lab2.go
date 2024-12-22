package main

import (
	"errors"
	"fmt"
	"math"
)

// Задание 1. Массивы и срезы

// 1. Функция для форматирования IP-адреса
func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// 2. Функция для получения всех четных чисел в диапазоне
func listEven(start, end int) ([]int, error) {
	if start > end {
		return nil, errors.New("левая граница диапазона больше правой")
	}
	var evens []int
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	return evens, nil
}

// Задание 2. Карты

// 1. Функция для подсчета вхождений каждого символа
func countCharacters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	return counts
}

// Задание 3. Структуры, методы и интерфейсы

// 1. Структура "Точка"
type Point struct {
	X, Y float64
}

// 2. Структура "Отрезок"
type Segment struct {
	Start, End Point
}

// 3. Метод для вычисления длины отрезка
func (s Segment) Length() float64 {
	dx := s.End.X - s.Start.X
	dy := s.End.Y - s.Start.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// 4. Структура "Треугольник"
type Triangle struct {
	A, B, C Point
}

// 5. Структура "Круг"
type Circle struct {
	Center Point
	Radius float64
}

// 6. Метод для вычисления площади треугольника
func (t Triangle) Area() float64 {
	a := Segment{t.A, t.B}.Length()
	b := Segment{t.B, t.C}.Length()
	c := Segment{t.C, t.A}.Length()
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

// Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 7. Интерфейс "Фигура"
type Shape interface {
	Area() float64
}

// 8. Функция для вывода площади фигуры
func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}

// Задание 4. Функциональное программирование

// 1. Функция Map
func Map(slice []float64, fn func(float64) float64) []float64 {
	result := make([]float64, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// 2. Функция для возведения в квадрат
func square(x float64) float64 {
	return x * x
}

func main() {
	// Тестирование функций из Задания 1
	ip := [4]byte{127, 0, 0, 1}
	fmt.Println("IP адрес:", formatIP(ip))

	evens, err := listEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Четные числа:", evens)
	}

	// Тестирование функций из Задания 2
	str := "hello, world"
	fmt.Println("Подсчет символов:", countCharacters(str))

	// Тестирование структур и методов из Задания 3
	triangle := Triangle{Point{0, 0}, Point{4, 0}, Point{0, 3}}
	circle := Circle{Point{0, 0}, 5}

	printArea(triangle)
	printArea(circle)

	// Тестирование функций из Задания 4
	numbers := []float64{1, 2, 3, 4, 5}
	squaredNumbers := Map(numbers, square)
	fmt.Println("Исходные числа:", numbers)
	fmt.Println("Квадраты чисел:", squaredNumbers)
}
