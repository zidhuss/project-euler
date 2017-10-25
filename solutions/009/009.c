#include <stdio.h>
#include <math.h>

int main() {
    int s = 1000;
    int r = (int) sqrt(s);
    for (int m = r; m > 1; m--)
        for (int n = m; n > 2; n--) {
            int a = m * m - n * n;
            int b = 2 * (m * n);
            int c = m * m + n * n;
            if (a + b + c == s) {
                printf("%d\n", a * b * c);
                return 0;
            }
        }
}
