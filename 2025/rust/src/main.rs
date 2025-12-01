#![feature(int_roundings)]
use std::{env, fs::read_to_string};
mod days;

use crate::days::one::part_one;

fn main() {
  let args: Vec<String> = env::args().collect();
  let day = &args[1];
  match day.as_str() {
    "1" => {
      part_one(handle_file("1"));
    }
    _ => {}
  }
}

fn handle_file(day: &str) -> Vec<String> {
  // NOTE: I should probably not make this a static path
  read_to_string(format!("/home/rond/code/aoc/src/input/day00{}.txt", day))
    .expect(format!("Couldn't load file: day{}.txt", day).as_str())
    .lines()
    .map(String::from)
    .collect()
}
