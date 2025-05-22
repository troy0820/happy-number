use std::collections::HashMap;
use std::collections::HashSet;
use std::io;

pub fn get_input_for_happy_number(mut s: String) -> i32 {
    println!("Pick a number that is happy:");
    io::stdin().read_line(&mut s).expect("failed to read line");
    let num = s.trim().parse::<i32>().expect("failed to parse number");
    return num;
}

fn process_number(mut x: i32) -> i32 {
    let mut sum = 0;
    let mut num;

    while 0 < x {
        num = x % 10;
        sum += num * num;
        x = x / 10;
    }
    return sum;
}

pub fn happy_number_hash_map(mut n: i32) -> bool {
    let mut m: HashMap<i32, bool> = HashMap::new();

    loop {
        if n == 1 {
            return true;
        }
        n = process_number(n);

        if m.get_mut(&n) == Some(&mut true) || n == 0 {
            return false;
        }
        m.insert(n, true);
    }
}

pub fn happy_number_hash_set(mut n: i32) -> bool {
    let mut m = HashSet::new();

    loop {
        if n == 1 {
            return true;
        }
        n = process_number(n);

        if m.contains(&n) || n == 0 {
            return false;
        }
        m.insert(n);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn should_return_true_with_hash_map() {
        let x = 13;
        let is_true = happy_number_hash_map(x);
        assert_eq!(is_true, true);
    }

    #[test]
    fn should_return_true_with_hash_set() {
        let x = 13;
        let is_true = happy_number_hash_set(x);
        assert_eq!(is_true, true);
    }

    #[test]
    fn should_return_false_with_hash_map() {
        let x = 2;
        let is_false = happy_number_hash_map(x);
        assert_eq!(is_false, false);
    }

    #[test]
    fn should_return_false_with_hash_set() {
        let x = 2;
        let is_false = happy_number_hash_set(x);
        assert_eq!(is_false, false);
    }
}
