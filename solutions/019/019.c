#include <stdio.h>

int main() {
    int dayOfYear = 365 % 7;
    int sundays = 0;
    int months[12] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
    for (int year = 1901; year < 2001; year++) {
        for (int i = 0; i < 12; i++) {
            int numberOfDays = months[i];
            if (year % 4 == 0 && i == 1)
                numberOfDays = 29;
            int day = dayOfYear % 7;
            dayOfYear += numberOfDays;
            if (day == 6)
                sundays++;
        }
    }
    printf("%d\n", sundays);
}
