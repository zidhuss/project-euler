#include <stdio.h>

int sumDivisors(int n) {
    int sum = 0;
    int i = 2;
    for (; i*i < n; i++)
        if ((n%i) == 0)
            sum += i + (n / i);
    if (i*i == n) sum += i;
    return sum + 1;
}

int main() {
    int amicableSum = 0;
    for (int a = 2; a < 10000; a++) {
        int b = sumDivisors(a);
        if ((b > a) && (sumDivisors(b) == a))
                amicableSum += a + b;
    }

    printf("%d\n", amicableSum);
}
