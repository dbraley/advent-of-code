use std::collections::HashMap;

use regex::Regex;

fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

fn part2(input: &str) -> String {
    let directions = input.lines().next().unwrap();
    dbg!(directions.len());


    let (coords, start_locs) = make_coord_map(input);

    // let mut cur_locs = start_locs;
    let mut product = 1;
    for i in 0..start_locs.len() {
        product *= find_solutions(start_locs[i].as_str(), directions, coords.to_owned());
        println!("");
    }

    // This is a pretty sketchy solution. It makes a number of assumptions to simplify the code, each of which could 
    // easily be wrong for a given data set, but do seem to hold true for mine.
    //  0. It assumes that a path will result in finding a solution more than once. While this doesn't have to be 
    //      true, the problem is most likely unsolvable if this isn't true.
    //  1. It assumes that the cadence of finding a solution on a given path is consistent. This is not necessarily 
    //      true though, it is possible that a given loop crosses a vialbe a 
    //  2. It assumes that 
    return (product * (directions.len() as i64)).to_string();
}

fn find_solutions(start_loc: &str, directions: &str, coords: HashMap<String, (String, String)>) -> i64 {
    let mut steps = 0;
    let mut solutions = 0;
    let mut last_found = 0;
    let mut cur_loc = start_loc;
    
    for i in 0..100 {
        // dbg!(i);
        for next_dir in directions.chars() {
            if cur_loc.ends_with("Z") {
                solutions += 1;
                println!("Found solution for {} at {} on step {}:{}({})", start_loc, cur_loc, steps, steps - last_found, i);
                last_found = steps;
                return i;
                // if solutions >= 5 {
                //     return i;
                // }
            }

            steps += 1;

            let possible: &(String, String) = coords.get(cur_loc).unwrap();

            match next_dir {
                'L' => cur_loc = possible.0.as_str(),
                _ => cur_loc = possible.1.as_str(),
            }

        }
    }
    return -1;
}

fn make_coord_map(input: &str) -> (HashMap<String, (String, String)>, Vec<String>) {
    let mut coords:HashMap<String, (String, String)> = HashMap::new();
    let mut start_locs:Vec<String> = Vec::new();
    let re = Regex::new(r"(?m)^([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)$").unwrap();

    for (_, [key, left, right]) in re.captures_iter(input).map(|c| c.extract()) {
        coords.insert(key.to_string(), (left.to_string(), right.to_string()));
        if key.ends_with("A") {
            dbg!(key);
            start_locs.push(key.to_string());
        }
    }
    return (coords, start_locs);
}

fn next_locs(coords: HashMap<String, (String, String)>, cur_locs: Vec<String>) -> Vec<String> {
    let res = Vec::new();

    return res;
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_it_works() {
        let actal = part2("\
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
        let actal = part2("\
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)");
        assert_eq!("6", actal);
    }

    #[test]
    fn test_it_works3() {
        let actal = part2("\
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
");
        assert_eq!("6", actal);
    }

    #[test_case("23456", 10203040506)]    
    fn test_cards_to_int(_: &str, _: i64) {
    }

}