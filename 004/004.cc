#include <iostream>

bool is_palindrome(int orig) {
    int rev = 0, n = orig;
    while (n  > 0) {
        rev = rev * 10 + n % 10;
        n /= 10;
    }
    return orig == rev;
}

int main() {
    int largest = 0, last = 999, first = 100;
    for (int a = last; a >= first; a--)
        for (int b = last; b >= first; b--) {
            int calculated = a * b;
            if (calculated > largest && is_palindrome(calculated))
                largest = calculated;
        }
    std::cout << largest << std::endl;
}
