#include <stdio.h>

int allOdd(int n) {
    while (n > 0) {
        int x = n % 10;
        if (x % 2 == 0)
            return 0;
        n /= 10;
    }
    return 1;
}

int reverse(int n) {
    int res = 0;
    while (n > 0) {
        int x = n % 10;
        res *= 10;
        res += x;
        n /= 10;
    }
    return res;
}

int main() {
    int count = 0;
    for (int i = 11; i < 1e8; i++) {
        if (i % 10 == 0) continue;
        if (allOdd(i + reverse(i))) count++;
    }
    printf("%d\n", count);
}
