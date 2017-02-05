fn sum(n: i32, k: i32) -> i32 {
	return k * ((((n - 1) / k) * (((n - 1) / k) + 1)) / 2);
}

fn main() {
    println!("{}", sum(1000, 3) + sum(1000, 5) - sum(1000, 15));
}
