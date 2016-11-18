public class _006 {
    public static void main(String[] args) {
        int sumSquare = 0, squareSum = 0;
        for (int i = 0; i <= 100; i++) {
            squareSum += i * i;
            sumSquare += i;
        }
        System.out.println(sumSquare * sumSquare - squareSum);
    }
}
