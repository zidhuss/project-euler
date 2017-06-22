#include <stdio.h>
#include <gmp.h>

int main() {
    int largest = 0;
    mpz_t z;
    mpz_init(z);
    char res[202];

    for (int a = 1; a <= 100; a++) {
        for (int b = 1; b <= 100; b++) {
            mpz_ui_pow_ui(z, a, b);
            mpz_get_str(res, 10, z);

            int sum = 0;
            for (char* c = res; *c; c++)
                sum += *c - '0';
            if (sum > largest)
                largest = sum;
        }
    }
    printf("%d\n", largest);
}
