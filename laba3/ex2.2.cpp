#include <iostream>
#include <vector>
#include <tuple>
#include <set>
#include <cmath>
#include <random>
#include <iomanip>
#include <algorithm>

using namespace std;

// Функция возведения в степень по модулю
int powMod(int base, int degree, int mod) {
    int result = 1;
    base %= mod;
    while (degree > 0) {
        if (degree % 2 == 1)
            result = (1LL * result * base) % mod;
        base = (1LL * base * base) % mod;
        degree /= 2;
    }
    return result;
}

// Решето Эратосфена
vector<int> sieveEratos(int N) {
    vector<bool> isPrime(N + 1, true);
    vector<int> primes;
    isPrime[0] = isPrime[1] = false;
    for (int i = 2; i <= N; ++i) {
        if (isPrime[i]) {
            primes.push_back(i);
            for (int j = i * 2; j <= N; j += i)
                isPrime[j] = false;
        }
    }
    return primes;
}

// Миллер — Рабин
bool rabin(int num, int k) {
    if (num < 2) return false;
    if (num == 2 || num == 3) return true;
    if (num % 2 == 0) return false;

    int s = 0;
    int d = num - 1;
    while (d % 2 == 0) {
        d /= 2;
        s++;
    }

    for (int i = 0; i < k; i++) {
        int a = 2 + rand() % (num - 3);
        int x = powMod(a, d, num);
        if (x == 1 || x == num - 1)
            continue;

        bool passed = false;
        for (int r = 1; r < s; ++r) {
            x = powMod(x, 2, num);
            if (x == num - 1) {
                passed = true;
                break;
            }
        }
        if (!passed)
            return false;
    }

    return true;
}

// Случайное число в диапазоне
int randDist(int min, int max) {
    random_device rd;
    mt19937 gen(rd());
    uniform_int_distribution<> dist(min, max);
    return dist(gen);
}

// Тест Поклингтона
bool poklington(int n, int t, const vector<int>& qList) {
    set<int> aSet;
    while (aSet.size() < t)
        aSet.insert(randDist(2, n - 2));

    for (int a : aSet) {
        if (powMod(a, n - 1, n) != 1)
            return false;
    }

    for (int a : aSet) {
        bool conditionMet = true;
        for (int q : qList) {
            if (powMod(a, (n - 1) / q, n) == 1) {
                conditionMet = false;
                break;
            }
        }
        if (conditionMet)
            return true;
    }

    return false;
}

// Формирование F и p
tuple<int, vector<int>> calcN(const vector<int>& primes, int bit) {
    int minBitF = bit / 2 + 1;
    int maxBitF = bit / 2 + 2;
    int minF = 1 << minBitF;
    int maxF = 1 << maxBitF;

    int F = 1;
    vector<int> qList;
    set<int> usedPrimes;

    while (F < minF) {
        int prime = primes[randDist(0, primes.size() - 1)];
        if (usedPrimes.count(prime)) continue;

        int alpha = randDist(1, 3);
        int powValue = pow(prime, alpha);

        if (F * powValue >= maxF)
            break;

        F *= powValue;
        qList.push_back(prime);
        usedPrimes.insert(prime);
    }

    int R = randDist(2, 10) * 2; // чётное
    int p = R * F + 1;

    return {p, qList};
}

int main() {
    const int bit = 10;
    const int t = 10;

    vector<int> primes = sieveEratos(500);
    vector<int> results;
    vector<string> test_results;

    while (results.size() < 10) {
        auto [p, qList] = calcN(primes, bit);
        if (p < 100 || p > 999) continue;  // Только трёхзначные

        if (poklington(p, t, qList)) {
            results.push_back(p);
            test_results.push_back(rabin(p, 3) ? "+" : "-");
        }
    }

    // Вертикальный вывод
    cout << left;
    cout << setw(5) << "№" << setw(10) << "P" << setw(6) << "Test" << "\n";
    cout << "---------------------------\n";

    for (int i = 0; i < 10; ++i) {
        cout << setw(5) << (i + 1)
             << setw(10) << results[i]
             << setw(6) << test_results[i] << "\n";
    }

    int k = count(test_results.begin(), test_results.end(), "-");
    cout << "---------------------------\n";
    cout << "Количество не прошедших тест: " << k << "\n";

    return 0;
}
