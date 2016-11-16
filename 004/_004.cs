using System;

public class _004 {
    public static void Main() {
        int last = 999;
        int first = 100;
        int largest = 0;
        for (int a = last; a >= first; a--)
            for (int b = last; b >= first; b--) {
                int calculated = a*b;
                if (calculated > largest && IsPalindrome(calculated)) {
                    largest = calculated;
                }
            }
        Console.WriteLine(largest);
    }

    public static bool IsPalindrome(int number) {
        int n = number;
        int reverse = 0;
        while (n > 0) {
            reverse = reverse * 10 + n % 10;
            n /= 10;
        }
        return number == reverse;
    }
}
