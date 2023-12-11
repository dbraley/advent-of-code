use std::collections::HashMap;

use regex::Regex;

fn main() {
    let input = include_str!("./input.txt");
    let output = part1(input);
    dbg!(output);
}

fn part1(input: &str) -> String {
    let directions = input.lines().next().unwrap();
    dbg!(directions);


    let mut coords: HashMap<&str, (&str, &str)> = HashMap::new();
    // let re = Regex::new(r"(?m)^([^:]+):([0-9]+):(.+)$").unwrap();
    let re = Regex::new(r"(?m)^([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)$").unwrap();

    for (_, [key, left, right]) in re.captures_iter(input).map(|c| c.extract()) {
        coords.insert(key, (left, right));
    }

    let mut cur_loc = "AAA";
    let mut steps = 0;
    for i in 0..100 {
        dbg!(i);
        for next_dir in directions.chars() {
            steps += 1;
            dbg!(cur_loc, steps);
            let possible = coords.get(cur_loc).unwrap();
            match next_dir {
                'L' => cur_loc = possible.0,
                _ => cur_loc = possible.1,
            }
            if cur_loc == "ZZZ" {
                dbg!(cur_loc);
                return steps.to_string();
            }
        }
    }

    return "".to_string();
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_it_works() {
        let actal = part1("\
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)");
        assert_eq!("2", actal);
    }
    
    #[test]
    fn test_it_works2() {
        let actal = part1("\
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)");
        assert_eq!("6", actal);
    }

    #[test_case("23456", 10203040506)]    
    fn test_cards_to_int(_: &str, _: i64) {
    }

}