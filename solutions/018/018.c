#include <stdio.h>
#include <stdlib.h>

#define NO_OF_ROWS 15


int maximumPathSum(int tree[NO_OF_ROWS][NO_OF_ROWS]) {
    for (int i = NO_OF_ROWS - 2; i >= 0; i--) {
        for (int j = 0; j <= i; j++) {
            int x = tree[i+1][j];
            int y = tree[i+1][j+1];
            tree[i][j] += (x > y) ? x : y;
        }
    }
    return tree[0][0];
}


int main() {
    FILE *fp = fopen("./resources/tree.txt", "r");

    int tree[NO_OF_ROWS][NO_OF_ROWS];

    int row = 0;
    int col = 0;

    char tmp[3];
    char *ptr = tmp;

    int c;

    while ((c = fgetc(fp)) != EOF) {
        if (c == ' ') {
            tree[row][col++] = atoi(tmp);
            ptr = tmp;
        } else if (c == '\n') {
            tree[row][col++] = atoi(tmp);
            ptr = tmp;
            row++;
            col = 0;
        } else {
            *ptr++ = c;
        }
    }

    printf("%d\n", maximumPathSum(tree));
}
