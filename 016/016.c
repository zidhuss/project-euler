#include <stdio.h>
#include <gmp.h>

int main() {
    mpz_t z;
    mpz_init(z);
    mpz_ui_pow_ui(z, 2, 1000);
    char res[400];
    mpz_get_str(res, 10, z);

    int sum = 0;
    for (char* c = res; *c; c++)
        sum += *c - '0';

    printf("%d\n", sum);
}
