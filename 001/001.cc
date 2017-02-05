#include <iostream>

int sum(int n, int k) {
	return k * (((n-1) / k) * (((n-1) / k) + 1)) / 2;
}

int main() {
    int solution = sum(1000, 3) + sum(1000, 5) - sum(1000, 15);
    std::cout << solution << std::endl;
}
