fn is_palindrome(orig : i32) -> bool {
    let mut rev = 0;
    let mut n = orig;
    while n > 0 {
        rev = rev * 10 + n % 10;
        n /= 10;
    }
    return orig == rev;
}

fn main() {
    let mut largest = 0;
    let last = 999;
    let first = 100;

    let mut a = last;
    while a >= first {
        let mut b = last;
        while b >= first {
            let calculated = a * b;
            if calculated > largest && is_palindrome(calculated) {
                largest = calculated;
            }
            b -= 1;
        }
        a -= 1;
    }
    println!("{}", largest);
}
