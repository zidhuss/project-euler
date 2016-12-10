#include <stdio.h>

int a(int n) {
    return n*(n+2)/4 + (n%4)/3 + 1;
}

int main() {
    int n = 1001;
    int sum = 0;
    for (int i = 1; a(i) <= n*n; i++)
        sum += a(i);
    printf("%d\n", sum);
}
