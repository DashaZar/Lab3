package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

// Решето Эратосфена
func sieve(limit int) []int {
	isPrime := make([]bool, limit)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	for i := 2; i*i < limit; i++ {
		if isPrime[i] {
			for j := i * i; j < limit; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := []int{}
	for i := 2; i < limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// Быстрое модульное возведение в степень
func modPow(a, b, mod int64) int64 {
	result := int64(1)
	a = a % mod
	for b > 0 {
		if b&1 == 1 {
			result = (result * a) % mod
		}
		a = (a * a) % mod
		b >>= 1
	}
	return result
}

// Факторизация числа на простые множители
func factorize(n int64) [][2]int64 {
	factors := [][2]int64{}
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			count := int64(0)
			for n%i == 0 {
				n /= i
				count++
			}
			factors = append(factors, [2]int64{i, count})
		}
	}
	if n > 1 {
		factors = append(factors, [2]int64{n, 1})
	}
	return factors
}

// Тест Миллера (по вашей логике)
func millerTest(n int64, t int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	m := n - 1
	factors := factorize(m)

	rand.Seed(time.Now().UnixNano())
	used := make(map[int64]bool)
	aList := []int64{}

	for len(aList) < t {
		a := rand.Int63n(n-3) + 2 // [2, n-2]
		if !used[a] {
			used[a] = true
			aList = append(aList, a)
		}
	}

	// Шаг 1: a^(n-1) mod n == 1
	for _, a := range aList {
		if modPow(a, m, n) != 1 {
			return false
		}
	}

	// Шаг 2: a^((n-1)/q) mod n != 1 для всех q
	for _, f := range factors {
		q := f[0]
		allOnes := true
		for _, a := range aList {
			if modPow(a, m/q, n) != 1 {
				allOnes = false
				break
			}
		}
		if allOnes {
			return false
		}
	}

	return true
}

// Генерация m для n = 2m + 1
func generateM(primes []int, targetBits int) int64 {
	rand.Seed(time.Now().UnixNano())
	m := int64(1)
	currentBits := 0

	for currentBits < targetBits-1 {
		q := int64(primes[rand.Intn(len(primes))])
		alpha := rand.Intn(3) + 1 // [1..3]
		term := int64(math.Pow(float64(q), float64(alpha)))

		if math.Log2(float64(m))+math.Log2(float64(term)) > float64(targetBits-1) {
			continue
		}
		m *= term
		currentBits = int(math.Log2(float64(m))) + 1
	}
	return m
}

// Генерация кандидата простого числа n = 2m + 1
func generatePrime(primes []int, digits int, t int) int64 {
	minVal := int64(math.Pow10(digits - 1))
	maxVal := int64(math.Pow10(digits)) - 1
	targetBits := int(math.Ceil(math.Log2(float64(maxVal))))

	for {
		m := generateM(primes, targetBits-1)
		n := 2*m + 1
		if n < minVal || n > maxVal {
			continue
		}
		if millerTest(n, t) {
			return n
		}
	}
}

// Тест Миллера-Рабина
func millerRabin(n int64, iterations int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	d := n - 1
	s := 0
	for d%2 == 0 {
		d /= 2
		s++
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < iterations; i++ {
		a := rand.Int63n(n-3) + 2
		x := modPow(a, d, n)

		if x == 1 || x == n-1 {
			continue
		}

		composite := true
		for j := 0; j < s-1; j++ {
			x = modPow(x, 2, n)
			if x == n-1 {
				composite = false
				break
			}
		}
		if composite {
			return false
		}
	}
	return true
}

func main() {
	primes := sieve(500)
	digits := 3
	t := 5
	count := 10
	k := 0

	numbers := make([]int64, 0, count)
	results := make([]string, 0, count)

	for i := 0; i < count; i++ {
		var p int64
		for {
			p = generatePrime(primes, digits, t)
			if p >= 100 && p <= 999 {
				break
			}
		}
		isPrime := millerRabin(p, 5)
		numbers = append(numbers, p)
		if isPrime {
			results = append(results, "+")
		} else {
			results = append(results, "-")
			k++
		}
	}

	// Вывод результатов
	fmt.Printf("%-5s", "№")
	for i := 1; i <= count; i++ {
		fmt.Printf("%-6d", i)
	}
	fmt.Println("\n-------------------------------------------------")

	fmt.Printf("%-5s", "P")
	for _, num := range numbers {
		fmt.Printf("%-6d", num)
	}
	fmt.Println()

	fmt.Printf("%-5s", "Test")
	for _, res := range results {
		fmt.Printf("%-6s", res)
	}
	fmt.Println("\n-------------------------------------------------")

	fmt.Printf("K = %d\n", k)
}
