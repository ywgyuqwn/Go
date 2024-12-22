package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Задание 1: Обработка query-параметров
func greetHandler(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	if name == "" || age == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметры name и age обязательны"})
		return
	}

	c.String(http.StatusOK, "Меня зовут %s, мне %s лет", name, age)
}

// Задание 2: Маршрутизация для арифметических операций

// Функция для обработки арифметических операций
func handleOperation(c *gin.Context, operation string) {
	aStr := c.Query("a")
	bStr := c.Query("b")

	a, errA := strconv.ParseFloat(aStr, 64)
	b, errB := strconv.ParseFloat(bStr, 64)

	if errA != nil || errB != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Оба параметра a и b должны быть числами"})
		return
	}

	var result float64
	switch operation {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		if b == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Деление на ноль невозможно"})
			return
		}
		result = a / b
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

// Обработчики для каждого маршрута
func addHandler(c *gin.Context) {
	handleOperation(c, "add")
}

func subHandler(c *gin.Context) {
	handleOperation(c, "sub")
}

func mulHandler(c *gin.Context) {
	handleOperation(c, "mul")
}

func divHandler(c *gin.Context) {
	handleOperation(c, "div")
}

// Задание 3: Обработка JSON для подсчета символов
func countCharactersHandler(c *gin.Context) {
	var requestBody struct {
		Text string `json:"text"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON-формат"})
		return
	}

	counts := make(map[rune]int)
	for _, char := range requestBody.Text {
		counts[char]++
	}

	c.JSON(http.StatusOK, counts)
}

func main() {
	router := gin.Default()

	// Задание 1: Маршрут для приветствия с использованием query-параметров
	router.GET("/greet", greetHandler)

	// Задание 2: Маршруты для арифметических операций
	router.GET("/add", addHandler)
	router.GET("/sub", subHandler)
	router.GET("/mul", mulHandler)
	router.GET("/div", divHandler)

	// Задание 3: Маршрут для подсчета символов в строке
	router.POST("/count", countCharactersHandler)

	router.Run(":8080")
}
