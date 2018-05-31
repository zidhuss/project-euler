#include <stdio.h>

int gcd(int a, int b) {
    int c;
    while (b != 0) {
        c = b;
        b = a % b;
        a = c;
    }
    return a;
}

int main() {
    int num = 1;
    int den = 1;

    for (int i = 1; i < 10; i++) {
        for (int d = 1; d < i; d++) {
            for (int n = 1; n < d; n++) {
                int a = (n*10) + i;
                int b = (i*10) + d;
                if (((n*1.0)/d) == ((a*1.0)/b)) {
                    num *= a;
                    den *= b;
                }
            }
        }
    }

    printf("%d\n", den / gcd(num, den));
}
