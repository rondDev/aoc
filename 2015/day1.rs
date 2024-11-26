use std::fs::File;
use std::io::Read;
fn main() -> std::io::Result<()> {
    let mut input = String::new();
    let mut file = File::open("inputs/day1.txt")?;
    file.read_to_string(&mut input)?;
    let result = solve(&input, false);
    println!("Your answer for part one is: {}", result);
    let result = solve(&input, true);
    println!("Your answer for part two is: {}", result);
    Ok(())
}

fn solve(input: &str, part_two: bool) -> i32 {
    let mut total = 0;
    for (i, ch) in input.chars().enumerate() {
        match ch {
            '(' => {
                total += 1;
            }
            ')' => {
                total -= 1;
            }
            _ => {
                // Do nothing
            }
        }
        if part_two && total < 0 {
            return (i + 1) as i32; // Have to add 1 since index is 0-based
        }
    }
    total
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_part_one() {
        assert_eq!(0, solve("(())"));
        assert_eq!(3, solve("((("));
        assert_eq!(3, solve("))((((("));
    }
}
