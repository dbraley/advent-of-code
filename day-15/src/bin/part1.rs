fn main() {
    let input = include_str!("./input.txt");
    let output = part1(input);
    dbg!(output);
}

fn part1(input: &str) -> String {
    let score:i64 = input.split(",").map(|s| hash_str(s)).sum();
    return score.to_string();
}

fn hash_str(input: &str) -> i64 {
    let ret = input.chars()
        .map(|c| c as i64)
        .fold(0, |acc, x| dbg!(dbg!(acc + x) * 17) % 256);
        
    return ret;
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_example_1() {
        let actal = part1("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7");
        assert_eq!("1320", actal);
    }

    #[test_case("HASH", 52)]
    #[test_case("rn=1", 30)]
    fn test_hash(input: &str, expected: i64) {
        assert_eq!(expected, hash_str(input));
    }
}