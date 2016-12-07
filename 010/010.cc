#include <iostream>
#include <cmath>

int main() {
    int limit = 2000000;
    int sievebound = (limit - 1) / 2;
    int crosslimit = (int) (sqrt(limit) - 1) / 2;
    int *sieve = new int[sievebound];

    for (int i = 1; i < crosslimit; i++)
        if (sieve[i] == 0)
            for (int j = 2*i*(i+1); j < sievebound; j += 2*i+1)
                sieve[j] = 1;

    long sum = 2;
    for (int i = 1; i < sievebound; i++)
        if (sieve[i] == 0)
            sum += 2*i+1;

    std::cout << sum << std::endl;
}
