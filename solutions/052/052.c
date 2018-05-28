#include <stdio.h>

int hasSameDigits(int a, int b) {
    int x[10] = {0};
    int y[10] = {0};

    while (a > 0) {
        x[a%10]++;
        a /= 10;
    }

    while (b > 0) {
        y[b%10]++;
        b /= 10;
    }

    for (int i = 0; i < 10; i++) {
        if (x[i] != y[i])
            return 0;
    }
    return 1;
}

int hasSameDigitsRange(int n) {
    for (int i = 2; i < 7; i++) {
        if (!hasSameDigits(n, n*i))
            return 0;
    }
    return 1;
}

int main() {
    int i = 1;
    while (!hasSameDigitsRange(i))
        i++;
    printf("%d\n", i);
}
