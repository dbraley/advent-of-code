fn main() {
    let input = include_str!("./input.txt");
    let output = part1(input);
    dbg!(output);
}

fn part1(input: &str) -> String {
    return input.lines()
        .map(|line| value(line))
        .inspect(|v| _ = dbg!(v))
        .fold(0, |acc, v| acc + v)
        .to_string();
}

fn value(input: &str) -> u32 {
    // Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
    let segs = input.split(':');
    let mut vals = segs.last().unwrap().split('|');
    let win_vals = num_vec(vals.next().unwrap());
    let card_vals = num_vec(vals.next().unwrap());
    let mut points = 0;
    for v in win_vals {
        if card_vals.contains(&v) {
            match points {
                0 => points += 1,
                _ => points *= 2,
            }
        }
    }
    return points;
}

fn num_vec(input: &str) -> Vec<u32> {
    let results: Vec<u32> = input.split(' ')
        .map(|s| s.parse::<u32>())
        .filter(|r| r.is_ok())
        .map(|r| r.unwrap())
        .collect();
    return results;
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_it_works() {
        let input = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11";
        assert_eq!("13", part1(input))
    }

    #[test_case("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8; "Card 1")]
    fn test_value(input: &str, expected: u32) {
        assert_eq!(expected, value(input));
    }

    #[test]
    fn test_num_vec() {
        let input = " 41 48 83 86 17 ";
        assert_eq!(vec![41, 48, 83, 86, 17], num_vec(input));
    }

}