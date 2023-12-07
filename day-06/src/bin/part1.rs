fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

fn part1(input: &str) -> String {
    let res = dbg!(margin(42, 284) )
        * dbg!(margin(68, 1005))
        * dbg!(margin(69, 1122))
        * dbg!(margin(85, 1341));
    return res.to_string();
}

fn part2(_: &str) -> String {
    return margin(42686985, 284100511221341).to_string();
}

fn margin(time: i64, dist:i64) -> i64 {
    let mut min = time;
    for i in 1..time {
        if i * (time - i) > dist {
            return time + 1 - i * 2;
        }
    }
    return 0;
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test_case(7, 9, 4)]
    #[test_case(15, 40, 8)]
    #[test_case(30, 200, 9)]
    fn test_margin(time: i64, dist: i64, expected: i64) {
        assert_eq!(expected, margin(time, dist));
    }

    #[test]
    fn test_part1() {
        let res = margin(7,9) * margin(15, 40) * margin(30, 200);
        assert_eq!(288, res);
    }
}