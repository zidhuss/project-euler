public class _003 {
    public static void main(String[] args) {
        System.out.println(largestPrimeFactor(600851475143L));
    }

    public static long largestPrimeFactor(long number) {
        long factor = 1;
        for (int i = 3; i * i <= number; i += 2) {
            if (number == 1)
                return factor;
            else if (number % i != 0)
                continue;
            else {
                factor = i;
                while (number % i == 0)
                    number /= i;
            }
        }
        return number;
    }
}
