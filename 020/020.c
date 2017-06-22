#include <stdio.h>
#include <gmp.h>

int main() {
    mpz_t z;
    mpz_init(z);
    mpz_fac_ui(z, 100);
    char res[200];
    mpz_get_str(res, 10, z);

    int sum = 0;
    for (char* c = res; *c; c++) {
        sum += *c - '0';
    }
    printf("%d\n", sum);
}
