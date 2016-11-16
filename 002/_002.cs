using System;

public class _002 {
    public static void Main() {
        int[] fibo = new int[36];
        fibo[0] = 0;
        fibo[1] = 1;
        for (int i = 2; i < 34; i++) {
            fibo[i] = fibo[i - 2] + fibo[i - 1];
        }
        int s = 0;
        for (int i = 0; i < 36; i++) {
            if (fibo[i] % 2 == 0)
                s += fibo[i];
        }
        Console.WriteLine(s);
    }
}
