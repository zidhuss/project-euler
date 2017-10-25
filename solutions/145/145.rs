fn all_odd(mut n : i32) -> bool {
    while n > 0 {
        let x = n % 10;
        if x % 2 == 0 {
            return false;
        }
        n /= 10;
    }
    return true;
}

fn reverse(mut n : i32) -> i32 {
    let mut res = 0;
    while n > 0 {
        let x = n % 10;
        res *= 10;
        res += x;
        n /= 10;
    }
    return res;
}

fn main() {
    let mut count = 0;
    for i in 11..100000000 {
        if i % 10 == 0 {
            continue;
        }
        if all_odd(i + reverse(i)) {
            count += 1;
        }
    }
    println!("{}", count);
}
