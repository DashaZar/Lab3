package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	// Чтение входных данных
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)

	// Проверка ограничений
	if n < 5 || n > 50000 {
		fmt.Println("n must be between 5 and 50000")
		return
	}
	if m < 4 || m > 100 {
		fmt.Println("m must be between 4 and 100")
		return
	}

	// Чтение массива чисел
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	// Создание массива префиксных сумм
	// prefix[i] = сумма элементов nums[0] до nums[i-1]
	prefix := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + int64(nums[i-1])
	}

	// Массив для динамического программирования
	// dp[i] - максимальная разница очков текущего игрока и оппонента
	// при игре на подмассиве nums[i:]
	dp := make([]int64, n+1)

	// Обратный проход по массиву (от последнего элемента к первому)
	for i := n - 1; i >= 0; i-- {
		// Инициализация минимально возможным значением
		maxScore := int64(math.MinInt64)

		// Перебор всех возможных ходов (1..m чисел)
		for j := 1; j <= m; j++ {
			// Проверка выхода за границы массива
			if i+j > n {
				break
			}

			// Вычисление суммы текущего хода
			currentSum := prefix[i+j] - prefix[i]
			
			// Расчет преимущества для текущего игрока:
			// currentSum - очки, взятые текущим игроком
			// dp[i+j] - преимущество оппонента на оставшейся части
			currentScore := currentSum - dp[i+j]
			
			// Обновление максимального преимущества
			if currentScore > maxScore {
				maxScore = currentScore
			}
		}
		dp[i] = maxScore
	}

	// Определение и вывод результата
	switch {
	case dp[0] > 0:
		fmt.Println(1) // Победа первого игрока
	case dp[0] < 0:
		fmt.Println(0) // Победа второго игрока
	default:
		fmt.Println("ничья") // Ничья
	}
}
