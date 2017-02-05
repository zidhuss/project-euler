using System;

public class _001 {
    public static void Main() {
        Console.WriteLine(Sum(1000, 3) + Sum(1000, 5) - Sum(1000, 15));
    }

    public static int Sum(int n, int k) {
        return k * ((((n - 1) / k) * (((n - 1) / k) + 1)) / 2);
    }
}
