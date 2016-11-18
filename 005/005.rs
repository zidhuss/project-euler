fn gcd(mut a : i32 , mut b : i32) -> i32 {
	while b != 0 {
        let c = b;
        b = a % b;
        a = c;
	}
	return a
}

fn lcm(a : i32, b : i32) -> i32 {
    return a * b / gcd(a, b);
}

fn main() {
    let mut result = 1;
    for i in 1..20 {
        result = lcm(result, i);
    }
    println!("{}", result);
}
