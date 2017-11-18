use std::fs::File;
use std::io::Read;

fn main() {
    let mut numbers: [u8;1000] = [0;1000];

    let file = File::open("./resources/digits.txt");
    let mut i = 0;

    for byte in file.unwrap().bytes() {
        let b = byte.unwrap();
        if b >= 48 {
            numbers[i] = b - 48;
            i += 1;
        }
    }

    let size = 13;
    let mut largest: i64 = 0;

    for j in size..(1000-size) {
        let mut product: i64 = 1;
        for k in 0..size {
            product *= numbers[j+k] as i64;
        }
        if product > largest {
            largest = product;
        }
    }

    println!("{}", largest);
}
