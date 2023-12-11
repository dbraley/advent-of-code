// use std::collections::HashMap;

// use regex::Regex;

use std::collections::HashMap;

fn main() {
    let input = include_str!("./input.txt");
    let output = part1(input);
    dbg!(output);
}

type Col = i64;
type Row = i64;
type Point = (Col, Row);

fn north(p: Point) -> Point {
    (p.0, p.1 - 1)
}
fn east(p: Point) -> Point {
    (p.0+1, p.1)
}
fn west(p: Point) -> Point {
    (p.0 - 1, p.1)
}
fn south(p: Point) -> Point {
    (p.0, p.1 + 1)
}

trait Pipe {
    fn can_go(self, dir: char) -> bool;
    fn next(self, not: char) -> Option<char>;
}

impl Pipe for char {
    fn can_go(self, dir: char) -> bool {
        match dir {
            'N' => self == '|' || self == 'L' || self == 'J',
            'E' => self == '-' || self == 'L' || self == 'F',
            'W' => self == '-' || self == 'J' || self == '7',
            'S' => self == '|' || self == '7' || self == 'F',
            _ => false,
        }
    }
    fn next(self, not_this: char) -> Option<char> {
        match self {
            '|' => opposite('N', 'S', not_this),
            '-' => dbg!(opposite('E', 'W', not_this)),
            'L' => opposite('N', 'E', not_this),
            'J' => opposite('N', 'W', not_this),
            '7' => opposite('W', 'S', not_this),
            'F' => opposite('E', 'S', not_this),
            _ => None,
        }
    }
}

fn opposite(c1: char, c2: char, not_this: char) -> Option<char> {
    if not_this == c1 {
        return Some(c2);
    } 
    if not_this == c2 {
        return Some(c1);
    }
    return None;
}

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

fn part1(input: &str) -> String {
    let (coords, specials) = str_to_coords_with_specials(input, vec!['S']);
    let start = specials.get(&'S').unwrap();
    // let n = coords.get(&north(*start)).unwrap_or(&'.');
    let e = coords.get(&east(*start)).unwrap_or(&'.');
    if e.can_go('W') {
        let mut coord = east(*start);
        let mut from = 'W';
        for i in 1..20000 {
            let next_pipe = coords.get(&coord).unwrap();
            if next_pipe == &'S' {
                return (i/2).to_string();
            }
            match next_pipe.next(from) {
                Some('S') => {
                    coord = south(coord);
                    from = 'N';
                },
                Some('W') => {
                    coord = west(coord);
                    from = 'E';
                },
                Some('E') => {
                    coord = east(coord);
                    from = 'W';
                },
                Some('N') => {
                    coord = north(coord);
                    from = 'S';
                },
                _ => panic!()
            }
        }
    }
    let res = "error - need more iterations";
    return res.to_string();
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_example_1() {
        let actal = part1("\
-L|F7
7S-7|
L|7||
-L-J|
L|-JF");
        assert_eq!("4", actal);
    }

    #[test]
    fn test_example_2() {
        let actal = part1("\
7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ");
        assert_eq!("8", actal);
    }

    #[test]
    fn test_str_to_coords_with_specials() {
        let input = "\
S7
LJ";
        let (coords, specials) = str_to_coords_with_specials(input, vec!['S', 'Q']);
        
        assert_eq!(4,coords.len());
        assert_eq!(Some(&'S'), coords.get(&((0 as Col,0 as Row) as Point)));
        assert_eq!(Some(&'7'), coords.get(&((1 as Col,0 as Row) as Point)));
        assert_eq!(Some(&'L'), coords.get(&((0 as Col,1 as Row) as Point)));
        assert_eq!(Some(&'J'), coords.get(&((1 as Col,1 as Row) as Point)));
        
        assert_eq!(1,specials.len());
        assert_eq!(Some(&(0 as Col, 0 as Row)), specials.get(&'S'));
        assert_eq!(None, specials.get(&'Q'));
    }

    #[test_case('|', 'N', true)]    
    #[test_case('|', 'S', true)] 
    #[test_case('|', 'E', false)] 
    #[test_case('|', 'W', false)] 
    #[test_case('-', 'E', true)] 
    #[test_case('-', 'W', true)] 
    #[test_case('L', 'N', true)] 
    #[test_case('L', 'E', true)] 
    #[test_case('J', 'N', true)] 
    #[test_case('J', 'W', true)] 
    #[test_case('7', 'W', true)] 
    #[test_case('7', 'S', true)]
    #[test_case('F', 'E', true)]
    #[test_case('F', 'S', true)]
    fn test_next_val(pipe: char, dir: char, expected: bool) {
        assert_eq!(expected, pipe.can_go(dir));
    }

}