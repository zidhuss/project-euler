#include <stdio.h>
#include <string.h>

int main() {
    FILE *fp = fopen("resources/cipher.txt", "r");
    char txt[5000] = {0};
    fgets(txt, 5000, fp);

    char* key = "god";
    char *c = txt;
    int sum = 0;

    *(txt + strlen(txt) - 1) = 0; // Trim new line
    int i = 0;
    while (*c) {
        if (*c == ',' || *c == '\n') c++;
        int x = 0;
        int n = 0;
        sscanf(c, "%d%n,", &x, &n);
        c += n;
        sum += x ^ (key[i%3]);
        i++;
    }
    printf("%d\n", sum);
}
