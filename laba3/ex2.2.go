package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Возведение в степень по модулю
func powMod(base, degree, mod int) int {
	result := 1
	base %= mod
	for degree > 0 {
		if degree%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		degree /= 2
	}
	return result
}

// Решето Эратосфена
func sieveEratos(N int) []int {
	isPrime := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		isPrime[i] = true
	}
	primes := []int{}
	for i := 2; i <= N; i++ {
		if isPrime[i] {
			primes = append(primes, i)
			for j := i * 2; j <= N; j += i {
				isPrime[j] = false
			}
		}
	}
	return primes
}

// Тест Миллера-Рабина
func rabin(num, k int) bool {
	if num < 2 {
		return false
	}
	if num == 2 || num == 3 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	// Представляем num-1 в виде d*2^s
	s := 0
	d := num - 1
	for d%2 == 0 {
		d /= 2
		s++
	}

	for i := 0; i < k; i++ {
		a := 2 + rand.Intn(num-3) // [2, num-2]
		x := powMod(a, d, num)
		if x == 1 || x == num-1 {
			continue
		}

		passed := false
		for r := 1; r < s; r++ {
			x = powMod(x, 2, num)
			if x == num-1 {
				passed = true
				break
			}
		}
		if !passed {
			return false
		}
	}
	return true
}

// Генерация случайного числа в диапазоне [min, max]
func randDist(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// Тест Поклингтона
func poklington(n, t int, qList []int) bool {
	aSet := make(map[int]struct{})

	for len(aSet) < t {
		a := randDist(2, n-2)
		aSet[a] = struct{}{}
	}

	for a := range aSet {
		if powMod(a, n-1, n) != 1 {
			return false
		}
	}

	for a := range aSet {
		conditionMet := true
		for _, q := range qList {
			if powMod(a, (n-1)/q, n) == 1 {
				conditionMet = false
				break
			}
		}
		if conditionMet {
			return true
		}
	}
	return false
}

// Возведение числа в степень с float64 (для pow)
func intPow(base, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}

// Генерация кандидата p и списка q
func calcN(primes []int, bit int) (int, []int) {
	minBitF := bit/2 + 1
	maxBitF := bit/2 + 2
	minF := 1 << minBitF
	maxF := 1 << maxBitF

	F := 1
	qList := []int{}
	usedPrimes := make(map[int]struct{})

	for F < minF {
		prime := primes[randDist(0, len(primes)-1)]
		if _, exists := usedPrimes[prime]; exists {
			continue
		}

		alpha := randDist(1, 3)
		powValue := intPow(prime, alpha)

		if F*powValue >= maxF {
			break
		}

		F *= powValue
		qList = append(qList, prime)
		usedPrimes[prime] = struct{}{}
	}

	R := randDist(2, 10) * 2 // Чётное от 4 до 20 (2*2..10*2)
	p := R*F + 1

	return p, qList
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const bit = 10
	const t = 10

	primes := sieveEratos(500)
	results := []int{}
	testResults := []string{}

	for len(results) < 10 {
		p, qList := calcN(primes, bit)

		if p < 100 || p > 999 {
			continue
		}

		if poklington(p, t, qList) {
			results = append(results, p)
			if rabin(p, 3) {
				testResults = append(testResults, "+")
			} else {
				testResults = append(testResults, "-")
			}
		}
	}

	fmt.Printf("%-5s %-10s %-6s\n", "№", "P", "Test")
	fmt.Println("---------------------------")
	for i := 0; i < 10; i++ {
		fmt.Printf("%-5d %-10d %-6s\n", i+1, results[i], testResults[i])
	}

	k := 0
	for _, res := range testResults {
		if res == "-" {
			k++
		}
	}

	fmt.Println("---------------------------")
	fmt.Printf("Количество не прошедших тест: %d\n", k)
}

