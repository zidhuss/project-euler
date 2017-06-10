#include <stdio.h>


int main() {
    int numbers[1000];
    FILE *fp;
    fp = fopen("./resources/digits.txt", "r");

    int i = 0;
    for (char c = fgetc(fp); c != EOF; c = fgetc(fp)) {
        if (c != '\n') {
            numbers[i] = c - 48;
            i++;
        }
    }

    int size = 13;
    long largest = 0;

    for (int j = size; j < 1000-size; j++) {
        long product = 1;
        for (int k = 0; k < size; k++) {
            product *= numbers[k+j];
        }
        if (product > largest) largest = product;
    }

    printf("%li\n", largest);
}
