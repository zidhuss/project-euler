#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <gmp.h>

int main() {
    mpz_t fn;
    mpz_init(fn);
    char* str = calloc(1500, sizeof(char));

    int n = 100;
    while (strlen(str) < 1000) {
        mpz_fib_ui(fn, n);
        mpz_get_str(str, 10, fn);
        n++;
    }
    printf("%d\n", n-1);
}
