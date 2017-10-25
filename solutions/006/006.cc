#include <iostream>

int main() {
    int sumSquare = 0, squareSum = 0;
    for (int i = 0; i <= 100; i++) {
        squareSum += i * i;
        sumSquare += i;
    }
    std::cout << sumSquare * sumSquare - squareSum << std::endl;
}
