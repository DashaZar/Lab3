/*
 * Программа для генерации и проверки простых чисел с использованием теста Миллера-Рабина.
 * Формат вывода соответствует заданному шаблону с таблицей результатов.
 */

#include <iostream>
#include <vector>
#include <random>
#include <iomanip>
#include <cmath>
#include <unordered_set>

using namespace std;

// ===============================================
// 1. Решето Эратосфена для поиска простых чисел < limit
// ===============================================
vector<int> sieve(int limit) {
    vector<bool> is_prime(limit, true);
    is_prime[0] = is_prime[1] = false;

    for (int i = 2; i * i < limit; ++i) {
        if (is_prime[i]) {
            for (int j = i * i; j < limit; j += i)
                is_prime[j] = false;
        }
    }

    vector<int> primes;
    for (int i = 2; i < limit; ++i) {
        if (is_prime[i]) primes.push_back(i);
    }

    return primes;
}

// ===============================================
// 2. Быстрое модульное возведение в степень (a^b % mod)
// ===============================================
long long mod_pow(long long a, long long b, long long mod) {
    long long result = 1;
    a %= mod;

    while (b > 0) {
        if (b & 1)
            result = (result * a) % mod;
        a = (a * a) % mod;
        b >>= 1;
    }

    return result;
}

// ===============================================
// 3. Факторизация числа (n = q1^α1 * q2^α2 * ...)
// ===============================================
vector<pair<long long, int>> factorize(long long n) {
    vector<pair<long long, int>> factors;

    for (long long i = 2; i * i <= n; ++i) {
        if (n % i == 0) {
            int count = 0;
            while (n % i == 0) {
                n /= i;
                count++;
            }
            factors.emplace_back(i, count);
        }
    }

    if (n > 1)
        factors.emplace_back(n, 1);

    return factors;
}

// ===============================================
// 4. Тест Миллера на основе теоремы Сэлфриджа
// ===============================================
bool miller_test(long long n, int t) {
    if (n < 2) return false;
    if (n == 2 || n == 3) return true;
    if (n % 2 == 0) return false;

    long long m = n - 1;
    auto factors = factorize(m);

    random_device rd;
    mt19937 gen(rd());
    uniform_int_distribution<long long> dis(2, n - 2);

    unordered_set<long long> used;
    vector<long long> a_list;

    while ((int)a_list.size() < t) {
        long long a = dis(gen);
        if (used.insert(a).second)
            a_list.push_back(a);
    }

    for (long long a : a_list) {
        if (mod_pow(a, m, n) != 1)
            return false;
    }

    for (auto [q, _] : factors) {
        bool all_ones = true;
        for (long long a : a_list) {
            if (mod_pow(a, m / q, n) != 1) {
                all_ones = false;
                break;
            }
        }
        if (all_ones)
            return false;
    }

    return true;
}

// ===============================================
// 5. Генерация числа m для n = 2m + 1
// ===============================================
long long generate_m(const vector<int>& primes, int target_bits, mt19937& gen) {
    uniform_int_distribution<int> prime_dist(0, primes.size() - 1);
    uniform_int_distribution<int> exp_dist(1, 3);

    long long m = 1;
    int current_bits = 0;

    while (current_bits < target_bits - 1) {
        int q = primes[prime_dist(gen)];
        int alpha = exp_dist(gen);
        long long term = pow(q, alpha);

        if (log2(m) + log2(term) > target_bits - 1) continue;

        m *= term;
        current_bits = log2(m) + 1;
    }

    return m;
}

// ===============================================
// 6. Генерация простого числа n = 2m + 1
// ===============================================
long long generate_prime(const vector<int>& primes, int digits, int t) {
    int min_val = pow(10, digits - 1);
    int max_val = pow(10, digits) - 1;
    int target_bits = ceil(log2(max_val));

    random_device rd;
    mt19937 gen(rd());

    while (true) {
        long long m = generate_m(primes, target_bits - 1, gen);
        long long n = 2 * m + 1;

        if (n < min_val || n > max_val)
            continue;

        if (miller_test(n, t))
            return n;
    }
}

// ===============================================
// 7. Тест Миллера-Рабина для финальной проверки
// ===============================================
bool miller_rabin(long long n, int iterations = 5) {
    if (n < 2) return false;
    if (n == 2 || n == 3) return true;
    if (n % 2 == 0) return false;

    long long d = n - 1;
    int r = 0;

    while (d % 2 == 0) {
        d /= 2;
        ++r;
    }

    random_device rd;
    mt19937 gen(rd());
    uniform_int_distribution<long long> dis(2, n - 2);

    for (int i = 0; i < iterations; ++i) {
        long long a = dis(gen);
        long long x = mod_pow(a, d, n);

        if (x == 1 || x == n - 1)
            continue;

        bool continue_outer = false;
        for (int j = 0; j < r - 1; ++j) {
            x = mod_pow(x, 2, n);
            if (x == n - 1) {
                continue_outer = true;
                break;
            }
        }

        if (continue_outer)
            continue;

        return false;
    }

    return true;
}

// ===============================================
// 8. Главная функция с требуемым форматом вывода
// ===============================================
int main() {
    vector<int> primes = sieve(500);
    const int digits = 3;
    const int t = 5;
    const int count = 10;
    int k = 0;

    // Векторы для хранения результатов
    vector<long long> numbers;
    vector<string> results;

    // Генерация и проверка чисел
    for (int i = 0; i < count; ++i) {
        long long p;
        do {
            p = generate_prime(primes, digits, t);
        } while (p < 100 || p > 999);

        bool is_prime = miller_rabin(p, 5);
        numbers.push_back(p);
        results.push_back(is_prime ? "+" : "-");

        if (!is_prime) ++k;
    }

    // Вывод заголовков таблицы
    cout << left << setw(5) << "№";
    for (int i = 1; i <= count; ++i) {
        cout << setw(5) << i;
    }
    cout << "\n-------------------------------------------------\n";

    // Вывод чисел
    cout << setw(5) << "P";
    for (auto num : numbers) {
        cout << setw(5) << num;
    }
    cout << "\n";

    // Вывод результатов
    cout << setw(5) << "Test";
    for (auto res : results) {
        cout << setw(5) << res;
    }
    cout << "\n-------------------------------------------------\n";

    // Вывод значения K
    cout << "K = " << k << endl;

    return 0;
}
