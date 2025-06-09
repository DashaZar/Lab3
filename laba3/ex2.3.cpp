#include <iostream>
#include <vector>
#include <cmath>
#include <random>
#include <iomanip>

using namespace std;

// Функция реализует алгоритм "Решето Эратосфена" для поиска всех простых чисел до заданного предела
vector<int> sieve(int limit) {
    vector<bool> is_prime(limit + 1, true); // Массив флагов для отметки простых чисел
    vector<int> primes; // Вектор для хранения найденных простых чисел
    
    is_prime[0] = is_prime[1] = false; // 0 и 1 не являются простыми
    
    // Основной цикл алгоритма: последовательно отсеиваем составные числа
    for (int i = 2; i <= limit; ++i) {
        if (is_prime[i]) { // Если i простое
            primes.push_back(i); // Добавляем в результат
            // Помечаем все кратные i как составные
            for (int j = i * 2; j <= limit; j += i)
                is_prime[j] = false;
        }
    }
    return primes;
}

// Функция для быстрого возведения в степень по модулю (оптимизированный алгоритм)
long long mod_pow(long long base, long long exp, long long mod) {
    long long result = 1;
    base %= mod; // Уменьшаем base для работы с большими числами
    
    // Алгоритм основан на двоичном разложении показателя степени
    while (exp > 0) {
        if (exp % 2) // Если текущий бит степени установлен
            result = (result * base) % mod;
        base = (base * base) % mod; // Возводим base в квадрат
        exp /= 2; // Переходим к следующему биту
    }
    return result;
}

// Вероятностный тест простоты Миллера-Рабина (k - количество проверок)
bool is_probably_prime(long long n, int k = 10) {
    // Обработка тривиальных случаев
    if (n < 2 || n % 2 == 0) 
        return n == 2;
    
    // Представим n-1 в виде d*2^s
    long long d = n - 1;
    int s = 0;
    while (d % 2 == 0) {
        d /= 2;
        ++s;
    }
    
    // Генератор случайных чисел
    mt19937 gen(random_device{}());
    uniform_int_distribution<long long> dist(2, n - 2);
    
    // Проводим k итераций теста
    for (int i = 0; i < k; ++i) {
        long long a = dist(gen); // Выбираем случайное основание
        long long x = mod_pow(a, d, n);
        
        if (x == 1 || x == n - 1) 
            continue;
        
        // Повторяем s-1 раз проверку
        bool passed = false;
        for (int r = 1; r < s; ++r) {
            x = mod_pow(x, 2, n);
            if (x == n - 1) {
                passed = true;
                break;
            }
        }
        if (!passed) 
            return false;
    }
    return true; // Число вероятно простое
}

// Генерация простого числа p по алгоритму ГОСТ Р 34.10 (трехзначные числа)
bool generate_gost_prime(int q, long long& p_out) {
    // Инициализация генератора случайных чисел
    mt19937 gen(random_device{}());
    uniform_real_distribution<> dist(0.0, 1.0);
    
    // Вычисление параметра N
    double xi = dist(gen); // Случайное число от 0 до 1
    int N = static_cast<int>(100.0 / (2 * q) + xi * (100.0 / q));
    if (N % 2 == 1) N++; // Гарантируем четность N
    
    // Поиск подходящего p вида (N+u)*q + 1
    for (int u = 0; ; u += 2) { // u увеличивается на 2 для сохранения четности
        long long p = (N + u) * q + 1;
        if (p > 999) break; // Ограничение на трехзначные числа
        if (p < 100) continue; // Пропускаем двухзначные
        
        // Проверка по теореме Диемитко с основанием a=2
        if (mod_pow(2, p - 1, p) == 1 && mod_pow(2, N + u, p) != 1) {
            p_out = p;
            return true;
        }
    }
    return false; // Подходящее p не найдено
}

// Функция для вывода результатов в виде форматированной таблицы
void print_table(const vector<long long>& p_list, const vector<char>& results, int k) {
    // Шапка таблицы
    cout << "+----+---------+-----------------------------+-----+\n";
    cout << "| №  |    p    | Результат вероятностным тестом |  K  |\n";
    cout << "+----+---------+-----------------------------+-----+\n";
    
    // Строки с данными
    for (size_t i = 0; i < p_list.size(); ++i) {
        cout << "| " << setw(2) << i + 1 << " | "
             << setw(7) << p_list[i] << " | "
             << setw(27) << results[i] << " | "
             << setw(3) << (i == 0 ? to_string(k) : "") << " |\n";
    }
    
    // Нижняя граница таблицы
    cout << "+----+---------+-----------------------------+-----+\n";
}

int main() {
    // Генерация списка простых чисел до 500
    vector<int> primes = sieve(500);
    vector<long long> p_list; // Список сгенерированных чисел p
    vector<char> test_results; // Результаты тестирования ('+' или '-')
    int k = 0; // Счетчик ложноположительных результатов
    
    int count = 0; // Счетчик успешно сгенерированных чисел p
    for (size_t i = 0; i < primes.size() && count < 10; ++i) {
        int q = primes[i];
        if (q >= 100) continue; // Используем только малые q (двухзначные)
        
        long long p;
        if (generate_gost_prime(q, p)) {
            // Проверка числа p тестом Миллера-Рабина
            bool probable = is_probably_prime(p);
            
            // Ложноположительный случай: p соответствует ГОСТ, но тест не пройден
            if (!probable) k++;
            
            // Сохраняем результаты
            p_list.push_back(p);
            test_results.push_back(probable ? '+' : '-');
            count++;
        }
    }
    
    // Вывод результатов
    print_table(p_list, test_results, k);
    return 0;
}
