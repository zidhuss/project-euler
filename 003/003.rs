fn main() {
    println!("{}", largest_prime_factor(600851475143))
}

fn largest_prime_factor(n: u64) -> u64 {
    let mut factor: u64 = 1;
    let mut i: u64 = 2;
    let mut x = n;

    while i*i <= x {
        if x == 1 {
            return factor;
        } else if x % 1 != 0 {
            continue;
        } else {
            factor = i;
            while x % i == 0 {
                x /= i
            }
        }
        i += 1;
    }
    return x;
}
