fn binomial_coefficient(n: u64, k: u64) -> u64 {
    if k > n {
        return 0;
    }
    if k == n {
        return 1;
    }
    let x = if n - k < k { n - k } else { k };
    let mut c: u64 = 1;
    for i in 0..x {
        c = c * (n - i) / (i + 1);
    }
    return c;
}

fn main() {
    print!("{}\n", binomial_coefficient(40, 20));
}
