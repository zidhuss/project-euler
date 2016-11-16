using System;

public class _001 {
    public static void Main() {
        int s = 0;
        for (int i = 1; i < 1000; i++) {
            if (i % 3 == 0 || i % 5 == 0)
                s += i;
        }
        Console.WriteLine(s);
    }
}
