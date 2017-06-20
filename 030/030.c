#include <stdio.h>

int intPow(int n, int e) {
    int s = 1;
    for (int i = 0; i < e; i++)
        s *= n;
    return s;
}

int main() {
    int sum = 0;
    for (int i = 355000; i > 2; i--) {
        int j = i;
        int raisedSum = 0;

        while (j > 0) {
            raisedSum += intPow((j % 10), 5);
            j /= 10;
        }

        if (i == raisedSum)
            sum += i;
    }
    printf("%d\n", sum);
}
