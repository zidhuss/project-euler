fn main() {
    let mut sum_square = 0;
    let mut square_sum = 0;
    for i in 1..100 {
        square_sum += i * i;
        sum_square += i;
    }
    println!("{}", sum_square * sum_square - square_sum);
}
