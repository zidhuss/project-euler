public class _001 {
    public static void main(String[] args) {
        System.out.println(sum(1000, 3) + sum(1000, 5) - sum(1000, 15));
    }

    public static int sum(int n, int k) {
        return k * ((((n - 1) / k) * (((n - 1) / k) + 1)) / 2);
    }
}
