fn main() {
    let input = include_str!("./input.txt");
    let output = part1(input);
    dbg!(output);
}

fn part1(input: &str) -> String {
    let sum = input.lines().map(|v| line_value(v)).fold(0, |acc, x| acc + x);
    return sum.to_string();
}

fn line_value(input: &str) -> u32 {
   let digits:Vec<_> = input.chars().filter(|c| c.is_numeric()).map(|c| c.to_digit(10).unwrap()).collect();
   let tens = digits.first().unwrap().clone();
   let ones = digits.last().unwrap().clone();
   let value = tens*10 + ones;
   dbg!(value);
    return value;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = part1("1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet");
        assert_eq!(result, "142")
    }

    #[test]
    fn test_line_1_value() {
        let result = line_value("1abc2");
        assert_eq!(12, result)
    }

    #[test]
    fn test_line_2_value() {
        let result = line_value("pqr3stu8vwx");
        assert_eq!(38, result)
    }
}