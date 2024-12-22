package main

import (
	"fmt"
	"image/color"
	"image/draw"
	"image/png"
	_ "image/png"
	"os"
	"time"
)

func filter(drawImg draw.RGBA64Image) {
	bounds := drawImg.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			px := drawImg.RGBA64At(x, y)
			avg := uint16((uint32(px.R) + uint32(px.G) + uint32(px.B)) / 3)
			drawImg.SetRGBA64(x, y, color.RGBA64{R: avg, G: avg, B: avg, A: px.A})
		}
	}
}

func main() {
	// Открываем файл изображения
	inFile, err := os.Open("input.png")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer inFile.Close()

	// Декодируем PNG
	img, err := png.Decode(inFile)
	if err != nil {
		fmt.Println("Ошибка декодирования PNG:", err)
		return
	}

	// Пытаемся преобразовать к типу draw.RGBA64Image
	drawImg, ok := img.(draw.RGBA64Image)
	if !ok {
		fmt.Println("Не удалось преобразовать изображение к draw.RGBA64Image")
		return
	}

	// Замеряем время обработки
	start := time.Now()

	// Применяем фильтр
	filter(drawImg)

	elapsed := time.Since(start)
	fmt.Printf("Время обработки: %v\n", elapsed)

	// Создаём выходной файл
	outFile, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Ошибка создания выходного файла:", err)
		return
	}
	defer outFile.Close()

	// Сохраняем результат
	err = png.Encode(outFile, drawImg)
	if err != nil {
		fmt.Println("Ошибка сохранения обработанного изображения:", err)
		return
	}

	fmt.Println("Обработка завершена. Результат в output.png")
}
