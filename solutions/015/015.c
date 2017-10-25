#include <stdio.h>

unsigned long binomialCoefficient(unsigned long n, unsigned long k) {
    if (k > n)
        return 0;
    if (k == n)
        return 1;
    if (n-k < k)
        k = n - k;
    unsigned long c = 1;
    for (unsigned long i = 0; i < k; i++)
        c = c * (n - i) / (i + 1);
    return c;
}

int main() {
    printf("%lu\n", binomialCoefficient(40, 20));
}
