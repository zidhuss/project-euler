fn sumDivisors(n : i64) -> i64 {
    let mut sum = 0;
    let mut i = 2;
    while i*i < n {
        if (n%i) == 0 {
            sum += i + (n / i);
        }
        if i*i == n {
            sum += i;
        }
        i += 1;
    }
    return sum+1;
}

fn main() {
    let mut amicableSum = 0;
    for a in 2..10000 {
        let mut b = sumDivisors(a);
        if b > a && sumDivisors(b) == a {
            amicableSum += a + b;
        }
    }
    println!("{}", amicableSum);
}
