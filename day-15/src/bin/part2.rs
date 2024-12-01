use lazy_static::lazy_static;
use regex::Regex;

lazy_static! {
    static ref PUT_LENS:Regex = Regex::new(r"(.*)=(\d)$").unwrap();
    static ref POP_LENS:Regex = Regex::new(r"(.*)-$").unwrap();
}

fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

#[derive(Clone, Debug)]
struct Lens {
    label: String,
    focal_length: u8,
}

fn part2(input: &str) -> String {
    let mut boxes: Vec<Vec<Lens>> = vec![vec![]; 256];
    for s in input.split(",") {
        if let Some(captures) = PUT_LENS.captures(s) {
            let l = Lens{
                label: captures[1].to_string(),
                focal_length: captures[2].parse().unwrap(),
            };
            let mut added = false;
            for b in boxes[hash_str(l.label.as_str())].iter_mut() {
                if b.label == l.label {
                    b.focal_length = l.focal_length;
                    added = true;
                    break;
                }
            }
            if !added {
                boxes[hash_str(l.label.as_str())].push(l);
            }
        }
        if let Some(captures) = POP_LENS.captures(s) {
            let label = captures[1].to_string();
            boxes[hash_str(&label)].retain(|l| l.label != label);
        }
    }
    let score = score_boxes(&boxes);
    // dbg!(boxes);
    return score.to_string();
}

fn score_boxes(boxes: &Vec<Vec<Lens>>) -> i64 {
    let mut score = 0;
    for (bnum, b) in boxes.iter().enumerate() {
        for (lnum, l) in b.iter().enumerate() {
            score += (bnum as i64 + 1) * (lnum as i64 + 1) * (l.focal_length as i64);
        }
    }
    return score;
}

fn hash_str(input: &str) -> usize {
    let ret = input.chars()
        .map(|c| c as usize)
        .fold(0, |acc, x| ((acc + x) * 17) % 256);
        
    return ret;
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_example_1() {
        let actal = part2("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7");
        assert_eq!("145", actal);
    }

    #[test_case("HASH", 52)]
    #[test_case("rn", 0)]
    #[test_case("cm", 0)]
    #[test_case("qp", 0)]
    fn test_hash(input: &str, expected: usize) {
        assert_eq!(expected, hash_str(input));
    }
}