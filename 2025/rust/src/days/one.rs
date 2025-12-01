use std::time;

pub fn part_one(input: Vec<String>) {
  let naive_one = time_function(naive, input.clone());
  println!(
    "Password for part 1 (naive): {} | Time elapsed: {:?}",
    naive_one.0, naive_one.1
  );
  let naive_two = time_function(naive_two, input.clone());
  println!(
    "Password for part 2 (naive): {} | Time elapsed: {:?}",
    naive_two.0, naive_two.1
  );
  let math_one = time_function(math_one, input.clone());
  println!(
    "Password for part 1 (math): {} | Time elapsed: {:?}",
    math_one.0, math_one.1
  );
  let red = time_function(reduction, input.clone());
  println!(
    "Password for part 1 (reduce): {} | Time elapsed: {:?}",
    red.0, red.1
  );
}

fn time_function(fun: fn(Vec<String>) -> i128, arg: Vec<String>) -> (i128, time::Duration) {
  let start_time = time::Instant::now();
  let part_one = fun(arg);
  let total_time = start_time.elapsed();
  (part_one, total_time)
}

fn reduction(input: Vec<String>) -> i128 {
  let mut current_value = 50;
  input.iter().fold(0, |acc, n| {
    let mut num: i128 = n[1..].parse().unwrap();
    if n.chars().nth(0).unwrap() == 'L' {
      num = -num;
    }
    current_value = (current_value + num) % 100;
    if current_value == 0 {
      return acc + 1;
    }
    acc
  })
}

fn math_one(input: Vec<String>) -> i128 {
  let mut total_zero = 0;
  let mut current_value = 50;
  for i in 0..input.len() {
    let direction: &char = &input[i].chars().nth(0).unwrap();
    let mut num: i128 = input[i][1..].parse().unwrap();
    if direction == &'L' {
      num = -num;
    }
    current_value = (current_value + num) % 100;
    if current_value == 0 {
      total_zero += 1;
    }
  }
  total_zero
}

fn naive(input: Vec<String>) -> i128 {
  // NOTE: Naive loop, bad performance
  let mut total_zero = 0;
  // NOTE: It starts at 50
  let mut current_value = 50;
  for n in 0..input.len() {
    let direction = &input[n].chars().nth(0).unwrap();
    let num: i128 = input[n][1..].parse().unwrap();
    for _ in 0..num {
      if direction == &'L' {
        if current_value - 1 < 0 {
          current_value = 99;
        } else {
          current_value -= 1;
        }
      } else if direction == &'R' {
        if current_value + 1 > 99 {
          current_value = 0;
        } else {
          current_value += 1;
        }
      }
    }
    if current_value == 0 {
      total_zero += 1;
    }
  }
  total_zero
}

fn naive_two(input: Vec<String>) -> i128 {
  // NOTE: Naive loop, bad performance
  let mut total_zero = 0;
  // NOTE: It starts at 50
  let mut current_value = 50;
  for n in 0..input.len() {
    let direction = &input[n].chars().nth(0).unwrap();
    let num: i128 = input[n][1..].parse().unwrap();
    for i in 0..num {
      if direction == &'L' {
        if current_value - 1 < 0 {
          current_value = 99;
        } else {
          current_value -= 1;
        }
      } else if direction == &'R' {
        if current_value + 1 > 99 {
          current_value = 0;
        } else {
          current_value += 1;
        }
      }
      if current_value == 0 {
        total_zero += 1;
      }
    }
  }
  total_zero
}

#[test]
fn test_naive_parts() {
  let input = [
    "L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82",
  ]
  .map(String::from)
  .to_vec();

  let first = naive(input.clone());
  assert_eq!(first, 3);
  let second = naive_two(input);
  assert_eq!(second, 6);
}

#[test]
fn test_math_parts() {
  let input = [
    "L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82",
  ]
  .map(String::from)
  .to_vec();

  let first = math_one(input.clone());
  assert_eq!(first, 3);
  // let second = naive_two(input);
  // assert_eq!(second, 6);
}

#[test]
fn test_recursive_parts() {
  let input = [
    "L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82",
  ]
  .map(String::from)
  .to_vec();

  let first = launch_recur(input.clone());
  assert_eq!(first, 3);
}

#[test]
fn longer_test_part_one() {
  let input = [
    "L3", "R49", "R11", "L18", "R41", "L49", "R40", "R38", "R8", "L8", "L14", "L45", "R38", "L46",
    "R45", "R37", "R24", "L43", "L45", "L44", "L2", "R22", "R8", "R50", "R49",
  ]
  .map(String::from)
  .to_vec();

  assert_eq!(naive(input.clone()), 1);
  assert_eq!(math_one(input.clone()), 1);
  assert_eq!(launch_recur(input), 1);
}
