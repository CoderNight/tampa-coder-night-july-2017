use std::env;
use std::fs::File;
use std::io::BufReader;
use std::collections::HashMap;
use std::io::prelude::*;

fn main() {
    let file = get_file(env::args().collect()).unwrap();
    let reader = BufReader::new(file);
    let task = read_missing_number_task(reader);
    find_missing_numbers(task)
}

fn read_missing_number_task(reader: BufReader<File>) -> MissingNumberTask {
    let mut lines = reader.lines();
    MissingNumberTask {
        a_length: lines.next().expect("could not read length of first set").unwrap(),
        a_string: lines.next().expect("could not read content of first set").unwrap(),
        b_length: lines.next().expect("could not read length of second set").unwrap(),
        b_string: lines.next().expect("could not read content of second set").unwrap(),
    }
}

fn get_file(args: Vec<String>) -> Result<File, &'static str> {
    if args.len() < 2 {
        return Err("Not enough args");
    }
    Ok(File::open(&args[1]).expect("Could not open file"))
}

#[derive(Debug)]
struct MissingNumberTask {
    a_length: String,
    a_string: String,
    b_length: String,
    b_string: String,
}

fn find_missing_numbers(task: MissingNumberTask) {
    let numbers_a = parse_numbers(task.a_string);
    let numbers_b = parse_numbers(task.b_string);
    let num_freq_b = calculate_number_frequency(numbers_b);
    let missing = decrement_frequencies(num_freq_b, numbers_a);
    print_frequencies(missing);
}

fn parse_numbers(num_string: String) -> Vec<i16> {
    num_string
        .split_whitespace()
        .map( |n| n.parse::<i16>().unwrap() )
        .collect()
}

fn calculate_number_frequency(num_array: Vec<i16>) -> HashMap<i16, i16> {
    let mut result = HashMap::new();
    for v in num_array {
        let count = result.entry(v).or_insert(0);
        *count += 1;
    }
    result
}
 
fn decrement_frequencies(mut freq_map: HashMap<i16, i16>, num_array: Vec<i16>) -> HashMap<i16, i16> {
    for v in num_array {
        let count = freq_map.entry(v).or_insert(0);
        *count -= 1;
    }
    freq_map
}

fn print_frequencies(freqs: HashMap<i16, i16>) {
    let mut result = Vec::new();
    for k in freqs.keys() {
        let mut i = freqs.get(k).unwrap().clone();
        while i > 0 {
            result.push(k);
            i -= 1;
        }
    }
    result.sort();
    for v in result {
        print!("{} ", v);
    }
}
