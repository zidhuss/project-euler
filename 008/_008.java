import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;

public class _008 {
    public static void main(String[] args) {
        try {
            long[] numbers = new String(Files.readAllBytes(Paths.get("./resources/digits.txt")))
                .replaceAll("\n", "").chars().mapToLong(x -> x - '0').toArray();
            System.out.println(adjacentProduct(numbers, 13));
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private static long adjacentProduct(long[] numbers, int size) {
        long[] first13 = Arrays.copyOfRange(numbers, 0, size);
        long largest = Arrays.stream(first13).reduce(1, (a, b) -> a * b);
        for (int i = size; i < numbers.length; i++) {
            long[] current13 = Arrays.copyOfRange(numbers, i-size, i);
            long currentProduct = Arrays.stream(current13).reduce(1, (a, b) -> a * b);
            if (currentProduct > largest)
                largest = currentProduct;
        }
        return largest;
    }
}
