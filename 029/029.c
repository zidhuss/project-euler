#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <gmp.h>

int main() {
    mpz_t z;
    mpz_init(z);

    char *numberStrings[10000];
    int count = 0;

    char n[500];

    for (unsigned int a = 2; a <= 100; a++)
        for (unsigned int b = 2; b <= 100; b++) {
            int found = 0;
            mpz_ui_pow_ui(z, a, b);
            mpz_get_str(n, 10, z);
            for (int i = 0; i < count; i++) {
                if (strcmp(numberStrings[i], n) == 0) {
                    found = 1;
                    break;
                }
            }
            if (!found) {
                numberStrings[count] = malloc(500);
                strcpy(numberStrings[count], n);
                count++;
            }
        }
    printf("%d\n", count);
}
