#include <iostream>
#include <cmath>

int main() {
    int s = 1000;
    int r = (int) sqrt(s);
    for (int m = r; m > 1; m--)
        for (int n = m; n > 2; n--) {
            int a = m * m - n * n;
            int b = 2 * (m * n);
            int c = m * m + n * n;
            if (a + b + c == s) {
                std::cout << a * b * c << std::endl;
                return 0;
            }
        }
}
