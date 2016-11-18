public class _005 {
    public static void main(String[] args) {
        int result = 1;
        for (int i = 1; i < 20; i++)
            result = lcm(result, i);
        System.out.println(result);
    }

    public static int lcm(int a, int b) {
        return a * b / gcd(a, b);
    }

    public static int gcd(int a, int b) {
        while (b != 0) {
            int c = b;
            b = a % b;
            a = c;
        }
        return a;
    }
}
