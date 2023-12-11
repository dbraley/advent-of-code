// use std::collections::HashMap;

// use regex::Regex;

use itertools::Itertools;
use std::cmp;

use std::collections::{HashMap, HashSet};

fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

// Reused from day-10
type Col = i64;
type Row = i64;
// type Point = (Col, Row);

// fn north(p: Point) -> Point {
//     (p.0, p.1 - 1)
// }
// fn east(p: Point) -> Point {
//     (p.0+1, p.1)
// }
// fn west(p: Point) -> Point {
//     (p.0 - 1, p.1)
// }
// fn south(p: Point) -> Point {
//     (p.0, p.1 + 1)
// }

fn str_to_coords_with_specials(input: &str, specials: Vec<char>) -> (HashMap<(Col, Row), char>, HashMap<char, (Col,Row)>) {
    let mut coord_map = HashMap::new();
    let mut special_map = HashMap::new();
    for (row, line) in input.lines().enumerate() {
        for (col, c) in line.chars().enumerate() {
            let p = (col as Col, row as Row);
            coord_map.insert(p, c);
            if specials.contains(&c) {
                special_map.insert(c, p);
            }
        }
    }
    return (coord_map, special_map);
}

fn part2(input: &str) -> String {
    let (coords, _) = str_to_coords_with_specials(input, Vec::new());
    let galaxies:Vec<(i64,i64)> = coords.iter()
    .filter(|entry| entry.1 == &'#')
    .map(|g| g.0.clone())
    .collect();

    let mut empty_rows = HashSet::new();
    for i in 0..input.lines().count() {
        empty_rows.insert(i as i64);
    }
    let mut empty_cols = HashSet::new();
    for i in 0..input.lines().next().unwrap().len() {
        empty_cols.insert(i as i64);
    }

    for (c, r) in &galaxies {
        empty_cols.remove(c);
        empty_rows.remove(r);
    }

    let mut sum_dist = 0;
    for pair in galaxies.into_iter().combinations(2){
        let g1 = pair.get(0).unwrap();
        let g2 = pair.get(1).unwrap();
        let min_col = cmp::min(g1.0, g2.0);
        let max_col = cmp::max(g1.0, g2.0);
        let extra_cols = empty_cols.iter().filter(|c| *c > &min_col && *c < &max_col).count() as i64;
        let min_row = cmp::min(g1.1, g2.1);
        let max_row = cmp::max(g1.1, g2.1);
        let extra_rows = empty_rows.iter().filter(|c| *c > &min_row && *c < &max_row).count() as i64;

        sum_dist += (max_col - min_col) + (extra_cols*999999) + (max_row - min_row) + (extra_rows*999999);
    }
    return sum_dist.to_string();
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_example_1() {
        let actal = part2("\
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....");
        assert_eq!("374", actal);
    }

}