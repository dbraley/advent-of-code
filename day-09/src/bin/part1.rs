// use std::collections::HashMap;

// use regex::Regex;

fn main() {
    let input = include_str!("./input.txt");
    let output = part1(input);
    dbg!(output);
}

fn part1(input: &str) -> String {
    let res = input.lines()
        .inspect(|line| _ = dbg!(line))
        .map(|line| line_to_num_vec(line))
        .map(|vec| next_val(vec))
        .fold(0, |acc, x| acc + x);
    return res.to_string();
}

fn line_to_num_vec(line: &str) -> Vec<i64> {
    return line.split(' ')
        // .inspect(|line| _ = dbg!(line))
        .map(|s| s.parse::<i64>().unwrap())
        .collect();
}

fn next_val(numbers: Vec<i64>) -> i64 {
    let mut line = numbers;
    let mut guess = 0;
    for _ in 0..100 {
        if line.iter()
                // .inspect(|v| _ = dbg!(v))
                .filter(|v| !0.eq(*v))
                .count() == 0 {
            return guess;
        }
        guess += line.last().unwrap();
        line = deltas(line);
    }
    panic!("Need more iterations in next_val");
}

fn deltas(line: Vec<i64>) -> Vec<i64> {
    let mut next_vec:Vec<i64> = Vec::new();
    let mut left = line.first().unwrap();
    // Intentionally skip the first element
    for i in 1..line.len() {
        let value = line.get(i).unwrap();
        let delta = value - left;
        next_vec.push(delta);
        left = value;
    }
    // return dbg!(next_vec);
    return next_vec;
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_it_works() {
        let actal = part1("\
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
");
        assert_eq!("114", actal);
    }

    #[test_case("0 3 6 9 12 15", 18)]    
    #[test_case("1 3 6 10 15 21", 28)]    
    #[test_case("10 13 16 21 30 45", 68)]    
    fn test_next_val(input: &str, expected: i64) {
        assert_eq!(expected, next_val(line_to_num_vec(input)));
    }

}