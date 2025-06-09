#include <iostream>
#include <fstream>
#include <cmath>
#include <iomanip>

using namespace std;

int main() {
    double Xstart = -6.0, Xend = 8.0, dx = 0.2;
    double x, y;

    ofstream fout("output_table.txt");
    fout << fixed << setprecision(2);
    cout << fixed << setprecision(2);

    fout << "  x\t|\ty\n";
    fout << "-------------------\n";
    cout << "  x\t|\ty\n";
    cout << "-------------------\n";

    for (x = Xstart; x <= Xend; x += dx) {
        if (x <= -2.0) {
            y = x + 3; // участок 1
        } else if (x > -2.0 && x < 0.0) {
            y = pow(3, x); // участок 2
        } else if (x >= 0.0 && x <= 6.0) {
            double R = 3.0;
            double centerX = 3.0, centerY = 3.0;
            double expr = R * R - (x - centerX) * (x - centerX);
            if (expr >= 0) {
                y = centerY + sqrt(expr); // Верхняя полусфера (исправлено)
            } else {
                y = NAN;
            }
        } else if (x > 6.0) {
            y = -1.5 * x + 12; // участок 4
        }

        if (!isnan(y)) {
            fout << setw(5) << x << "\t|\t" << y << endl;
            cout << setw(5) << x << "\t|\t" << y << endl;
        } else {
            fout << setw(5) << x << "\t|\tN/A" << endl;
            cout << setw(5) << x << "\t|\tN/A" << endl;
        }
    }

    fout.close();
    cout << "\nРезультаты также записаны в файл output_table.txt\n";

    return 0;
}
