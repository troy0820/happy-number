mod happy;

fn main() {
    let happy_number = String::new();
    let num: i32 = happy::get_input_for_happy_number(happy_number);
    println!(
        "Happy number with hash set {} is : {:?}",
        num,
        happy::happy_number_hash_set(num)
    );
    println!(
        "Happy number with hash map {} is : {:?}",
        num,
        happy::happy_number_hash_map(num)
    );
}
