using System;

public class _015 {
    public static void Main(String[] args) {
        Console.WriteLine(BinomialCoefficient(40, 20));
    }

    public static long BinomialCoefficient(long n, long k) {
        if (k < 0 || k > n)
            return 0;
        if (k == 0 || k == n)
            return 1;
        if (n-k < k)
            k = n - k;
        long c = 1;
        for (long i = 0; i < k; i++)
            c = c * (n - i) / (i + 1);
        return c;
    }
}
