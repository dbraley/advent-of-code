use std::collections::HashMap;
use std::cmp;

fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

// Library functions

// Reused from day-10
type Col = i64;
type Row = i64;

fn str_to_coords_with_specials(input: &str, specials: Vec<char>) -> (HashMap<(Col, Row), char>, HashMap<char, (Col,Row)>, (Col,Row)) {
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
    let max = input.lines().enumerate().last().map(|(row, line)| ((line.len() - 1) as Col, row as Row)).unwrap_or((-1,-1));
    return (coord_map, special_map, max);
}

// \Library


fn part2(input: &str) -> String {
    let grids = input.split("\n\n");
    let mut score = 0;
    for (n, grid_str) in grids.enumerate() {
        let (grid, _, max) = str_to_coords_with_specials(grid_str, vec![]);
        println!("grid {}: {}x{} (last: {})", n, max.0, max.1, grid.get(&(max.0, max.1)).unwrap_or(&'!'));
        for vertex in 0..max.0 {
            if is_vertical_reflection_point(&grid, vertex, max) {
                score += vertex + 1;
            }
        }
        for vertex in 0..max.1 {
            if is_horizontal_reflection_point(&grid, vertex, max) {
                score += (vertex + 1) * 100;
            }
        }
    }
    return score.to_string();
}

fn is_vertical_reflection_point(grid: &HashMap<(i64, i64), char>, vertex: i64, max: (Col, Row)) -> bool {
    let range = cmp::min(vertex + 1, max.0 - vertex);
    let mut fixed_1 = false;
    // println!("\tChecking for mirror after col {} with width {}", vertex, range);
    for offset in 0..range {
        let left = vertex - offset;
        let right = vertex + 1 + offset;
        for r in 0..(max.1+1) {
            if grid.get(&(left, r)) != grid.get(&(right, r)) {
                if !fixed_1 {
                    fixed_1 = true;
                } else {
                    return false;
                }
            }
        }
    }
    if fixed_1 {
        println!("\t\tCol {} was a mirror!!!", vertex);
    }
    return fixed_1;
}

fn is_horizontal_reflection_point(grid: &HashMap<(i64, i64), char>, vertex: i64, max: (Col, Row)) -> bool {
    let range = cmp::min(vertex + 1, max.1 - vertex);
    let mut fixed_1 = false;
    // println!("\tChecking for mirror after row {} with width {}", vertex, range);
    for offset in 0..range {
        let up = vertex - offset;
        let down = vertex + 1 + offset;
        for c in 0..(max.0+1) {
            if grid.get(&(c, up)) != grid.get(&(c, down)) {
                if !fixed_1 {
                    fixed_1 = true;
                } else {
                    return false;
                }
            }
        }
    }
    if fixed_1 {
        println!("\t\tRow {} was a mirror!!!", vertex);
    }
    return fixed_1;
}

#[cfg(test)]
mod tests {
    use super::*;
    // use test_case::test_case;

    #[test]
    fn test_example_1() {
        let actal = part2("\
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#");
        assert_eq!("400", actal);
    }
}