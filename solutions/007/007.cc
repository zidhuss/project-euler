#include <iostream>

int main() {
    int n = 10001, lenPrimes = 1;
    int primes[n];
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
    std::cout << primes[n-1] << std::endl;
}
