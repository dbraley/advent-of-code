// use std::collections::HashMap;

// use regex::Regex;

use std::collections::HashMap;

fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
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
            '-' => opposite('E', 'W', not_this),
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

fn part2(input: &str) -> String {
    let (coords, specials) = str_to_coords_with_specials(input, vec!['S']);
    let start = specials.get(&'S').unwrap();
    let path = get_path(coords, *start);

    // States [out, in]
    let mut state = "out";
    let mut boundary:Option<char> = None;
    let mut in_count = 0;
    for (row, line) in input.lines().enumerate() {
        let mut row_str = String::new();
        for (col, c) in line.chars().enumerate() {
            match path.get(&(col as Col, row as Row)) {
                Some(c) => {
                    row_str += c.to_string().as_str();
                    match state {
                        "out" => {
                            match c {
                                '|' => {state = "in"},
                                'F' => {boundary = Some('F');},
                                'L' => {boundary = Some('L');}
                                '7' => {
                                    match boundary {
                                        Some('F') => {boundary = None;}, // Where out, still out
                                        Some('L') => {
                                            boundary = None;
                                            state = "in";
                                        },
                                        _ => panic!("What?"),
                                    }
                                },
                                'J' => {
                                    match boundary {
                                        Some('L') => {boundary = None;}, // Where out, still out
                                        Some('F') => {
                                            boundary = None;
                                            state = "in";
                                        },
                                        _ => panic!("What?"),
                                    }
                                },
                                _ => (),
                            }
                        },
                        "in" => {
                            match c {
                                '|' => {state = "out"},
                                'F' => {boundary = Some('F');},
                                'L' => {boundary = Some('L');},
                                '7' => {
                                    match boundary {
                                        Some('F') => {boundary = None;}, // Where in, still in
                                        Some('L') => {
                                            boundary = None;
                                            state = "out";
                                        },
                                        _ => panic!("What?"),
                                    }
                                },
                                'J' => {
                                    match boundary {
                                        Some('L') => {boundary = None;}, // Where in, still in
                                        Some('F') => {
                                            boundary = None;
                                            state = "out";
                                        },
                                        _ => panic!("What?"),
                                    }
                                },
                                _ => (),
                            }
                        },
                        _ => panic!("what?"),
                    }
                },
                None => {
                    match state {
                        "out" => {
                            row_str += "O";
                        },
                        "in" => {
                            row_str += "I";
                            in_count += 1;
                        },
                        _ => panic!("bad state"),
                    }
                },
            }
        }
        println!("{}", row_str);
    }

    return in_count.to_string();
}

fn get_path(coords: HashMap<Point,char>, start: Point) -> HashMap<Point, char> {
    let e = coords.get(&east(start)).unwrap_or(&'.');
    let mut path: HashMap<Point,char> = HashMap::new();
    path.insert(start, *coords.get(&start).unwrap());
    if e.can_go('W') {
        let mut coord = east(start);
        let mut from = 'W';
        for i in 1..20000 {
            let next_pipe = coords.get(&coord).unwrap();
            if next_pipe == &'S' {
                match from {
                    'S' => _ = path.insert(coord, 'F'),
                    'W' => _ = path.insert(coord, '-'),
                    'N' => _ = path.insert(coord, 'L'),
                    _ => (),
                }
                return path;
            }
            path.insert(coord, *coords.get(&coord).unwrap());
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
    panic!("Couldn't go east from start");
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_example_1() {
        let actal = part2("\
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........");
        assert_eq!("4", actal);
    }

    #[test]
    fn test_example_2() {
        let actal = part2("\
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
");
        assert_eq!("8", actal);
    }

    #[test]
    fn test_example_3() {
        let actal = part2("\
FF7S7F7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L");
        assert_eq!("10", actal);
    }

    #[test]
    fn test_get_path() {
        let input = "\
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........";
        let (coords,_) = str_to_coords_with_specials(input, vec!['S']);
        let path = get_path(coords, (1,1));
        assert_eq!(46, path.len());
        assert_eq!(true, path.contains_key(&(1,1)));
        assert_eq!(true, path.contains_key(&(2,1)));
        assert_eq!(true, path.contains_key(&(4,5)));
        assert_eq!(false, path.contains_key(&(0,0)));
        assert_eq!(false, path.contains_key(&(4,4)));
        assert_eq!(false, path.contains_key(&(7,6)));
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