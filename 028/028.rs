fn a(n: i64) -> i64 {
    return n*(n+2)/4 + (n%4)/3 + 1;
}

fn main() {
    let n = 1001;
    let mut sum = 0;
    let mut i = 1;
    while a(i) <= n*n {
        sum += a(i);
        i += 1;
    }
    println!("{}", sum);
}
