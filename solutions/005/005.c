#include <stdio.h>

int gcd(int a, int b) {
    int c;
    while (b != 0) {
        c = b;
        b = a % b;
        a = c;
    }
    return a;
}

int lcm(int a, int b) {
    return a * b / gcd(a, b);
}

int main() {
    int result = 1;
    for (int i = 1; i < 20; i++)
        result = lcm(result, i);
    printf("%d\n", result);
}
