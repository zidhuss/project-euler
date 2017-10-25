#include <iostream>

int main() {
    int fibo[36] = {0, 1};
    for (int i = 2; i < 34; i++) {
        fibo[i] = fibo[i - 2] + fibo[i - 1];
    }
    int s = 0;
    for (int i = 0; i < 36; i++) {
        if (fibo[i] % 2 == 0)
            s += fibo[i];
    }
    std::cout << s << std::endl;
}
