use std::collections::BTreeMap;

fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

fn part2(input: &str) -> String {
    let hands_and_bids = read_input_to_map(input);
    return winnings(hands_and_bids).to_string();
}

fn read_input_to_map(input: &str) -> BTreeMap<i64,i64> {
    let mut hands_and_bids:BTreeMap<i64,i64> = BTreeMap::new();
    for line in input.lines(){
        let mut iter = line.split(" ");
        let hand = iter.next().unwrap();
        let bid = iter.next().unwrap().parse::<i64>().unwrap();
        hands_and_bids.insert(cards_to_int(hand), bid);
    }
    return dbg!(hands_and_bids);
}

fn cards_to_int(cards: &str) -> i64 {
    let card_val = cards.chars()
        .map(|c| card_to_num(c))
        .fold(0, |acc, x| acc * 100 + x);
    return 10000000000 * cards_to_hand(cards) + card_val;
}

// 7 = five of a kind
// 6 = four of a kind
// 5 = full house
// 4 = three of a kind
// 3 = two pair
// 2 = one pair
// 1 = highest
fn cards_to_hand(cards: &str) -> i64 {
    let mut best = 1; // highest card
    for c in ['2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A', 'J'] {
        let count = cards.chars().filter(|card| card == &c || card == &'J').count();
        match count {
            5 => best = 7, // 5 of a kind
            4 => {
                best = biggest(6, best); // four of a kind
            },
            3 => {
                let remaining = cards.chars().filter(|card| card != &c && card != &'J').collect::<String>();
                match cards_to_hand(&remaining) {
                    2 => best = biggest(5, best), // full house
                    _ => best = biggest(4,best), // three of a kind
                }
            },
            2 => {
                let remaining = cards.chars().filter(|card| card != &c && card != &'J').collect::<String>();
                match cards_to_hand(&remaining) {
                    4 => best = biggest(5, best), // full house
                    2 => best = biggest(3, best), // two pair
                    _ => best = biggest(2, best), // one pair
                }
            },
            _ => continue
        }
    }
    return best;
}

fn biggest(a:i64, b:i64) -> i64 {
    if a > b {
        return a;
    }
    return b;
}

fn card_to_num(card: char) -> i64 {
    match card {
        'T' => return 10,
        'J' => return 1,
        'Q' => return 12,
        'K' => return 13,
        'A' => return 14,
        _ => return i64::from(card.to_digit(10).unwrap_or(0))
    }
}

fn winnings(hands_and_bids: BTreeMap<i64, i64>) -> i64 {
    let mut rank_val:i64 = 1;
    let mut result:i64 = 0;
    for (hand, bid) in hands_and_bids.iter() {
        // dbg!(hand, bid);
        _ = hand;
        result += bid * rank_val;
        rank_val += 1;
    }

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn test_it_works() {
        let test_input = "32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483";
        assert_eq!("5905", part2(test_input));
    }

    #[test_case("23456", 10203040506)]
    #[test_case("AKQJT", 11413121110)]
    #[test_case("32T3K", 20302100313)]
    fn test_cards_to_int(cards: &str, expected: i64) {
        assert_eq!(expected, cards_to_int(cards));
    }

    #[test]
    fn test_winnings() {
        let mut hands_and_bids:BTreeMap<i64, i64> = BTreeMap::new();
        hands_and_bids.insert(10, 1);
        hands_and_bids.insert(5, 5);
        hands_and_bids.insert(1, 2);
        assert_eq!(2*1 + 5*2 + 1*3, winnings(hands_and_bids));
    }

    #[test_case('2', 2)]
    #[test_case('3', 3)]
    #[test_case('8', 8)]
    #[test_case('9', 9)]
    #[test_case('T', 10)]
    #[test_case('J', 1)]
    #[test_case('Q', 12)]
    #[test_case('K', 13)]
    #[test_case('A', 14)]
    fn test_card_to_number(card: char, expected: i64) {
        assert_eq!(expected, card_to_num(card));
    }

    #[test_case("AAAAA", 7)]
    #[test_case("AAAAK", 6)]
    #[test_case("AAAKK", 5)]
    #[test_case("AAKKK", 5)]
    #[test_case("AAAKQ", 4)]
    #[test_case("AAKKQ", 3)]
    #[test_case("AAKQT", 2)]
    #[test_case("A2345", 1)]
    #[test_case("32T3K", 2)]
    #[test_case("T55J5", 6)]
    #[test_case("KK677", 3)]
    #[test_case("KTJJT", 6)]
    #[test_case("QQQJA", 6)]
    fn test_cards_to_hand(cards: &str, expected: i64) {
        assert_eq!(expected, cards_to_hand(cards));
    }
}