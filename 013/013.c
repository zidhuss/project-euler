#include <stdio.h>
#include <string.h>

int main() {
    FILE *fp = fopen("./resources/numbers.txt", "r");
    char line[60] = {0};
    unsigned long sum = 0;
    unsigned long n = 0;

    while (fgets(line, 80, fp)) {
        *(line + 11) = 0;
        sscanf(line, "%lu", &n);
        sum += n;
    }

    while (sum >= 1e10)
        sum /= 10;

    printf("%lu\n", sum);
}
