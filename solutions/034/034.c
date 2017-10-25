#include <stdio.h>

int factorial(int n) {
    int s = 1;
    for (int i = n; i > 0; i--)
        s *= i;
    return s;
}

int main() {
	int sum = 0;
	for (int i = 4e5; i > 2; i--) {
		int j = i;
		int factorialSum = 0;

		while (j > 0) {
			factorialSum += factorial(j % 10);
			j /= 10;
		}

		if (i == factorialSum)
			sum += i;
	}
    printf("%d\n", sum);
}
