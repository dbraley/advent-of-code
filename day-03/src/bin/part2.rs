use std::collections::HashMap;
fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

fn part2(input: &str) -> String {
    let (gears, parts) = read(input);
    let mut sum = 0;
    for (key, _) in gears {
        let maybe_gears:Vec<&Part> = parts.iter()
            .filter(|p| p.in_range(key.0, key.1))
            .collect();
        // dbg!(maybe_gears);
        if maybe_gears.len() == 2 {
            let val = maybe_gears[0].value * maybe_gears[1].value;
            dbg!(val);
            sum += val;
        }
    }
    return sum.to_string();
}

#[derive(Debug)]
#[derive(Clone, Copy,PartialEq)]
struct Part {
    value: u32,
    row: isize,
    start: isize,
    end: isize,
}

impl Part {
    fn in_range(&self, row: isize, col: isize) -> bool {
        if row >= self.row - 1 && row <= self.row + 1 {
            if col >= self.start-1 && col <= self.end+1 {
                return true;
            }
        }
        return false;
    }
}

fn read(input: &str) -> (HashMap<(isize, isize), char>, Vec<Part>) {
    let mut making = false;
    let mut cur_part:Part = Part { value: 0, row: 0, start: 0, end: 0 };
    let mut parts:Vec<Part> = Vec::new();
    let mut points:HashMap<(isize, isize), char> = HashMap::new();
    for (row, line) in input.lines().enumerate() {
        // dbg!(row, line);
        for (col, c) in line.chars().enumerate() {
            // dbg!(col, c);
            if c == '*' {
                points.insert((isize::try_from(row).unwrap_or(0), isize::try_from(col).unwrap_or(0)), c);
            }
            if c.is_digit(10) {
                // dbg!(row, col, c);
                let v = c.to_digit(10).unwrap();
                if !making {
                    cur_part = Part { value: v, row: isize::try_from(row).unwrap_or(0), start: isize::try_from(col).unwrap_or(0), end: isize::try_from(col).unwrap_or(0) };
                    making = true;
                } else {
                    let old_part = cur_part.clone();
                    cur_part = Part {
                        value: old_part.value*10+v, 
                        row: old_part.row,
                        start: old_part.start,
                        end: isize::try_from(col).unwrap_or(0), 
                    }; 
                }
            } else {
                if making {
                    making = false;
                    parts.push(cur_part.to_owned());
                }
            }
        }
        if making {
            making = false;
            parts.push(cur_part.to_owned());
        }    
    }
    return (points, parts);
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_it_works() {
        let input = "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..";
        assert_eq!("467835", part2(input))
    }

    #[test]
    fn test_edge() {
        let input ="467..114..
...*......
..35..633.";
        assert_eq!("16345", part2(input));
    }

    #[test]
    fn test_find_parts() {
        let input = "467";
        let expected = vec![
            Part{value: 467, row: 0, start: 0, end: 2},
        ];

        let (_, actual) = read(input);
        assert_eq!(expected, actual);
    }
}