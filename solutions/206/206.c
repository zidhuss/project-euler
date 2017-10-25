#include <stdio.h>

int isValid(unsigned long long n) {
    if (n < 1e18 || n%10 != 0)
        return 0;

    unsigned long long i = 9;
    while (n > 0) {
        n /= 100;
        if (n%10 != i)
            return 0;
        i--;
    }

    return 1;
}

int main() {
    for (unsigned long long i = 1e9; i < 1e19; i += 2)
        if (i%3 == 0 || i%7 == 0)
            if (isValid(i*i)) {
                printf("%llu\n", i);
                break;
            }
}
