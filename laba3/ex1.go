package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	Xstart := -6.0
	Xend := 8.0
	dx := 0.2

	// Открываем файл для записи (создаётся новый или перезаписывается)
	file, err := os.Create("output_table.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Заголовок таблицы
	header := fmt.Sprintf("  x\t|\ty\n-------------------\n")
	fmt.Print(header)
	_, err = file.WriteString(header)
	if err != nil {
		panic(err)
	}

	for x := Xstart; x <= Xend; x += dx {
		var y float64

		if x <= -2.0 {
			y = x + 3 // участок 1
		} else if x > -2.0 && x < 1.0 {
			y = math.Pow(3, x) // участок 2
		} else if x >= 1.0 && x <= 5.0 {
			R := 3.0
			centerX, centerY := 3.0, 3.0
			expr := R*R - (x-centerX)*(x-centerX)
			if expr >= 0 {
				y = centerY - math.Sqrt(expr) // нижняя полусфера
			} else {
				y = math.NaN()
			}
		} else if x > 5.0 {
			y = -1.5*x + 10.5 // участок 4
		}

		// Форматируем с двумя знаками после запятой
		line := fmt.Sprintf("%5.2f\t|\t%5.2f\n", x, y)
		fmt.Print(line)
		_, err = file.WriteString(line)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("\nРезультаты также записаны в файл output_table.txt")
}
