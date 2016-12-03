#include <stdio.h>
#include <stdlib.h>

int main() {
    int n = 10001, lenPrimes = 1;
    int *primes = malloc(n * sizeof(int));
    primes[0] = 2;

    for (int x = 3; lenPrimes < n+1; x += 2) {
        for (int i = 0; i < lenPrimes; i++)
            if (x % primes[i] == 0)
                break;
            else if (primes[i]*primes[i] > x) {
                primes[lenPrimes++] = x;
                break;
            }
    }
    printf("%d\n", primes[n-1]);
}
