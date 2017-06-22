#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <gmp.h>

int main() {
    mpz_t sum;
    mpz_t x;
    mpz_init(sum);
    mpz_init(x);
    char *res = calloc(3001, sizeof(char));

    for (int i = 1; i <= 1000; i++) {
        mpz_ui_pow_ui(x, i, i);
        mpz_add(sum, sum, x);
    }

    mpz_get_str(res, 10, sum);
    char* last = res + (strlen(res) - 10);
    printf("%s\n", last);
}
