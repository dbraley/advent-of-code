// use std::collections::HashMap;

use regex::Regex;

use itertools::Itertools;

use core::num;
use std::collections::HashSet;

fn main() {
    let input = include_str!("./input.txt");
    let output = part1(input);
    dbg!(output);
}

fn part1(input: &str) -> String {
    let score:i64 = input.lines()
        .map(|line| score(line))
        .sum();
    return score.to_string();
}

// From day-09
fn line_to_num_vec(line: &str) -> Vec<i64> {
    return line.split(' ')
        // .inspect(|line| _ = dbg!(line))
        .map(|s| s.parse::<i64>().unwrap())
        .collect();
}

fn score(line: &str) -> i64 {
    let num_unknowns = line.chars().filter(|c| c == &'?').count();
    // if num_unknowns < 2 {
    //     return 1;
    // }
    // let num_known_broken:i64 = line.chars().filter(|c| c == &'#').count() as i64;
    let parts:Vec<&str> = line.split(' ').collect();
    let springs  = parts.get(0).unwrap().replace(".", "x");
    let counts = line_to_num_vec(parts.get(1).unwrap().replace(",", " ").as_str());
    let total_broken:i64 = counts.iter().sum();
    let min_size = total_broken + counts.len() as i64 - 1;
    println!("len {}, min_size {}, breaks {}, unknowns {}", springs.len(), min_size, counts.len() + 1, num_unknowns);
    // let mut missing:Vec<&str> = Vec::new();
    // for i in 0..num_unknowns as i64 {
    //     if i < (total_broken - num_known_broken) {
    //         missing.push("#");
    //     } else {
    //         missing.push("x");
    //     }
    // }
    // let mut count = 0;
    let reg = get_validation_regex(counts);
    // let permutations:HashSet<_> = missing.iter().permutations(num_unknowns).unique().collect();
    // for perm in permutations {
    //     let possible = perm.into_iter().fold(springs.to_string(), |springs, spring| springs.replacen("?", &spring, 1));
    //     if reg.is_match(&possible) {
    //         count += 1;
    //     }
    // }
    // println!("{} - {}", line, count);
    // return count;
    return score_recurse(springs, num_unknowns as i64, &reg);
}

fn score_recurse(springs: String, num_unknowns: i64, reg: &Regex) -> i64 {
    if num_unknowns <= 0 {
        if reg.is_match(springs.as_str()) {
            return 1;
        } else {
            return 0;
        }
    }
    return score_recurse(springs.replacen("?", "x", 1), num_unknowns-1, &reg)
        + score_recurse(springs.replacen("?", "#", 1), num_unknowns-1, &reg);
}

fn get_validation_regex(v: Vec<i64>) -> Regex {
    let pts:Vec<_> = v.iter().map(|i| format!(r"[#]{{{}}}", i)).collect();
    // let reg_str = pts.join(r"x+");
    let reg_str = format!(r"^x*{}x*$", pts.join(r"x+"));
    return Regex::new(reg_str.as_str()).unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_example_1() {
        let actal = part1("\
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
");
        assert_eq!("21", actal);
    }

    #[test_case("#.#.### 1,1,3", 1 ; "No unknowns")]
    #[test_case("#?#.### 1,1,3", 1 ; "Only one unknown (good)")]
    #[test_case("?.#.### 1,1,3", 1 ; "Only one unknown (bad)")]
    #[test_case(".??...#...###. 1,1,3", 2; "Two permutations, two unknowns")]
    #[test_case(".??...#...?##. 1,1,3", 2; "Two permutations")]
    #[test_case("???.### 1,1,3", 1 ; "Three permutations, but only one is valid")]
    #[test_case(".??..??...?##. 1,1,3", 4)]
    fn test_score(input: &str, expected: i64) {
        assert_eq!(score(input), expected);
    }

    #[test]
    fn test_get_validation_regex() {
        let r = get_validation_regex(vec![1,1,3]);
        assert_eq!(true, r.is_match("#x#x###"));
        assert_eq!(false, r.is_match("##xx###"));
        assert_eq!(false, r.is_match("x##xxx#xxxx##x"));
        assert_eq!(false, r.is_match("x##xxx#xxx###x"));
    }

}