using System;

public class _007 {
    public static void Main(string[] args) {
        int[] primes = new int[10001];
        primes[0] = 2;
        int lenPrimes = 1;
        for (int x = 3; lenPrimes < 10001; x += 2)
            for (int i = 0; i < lenPrimes; i++)
                if (x % primes[i] == 0)
                    break;
                else if (primes[i] * primes[i] > x) {
                    primes[lenPrimes++] = x;
                    break;
                }
        Console.WriteLine(primes[primes.Length - 1]);
    }
}
