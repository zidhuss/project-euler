#include <stdio.h>

int squareDigits(int n) {
    int res = 0;
    while (n > 0) {
        int d = n % 10;
        res += d * d;
        n /= 10;
    }
    return res;
}

int main() {
    int count = 0;
    for (int i = 2; i < 1e7; i++) {
        int x = i;
        while (x != 1) {
            if (x == 89) {
                count++;
                break;
            }
            x = squareDigits(x);
        }
    }
    printf("%d\n", count);
}
