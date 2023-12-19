use std::collections::HashMap;
// use std::cmp;

fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

// Library functions

// Reused from day-10
type Col = i64;
type Row = i64;

fn str_to_coords_with_specials(input: &str, specials: Vec<char>) -> (HashMap<(Col, Row), char>, HashMap<char, (Col,Row)>, (Col,Row)) {
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
    let max = input.lines().enumerate().last().map(|(row, line)| ((line.len() - 1) as Col, row as Row)).unwrap_or((-1,-1));
    return (coord_map, special_map, max);
}

fn print_grid(grid: &HashMap<(i64,i64), char>, max: (Col, Row)) {
    println!("{}", grid_to_str(grid, max)); 
}

fn grid_to_str(grid: &HashMap<(i64,i64), char>, max: (Col, Row)) -> String {
    let rows:Vec<String> = (0..(max.1+1))
        .map( | row | (0..(max.0+1)).
            map(|col| grid.get(&(col, row))
            .unwrap_or(&' '))
            .collect())
        .collect();

    let res = rows.join("\n");
    return res;
}

// \Library

fn score(grid: &HashMap<(i64,i64), char>, max: (Col, Row)) -> i64 {
    let mut score = 0;
    for row in 0..(max.1+1) {
        let rocks_in_row = (0..(max.0+1))
            .map(|col| grid.get(&(col, row)))
            .filter(|c| c == &Some(&'O'))
            .count() as i64;
        let row_val = max.1 + 1 - row;
        score += rocks_in_row * row_val;
    }
    return score;
}


fn part2(input: &str) -> String {
    let (mut grid, _, max) = str_to_coords_with_specials(input, vec![]);

    let mut cache:HashMap<String, i64> = HashMap::new();
    let mut iter_score:HashMap<i64, i64> = HashMap::new();
    for i in 0..1000 as i64 { // 1000000000
        let grid_as_string = grid_to_str(&grid, max);
        if let Some(original_iter) = cache.get(&grid_as_string) {
            println!("Found loop! Iteration {} is the same as {}", i, original_iter);

            let guess = iter_score.get(& dbg!(dbg!(dbg!(1000000000 - original_iter) % dbg!(i - original_iter)) + original_iter));
            println!("Score for 1000000000 should be {}", guess.unwrap());

            // println!("{}:", i);
            // print_grid(&grid, max);
    
            return guess.unwrap().to_string();
        }

        cache.insert(grid_as_string, i);
        iter_score.insert(i, score(&grid, max));

        // println!("{}:", i);
        // print_grid(&grid, max);
        // println!("Score: {}", score(&grid, max));
        // println!();

        tilt_north(&mut grid, max);
        // print_grid(&grid, max);
        
        grid = rotate(&grid, max);
        // println!("rot 90");
        // print_grid(&grid, max);
        
        tilt_north(&mut grid, max);
        // println!("rot 90 and tilt");
        // print_grid(&grid, max);
    
        grid = rotate(&grid, max);
        // println!("rot 180");
        // print_grid(&grid, max);
        
        tilt_north(&mut grid, max);
        // println!("rot 180 and tilt");
        // print_grid(&grid, max);
    
        grid = rotate(&grid, max);
        // println!("rot 270");
        // print_grid(&grid, max);
    
        tilt_north(&mut grid, max);
        // println!("rot 270 and tilt");
        // print_grid(&grid, max);
    
        grid = rotate(&grid, max);
    }

    return score(&grid, max).to_string();
}

fn rotate(grid: &HashMap<(i64,i64), char>, max: (Col, Row)) -> HashMap<(i64, i64), char> {
    assert_eq!(max.0, max.1);
    let mut res: HashMap<(i64, i64), char> = HashMap::new();
    for ((orig_col, orig_row), v) in grid {
        res.insert((max.0 - orig_row, orig_col*1), v.clone());
    }
    return res;
}

fn tilt_north(grid: &mut HashMap<(i64,i64), char>, max: (Col, Row)) {
    for col in 0..(max.0 + 1) {
        for row in 0..(max.1 + 1) {
            match grid.get(&(col,row)) { // Looking for an empty space to pull a rock too
                Some('.') => {
                    // println!("Found an empty space at ({},{})", col, row);
                    for row2 in (row +1)..(max.1 +1) {
                        // println!("Checkint for rock at ({},{})", col, row2);
                        match grid.get(&(col,row2)) { // Looking for a rock to pull north
                            Some('.') => {}, // do nothing
                            Some('O') => {
                                grid.insert((col, row), 'O');
                                grid.insert((col, row2), '.');
                                break;
                            }, // do nothing
                            Some('#') => {
                                // if we've found a non-moving rock, there are no more rocks in this section to move.
                                // TODO - if there are no rocks, we can skip forward our search for open spaces as well.
                                break;
                            },
                            _ => panic!("err 2")
                        }
                    }
                }
                Some('O') => {}, // do nothing
                Some('#') => {}, // do nothing
                _ => panic!("err 1"),
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    // use test_case::test_case;

    #[test]
    fn test_example_1() {
        let actal = part2("\
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....");
        assert_eq!("64", actal);
    }
}