#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>


int nameScore(char *str) {
    int score = 0;
    while (*str) {
        score += *str - 64;
        str++;
    }
    return score;
}

int main() {
    FILE *fp = fopen("./resources/names.txt", "r");
    char names[5170][30];

    int index = 0;
    char *ptr = names[index];
    int in_name = 0;
    int c;

    while ((c = fgetc(fp)) != EOF) {
        if (c == ',') continue;

        if (c == '"') {
            if (in_name) {
                // finished name
                *ptr = 0;
                ptr = names[++index];
            }
            in_name = !in_name;
            continue;
        }

        if (isalpha(c))
            *ptr++ = c;
    }

    qsort(names, index, 30, strcmp);

    int sum = 0;
    for (int i = 0; i < index; i++) {
        /* printf("%s\n", names[i]); */
        sum += nameScore(names[i]) * (i + 1);
    }

    printf("%d\n", sum);

    fclose(fp);
}
