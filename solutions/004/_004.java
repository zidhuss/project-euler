public class _004 {
    public static void main(String[] args) {
        System.out.println(largestPalindromeProduct(999, 100));
    }

    public static boolean isPalindrome(int number) {
        int n = number;
        int reverse = 0;
        while (n > 0) {
            reverse = reverse * 10 + n % 10;
            n /= 10;
        }
        return number == reverse;
    }

    public static int largestPalindromeProduct(int from, int to) {
        int largest = 0;
        for (int a = from; a >= to; a--)
            for (int b = from; b >= to; b--) {
                int calculated = a*b;
                if (calculated > largest && isPalindrome(calculated))
                    largest = calculated;
            }
        return largest;
    }
}
