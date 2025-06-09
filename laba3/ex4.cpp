#include <iostream>
#include <vector>
#include <climits>

using namespace std;

int main() {
    int n, m;
    cin >> n >> m;

    // Проверка входных данных
    if (n < 5 || n > 50000) {
        cerr << "n must be between 5 and 50000" << endl;
        return 1;
    }
    if (m < 4 || m > 100) {
        cerr << "m must be between 4 and 100" << endl;
        return 1;
    }

    vector<int> nums(n);
    for (int i = 0; i < n; ++i) {
        cin >> nums[i];
    }

    // Вычисляем префиксные суммы
    vector<long long> prefix(n + 1, 0);
    for (int i = 1; i <= n; ++i) {
        prefix[i] = prefix[i - 1] + nums[i - 1];
    }

    // Динамическое программирование
    vector<long long> dp(n + 1, 0);
    
    for (int i = n - 1; i >= 0; --i) {
        long long max_score = LLONG_MIN;
        for (int j = 1; j <= m && i + j <= n; ++j) {
            long long current = (prefix[i + j] - prefix[i]) - dp[i + j];
            if (current > max_score) {
                max_score = current;
            }
        }
        dp[i] = max_score;
    }

    // Определяем победителя
    cout << (dp[0] > 0 ? 1 : 0) << endl;

    return 0;
} 
