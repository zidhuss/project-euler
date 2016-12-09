day_of_year = 365 % 7
sundays = 0
months = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
for year in range(1901, 2001):
    for i in range(12):
        number_of_days = months[i]
        if year % 4 == 0 and i == 1:
            number_of_days = 29
        day = day_of_year % 7
        day_of_year += number_of_days
        if day == 6:
            sundays += 1

print(sundays)
