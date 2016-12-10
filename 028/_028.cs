using System;

public class _028 {
    public static void Main(string[] args) {
        int n = 1001;
        int sum = 0;
        for (int i = 1; a(i) <= n*n; i++)
            sum += a(i);
        Console.WriteLine(sum);
    }

    private static int a(int n) {
        return n*(n+2)/4 + (n%4)/3 + 1;
    }
}
