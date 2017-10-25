using System;

public class _009 {
    public static void Main(string[] args) {
        int s = 1000;
        int r = (int) Math.Sqrt(s);
        for (int m = r; m > 1; m--)
            for (int n = m; n > 2; n--) {
                int a = m * m - n * n;
                int b = 2 * (m * n);
                int c = m * m + n * n;
                if (a + b + c == s) {
                    Console.WriteLine(a * b * c);
                    return;
                }
            }
    }
}
