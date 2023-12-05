use std::collections::HashMap;

fn main() {
    let input = include_str!("./input.txt");
    let output = part1(input);
    // not 538937
    // not 540775
    dbg!(output);
}

fn part1(input: &str) -> String {
    let (points, parts) = read(input);
    return parts.iter()
        .filter(|part| part.in_range(&points))
        // .inspect(|x| _ = dbg!(x))
        .fold(0, |acc, x| acc + x.value)
        .to_string();
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
    fn in_range(&self, points: &HashMap<(isize, isize), char>) -> bool {
        for r in (self.row - 1)..(self.row + 2) {
            for c in (self.start - 1)..(self.end + 2) {
                // dbg!(r, c, points.get(&(r,c)));
                // if points.get(&(r,c)).is_some_and(|char| !['.','0','1','2','3','4','5','6','7','8','9'].contains(char)) {
                if points.get(&(r,c)).is_some_and(|char| ['*','+','&','@','=','-','/','#','%','$'].contains(char)) {
                    // dbg!(&self, r, c);
                    return true;
                }
            }
        }
        dbg!(&self.value);
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
            points.insert((isize::try_from(row).unwrap_or(0), isize::try_from(col).unwrap_or(0)), c);
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
        assert_eq!("4361", part1(input))
    }

    #[test]
    fn test_edge() {
        let input =".-757
437*.";
        assert_eq!("1194", part1(input));
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

    #[test_case(".111.\n.....", false; "not in range")]
    #[test_case("#111.\n.....", true; "next to hash")]
    #[test_case(".111.\n....#", true; "diagnol to hash")]
    #[test_case("...#\n...1.", true; "diagnol up right to hash")]
    #[test_case("*111.\n.....", true; "next to star")]
    #[test_case("+111.\n.....", true; "next to plus")]
    #[test_case("$111.\n.....", true; "next to dollar")]
    #[test_case("-111.\n.....", true; "next to dash")]
    #[test_case("467.......\n...*......", true; "real data")]
    fn test_in_range(input: &str, expected: bool) {
        let (points, parts) = read(input);
        assert_eq!(1, parts.len());
        assert_eq!(expected, parts[0].in_range(&points));
    }

}