package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"sync"
)

var kernel = [3][3]float64{
	{0.0625, 0.125, 0.0625},
	{0.125, 0.25, 0.125},
	{0.0625, 0.125, 0.0625},
}

func applyKernel(img image.Image, x, y int) color.RGBA {
	var r, g, b float64

	// Размер ядра свёртки (например, 3x3)
	kernelSize := len(kernel)

	for i := 0; i < kernelSize; i++ {
		for j := 0; j < kernelSize; j++ {
			// Координаты пикселей
			px := x + i - kernelSize/2
			py := y + j - kernelSize/2

			// Граничные проверки
			if px >= 0 && px < img.Bounds().Dx() && py >= 0 && py < img.Bounds().Dy() {
				rgba := color.RGBAModel.Convert(img.At(px, py)).(color.RGBA)
				r += float64(rgba.R) * kernel[i][j]
				g += float64(rgba.G) * kernel[i][j]
				b += float64(rgba.B) * kernel[i][j]
			}
		}
	}

	// Приведение к диапазону [0, 255]
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

func processRow(img image.Image, out *image.RGBA, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := 0; x < img.Bounds().Dx(); x++ {
		out.Set(x, y, applyKernel(img, x, y))
	}
}

func main() {
	// Открытие изображения
	file, err := os.Open("LogoBlue.png")
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Ошибка декодирования изображения: %v", err)
	}

	// Создание нового изображения для результата
	out := image.NewRGBA(img.Bounds())

	// Параллельная обработка строк
	var wg sync.WaitGroup
	for y := 0; y < img.Bounds().Dy(); y++ {
		wg.Add(1)
		go processRow(img, out, y, &wg)
	}
	wg.Wait()

	// Сохранение результата
	outFile, err := os.Create("output.png")
	if err != nil {
		log.Fatalf("Ошибка создания файла: %v", err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, out)
	if err != nil {
		log.Fatalf("Ошибка сохранения изображения: %v", err)
	}

	log.Println("Обработка завершена, изображение сохранено как output.png")
}
