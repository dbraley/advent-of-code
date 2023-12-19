use std::collections::HashMap;

fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

fn part2(input: &str) -> String {
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
    // if num_unknowns < 2 {
        //     return 1;
        // }
        // let num_known_broken:i64 = line.chars().filter(|c| c == &'#').count() as i64;
    let parts:Vec<&str> = line.split(' ').collect();

    let springs_1  = parts.get(0).unwrap();
    let springs = vec![*springs_1, springs_1, springs_1, springs_1, springs_1].clone().join("?");

    let num_unknowns = springs.chars().filter(|c| c == &'?').count();

    let counts_str_1 = parts.get(1).unwrap().replace(",", " ");
    let counts_str = vec![counts_str_1.clone(), counts_str_1.clone(), counts_str_1.clone(), counts_str_1.clone(), counts_str_1].clone().join(" ");
    let counts = line_to_num_vec(&counts_str);

    let total_broken:i64 = counts.iter().sum();
    let min_size = total_broken + counts.len() as i64 - 1;
    println!("len {}, min_size {}, breaks {}, unknowns {}", springs.len(), min_size, counts.len() + 1, num_unknowns);
    let mut cache:HashMap<(usize, usize), i64> = HashMap::new();
    return score_recurse(springs.as_str(), counts.as_slice(), min_size, &mut cache);
}

fn score_recurse(springs: &str, counts: &[i64], min_size_remaining: i64, cache: &mut HashMap<(usize, usize), i64>) -> i64 {
    if (springs.len() as i64) < min_size_remaining {
        return 0;
    }
    if springs.len() == 0 {
        if counts.len() == 0 {
            return 1;
        } else {
            return 0;
        }
    }
    // check cache
    if let Some(cached) = cache.get(&(springs.len(), counts.len())) {
        return *cached;
    }
    if springs.starts_with('.') {
        let score = score_recurse(&springs[1..], counts, min_size_remaining, cache);
        cache.insert((springs.len(), counts.len()),score);
        return score;
    }
    if springs.starts_with('?') {
        let operational_count = score_recurse(&springs[1..], &counts, min_size_remaining, cache);
        let broken_count = assume_broken_count(springs, counts, min_size_remaining, cache);
        let score = broken_count + operational_count;
        cache.insert((springs.len(), counts.len()),score);
        return score;
    }
    if springs.starts_with('#') {  
        let score = assume_broken_count(springs, counts, min_size_remaining, cache);
        cache.insert((springs.len(), counts.len()),score);
        return score;
    }
    panic!("Not ready yet!");
}

fn assume_broken_count(springs: &str, counts: &[i64], min_size_remaining: i64, cache: &mut HashMap<(usize, usize), i64>) -> i64 {
    match counts.first() {
        Some(group) => {
            for i in 0..*group {
                match springs.chars().nth(i as usize) {
                    Some('.') => return 0,
                    Some('#') => {}, // keep checking
                    Some('?') => {}, // keep checking
                    None => return 0,
                    _ => panic!("unknown char (1)")
                }
            }
            match springs.chars().nth(*group as usize) {
                Some('.') => {  // Group ends and next spring is operational
                    let next_springs:&str = &springs[(*group as usize + 1)..];
                    return score_recurse(next_springs, &counts[1..], min_size_remaining - group - 1, cache);
                },
                Some('?') => { // Group ends and next spring is unknown but operational
                    let next_springs:&str = &springs[(*group as usize + 1)..];
                    return score_recurse(next_springs, &counts[1..], min_size_remaining - group - 1, cache);
                },
                Some('#') => return 0, // Group continues after the count, this config is invalid
                None => { // At the end of the string, but still need to check if remaining counts
                    let next_springs:&str = &springs[(*group as usize)..];
                    return score_recurse(next_springs, &counts[1..], min_size_remaining - group, cache);
                },
                _ => panic!("unknown char (2)")
                
            }
        },
        None => return 0,
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_example_1() {
        let actal = part2("\
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
");
        assert_eq!("525152", actal);
    }

    #[test_case("???.### 1,1,3", 1; "line 1")]
    #[test_case(".??..??...?##. 1,1,3", 16384; "line 2")]
    #[test_case("?#?#?#?#?#?#?#? 1,3,1,6", 1; "line 3")]
    #[test_case("????.#...#... 4,1,1", 16; "line 4")]
    #[test_case("????.######..#####. 1,6,5", 2500; "line 5")]
    #[test_case("?###???????? 3,2,1", 506250; "line 6")]
    #[test_case("??##???????????#? 5,2,1,2", 1; "actual line 34")]
    fn test_score(input: &str, expected: i64) {
        assert_eq!(score(input), expected);
    }

    #[test_case("", &[], 1; "done")]
    #[test_case("", &[1], 0; "invalid: out of chars when needed")]
    #[test_case(".", &[], 1; "out of broken groups")]
    #[test_case("#", &[], 0; "invalid: out of gropus when needed")]
    #[test_case("#", &[1], 1; "1 broken group of 1")]
    #[test_case("###", &[3], 1; "3 broken group of 3")]
    #[test_case("#?#", &[3], 1; "3 broken or unknown group of 3")]
    #[test_case("#?#.", &[3], 1; "3 broken or unknown group of 3 followed by operational")]
    #[test_case("#?#?", &[3], 1; "3 broken or unknown group of 3 followed by unknown")]
    #[test_case("#?##", &[3], 0; "3 broken or unknown group of 4")]
    fn test_score_recurse(s: &str, c: &[i64], expected: i64) {
        let mut cache: HashMap<(usize, usize), i64> = HashMap::new();
        assert_eq!(expected, score_recurse(s, c, 0, &mut cache))
    }

}