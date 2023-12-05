

fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

fn part2(input: &str) -> String {
    let sum = input.lines().map(|v| line_value(v)).fold(0, |acc, x| acc + x);
    return sum.to_string();
}

fn line_value(input: &str) -> u32 {
    let input2 = input
        // Normal cases
        .replace("one", "one1one")
        .replace("two", "two2two")
        .replace("three", "three3three")
        .replace("four", "four4four")
        .replace("five", "five5five")
        .replace("six", "six6six")
        .replace("seven", "seven7seven")
        .replace("eight", "eight8eight")
        .replace("nine","nine9nine");
    let digits:Vec<_> = input2.chars().filter(|c| c.is_numeric()).map(|c| c.to_digit(10).unwrap()).collect();
    let tens = digits.first().unwrap().clone();
    // if digits.len() == 1 {
    //     dbg!(input, input2);
    //     return tens;
    // }
    let ones = digits.last().unwrap().clone();
    let value = tens*10 + ones;
    return value;
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn it_works() {
        let result = part2("two1nine
        eightwothree
        abcone2threexyz
        xtwone3four
        4nineeightseven2
        zoneight234
        7pqrstsixteen");
        assert_eq!(result, "281")
    }

    #[test_case("two1nine", 29)]
    #[test_case("eightwothree", 83)]
    #[test_case("abcone2threexyz", 13)]
    #[test_case("xtwone3four", 24)]
    #[test_case("4nineeightseven2", 42)]
    #[test_case("zoneight234", 14)]
    #[test_case("7pqrstsixteen", 76)]
    // If a line contains only one number, it is both the first and last.
    #[test_case("1", 11)]
    // When spelled out numbers overlap, they both count
    #[test_case("oneight", 18)]
    fn line_test(input: &str, expected: u32) {
        assert_eq!(expected, line_value(input))
    }
}