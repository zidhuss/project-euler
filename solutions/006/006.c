#include <stdio.h>

int main() {
    int sumSquare = 0, squareSum = 0;
    for (int i = 0; i <= 100; i++) {
        squareSum += i * i;
        sumSquare += i;
    }
    printf("%d\n", sumSquare * sumSquare - squareSum);
}
