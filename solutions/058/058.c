#include <stdio.h>

int isPrime(int n) {
    if (n <= 1)
        return 0;
    else if (n <= 3)
        return 1;
    else if ((n%2 == 0) || (n%3 == 0))
        return 0;

    int i = 5;
    while (i*i <= n) {
        if ((n%i == 0) || (n%(i+2) == 0))
            return 0;
        i += 6;
    }
    return 1;
}

int main() {
    int primeCount = 0;
    int x = 3;

    while (1) {
        for (int i = 3; i > 0; i--) {
            int n = x*x - i*(x-1);
            if (isPrime(n))
                primeCount++;
        }

        if (0.1 > ((double)primeCount) / (2*(x-1)+1))
            break;
        x += 2;
    }

    printf("%d\n", x);
}
