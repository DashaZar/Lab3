package main

import (
	"fmt"
	"math"
)

// Функция для вычисления наибольшего общего делителя (НОД)
func gcd(a, b int64) int64 {
	// Сделаем числа положительными
	a = abs(a)
	b = abs(b)
	
	// Алгоритм Евклида
	for b != 0 {
		temp := a % b
		a = b
		b = temp
	}
	return a
}

// Вспомогательная функция для получения абсолютного значения
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// Функция для быстрого возведения в степень
func power(base int64, exp int) int64 {
	// Любое число в степени 0 равно 1
	if exp == 0 {
		return 1
	}
	
	result := int64(1)
	currentBase := base
	currentExp := exp
	
	for currentExp > 0 {
		// Если степень нечётная, умножаем результат на текущее основание
		if currentExp%2 == 1 {
			result *= currentBase
		}
		// Возводим основание в квадрат
		currentBase *= currentBase
		// Уменьшаем степень вдвое
		currentExp /= 2
	}
	
	return result
}

// Функция для вычисления чисел Эйлера
func computeEuler(maxN int) [][]int64 {
	// Создаем 2D-срез размером (maxN+1) x (maxN+1)
	euler := make([][]int64, maxN+1)
	for i := range euler {
		euler[i] = make([]int64, maxN+1)
	}
	
	// Базовый случай: для пустой перестановки
	euler[0][0] = 1
	
	// Заполняем таблицу чисел Эйлера
	for n := 1; n <= maxN; n++ {
		// Перебираем все возможные количества подъемов (k)
		for k := 0; k < n; k++ {
			var term1 int64
			// Первое слагаемое: (n - k) * euler[n-1][k-1]
			if k > 0 {
				term1 = int64(n-k) * euler[n-1][k-1]
			}
			
			// Второе слагаемое: (k + 1) * euler[n-1][k]
			term2 := int64(k+1) * euler[n-1][k]
			
			// Суммируем оба слагаемых
			euler[n][k] = term1 + term2
		}
	}
	
	return euler
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	
	// Проверка диапазона входных значений
	if a < 1 || a > 10 || b < 1 || b > 10 {
		fmt.Println("Ошибка: числа должны быть от 1 до 10")
		return
	}
	
	// Обработка расходящегося ряда (b = 1)
	if b == 1 {
		fmt.Println("infinity")
		return
	}
	
	// Вычисляем числа Эйлера для заданного a
	euler := computeEuler(a)
	
	// Вычисление числителя дроби
	var numerator int64
	for j := 0; j < a; j++ {
		// Слагаемое: euler[a][j] * b^(a-j)
		numerator += euler[a][j] * power(int64(b), a-j)
	}
	
	// Вычисление знаменателя дроби
	denominator := power(int64(b-1), a+1)
	
	// Сокращение дроби
	common := gcd(numerator, denominator)
	numerator /= common
	denominator /= common
	
	// Вывод результата
	fmt.Printf("%d/%d\n", numerator, denominator)
}
