package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Быстрое возведение в степень по модулю
func powMod(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}

// Тест Миллера-Рабина на простоту
func isProbablePrime(n, k int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Представляем n-1 = 2^r * d
	d := n - 1
	r := 0
	for d%2 == 0 {
		d /= 2
		r++
	}

	for i := 0; i < k; i++ {
		a := rand.Intn(n-3) + 2 // случайное число от 2 до n-2
		x := powMod(a, d, n)

		if x == 1 || x == n-1 {
			continue
		}

		cont := false
		for j := 0; j < r-1; j++ {
			x = powMod(x, 2, n)
			if x == n-1 {
				cont = true
				break
			}
		}

		if cont {
			continue
		}

		return false
	}
	return true
}

// Генерация простых чисел до maxN (Решето Эратосфена)
func generateSmallPrimes(maxN int) []int {
	isPrime := make([]bool, maxN+1)
	for i := 2; i <= maxN; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= maxN; i++ {
		if isPrime[i] {
			for j := i * i; j <= maxN; j += i {
				isPrime[j] = false
			}
		}
	}
	primes := []int{}
	for i := 2; i <= maxN; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// Генерация случайного float64 [0,1)
func randDouble() float64 {
	return rand.Float64()
}

// Генерация случайного целого в диапазоне [a,b]
func randInt(a, b int) int {
	return rand.Intn(b-a+1) + a
}

// Генерация простого числа по алгоритму ГОСТ
func generateGOSTPrime(smallPrimes []int, t int) int {
	qBitLen := (t + 1) / 2
	qMin := 1 << (qBitLen - 1)
	qMax := (1 << qBitLen) - 1

	attempts := 0

	for attempts < 10000 {
		attempts++
		q := smallPrimes[randInt(0, len(smallPrimes)-1)]
		if q < qMin || q > qMax {
			continue
		}

		xi := randDouble()
		N := int((float64(1<<(t-1)) + xi*float64(1<<(t-1))) / float64(q))
		if N%2 != 0 {
			N++
		}

		for u := 0; u < 1000; u += 2 {
			p := (N + u)*q + 1

			if p < 100 || p > 999 {
				continue
			}

			if powMod(2, p-1, p) != 1 {
				continue
			}

			if powMod(2, N+u, p) == 1 {
				continue
			}

			return p
		}
	}
	return -1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	t := 10
	smallPrimes := generateSmallPrimes(100)

	experimentCount := 10

	pList := make([]int, 0, experimentCount)
	testResults := make([]rune, 0, experimentCount)
	rejectedCount := 0

	for i := 0; i < experimentCount; i++ {
		p := generateGOSTPrime(smallPrimes, t)
		if p == -1 {
			fmt.Println("Ошибка генерации простого числа.")
			i--
			continue
		}
		pList = append(pList, p)
		if isProbablePrime(p, 5) {
			testResults = append(testResults, '+')
		} else {
			testResults = append(testResults, '-')
			rejectedCount++
		}
	}

	// Вывод результатов
	fmt.Printf("№        ")
	for i := 1; i <= experimentCount; i++ {
		fmt.Printf("%6d", i)
	}
	fmt.Println()

	fmt.Printf("p        ")
	for _, p := range pList {
		fmt.Printf("%6d", p)
	}
	fmt.Println()

	fmt.Printf("Результат")
	for _, res := range testResults {
		fmt.Printf("%6c", res)
	}
	fmt.Println()

	fmt.Printf("K        %d\n", rejectedCount)
}
