using System;

public class p003 {
    public static void Main(String[] args) {
        Console.WriteLine(LargestPrimeFactor(600851475143));
    }

    public static long LargestPrimeFactor(long number) {
        long factor = 1;
        for (int i = 2; i * i <= number; i++) {
            if (number == 1)
                return factor;
            else if (number % i != 0)
                continue;
            else {
                factor = i;
                while (number % i == 0)
                    number /= i;
            }
        }
        return number;
    }
}
