using System;

public class _006 {
    public static void Main(String[] args) {
        int sumSquare = 0, squareSum = 0;
        for (int i = 0; i <= 100; i++) {
            squareSum += i * i;
            sumSquare += i;
        }
        Console.WriteLine(sumSquare * sumSquare - squareSum);
    }
}
