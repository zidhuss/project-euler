fn main() {
    let mut fibo: [u32; 36] = [0; 36];
    fibo[1] = 1;
    for i in 2..36 {
        fibo[i] = fibo[i - 2] + fibo[i - 1];
    }
    let mut s = 0;
    for i in 2..36 {
        if fibo[i] % 2 == 0 {
            s += fibo[i];
        }
    }
    println!("{}", s);
}
