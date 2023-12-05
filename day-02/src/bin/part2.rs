fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

fn part2(input: &str) -> String {
    let sum = input.lines().map(|line| line_value(line)).fold(0, |acc, x| acc + x);
    return sum.to_string();
}

fn line_value(line: &str) -> u32 {
    let (_, sets) = game_number_and_sets(line);
    return max_red(sets) * max_green(sets) * max_blue(sets)
}

fn game_number_and_sets(line: &str) -> (u32, &str) {
    let mut splitter = line.split(':');
    let game_str:String = splitter.next().unwrap().chars().skip(5).collect();
    let sets_str = splitter.next().unwrap();
    return (game_str.parse::<u32>().unwrap(), sets_str);
}

fn max_red(line: &str) -> u32 {
    return line.replace(",", ";").split(";")
        .filter(|s| s.contains("red"))
        .map(|s| s.trim().chars().take_while(|c| c.is_numeric()).collect::<String>().parse::<u32>().unwrap())
        .max().unwrap_or(0)
}

fn max_blue(line: &str) -> u32 {
    return line.replace(",", ";").split(";")
        .filter(|s| s.contains("blue"))
        .map(|s| s.trim().chars().take_while(|c| c.is_numeric()).collect::<String>().parse::<u32>().unwrap())
        .max().unwrap_or(0)
}

fn max_green(line: &str) -> u32 {
    return line.replace(",", ";").split(";")
        .filter(|s| s.contains("green"))
        .map(|s| s.trim().chars().take_while(|c| c.is_numeric()).collect::<String>().parse::<u32>().unwrap())
        .max().unwrap_or(0)
}


#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test]
    fn it_works() {
        let result = part2("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green");
        assert_eq!(result, "2286")
    }

    #[test_case("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48)]
    #[test_case("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12)]
    #[test_case("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560)]
    #[test_case("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630)]
    #[test_case("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36)]
    fn line_test(input: &str, expected: u32) {
        assert_eq!(expected, line_value(input))
    }

    #[test_case("3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 4)]
    #[test_case("3 blue; 2 green, 6 blue; 2 green", 0)]
    fn red_test(input: &str, expected: u32) {
        assert_eq!(expected, max_red(input));
    }

    #[test_case("3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 6)]
    #[test_case("4 red; 1 red, 2 green; 2 green", 0)]
    fn blue_test(input: &str, expected: u32) {
        assert_eq!(expected, max_blue(input));
    }

}