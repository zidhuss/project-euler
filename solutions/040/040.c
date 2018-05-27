#include <stdio.h>
#include <string.h>

#define LIMIT 1000000

int main() {
    char digits[LIMIT + 1] = {0};

    char temp[30];
    char *ptr = digits;

    for (int i = 0, x = 1; i < LIMIT; x++) {
        sprintf(temp, "%d", x);
        int len = strlen(temp);
        strncpy(ptr, temp, len);
        i += len;
        ptr += len;
    }

    int sum = 1;
    for (int i = 1; i <= LIMIT; i *= 10) {
        sum *= digits[i-1] - 48;
    }

    printf("%d\n", sum);
}
