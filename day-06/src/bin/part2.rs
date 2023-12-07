fn main() {
    let input = include_str!("./input.txt");
    let output = part2(input);
    dbg!(output);
}

fn part2(input: &str) -> String {
    return Board::new_from(input).run().to_string();
}

struct Board {
    // 1994747387
    seeds: Vec<(i64, i64)>,
    seed_to_soil_map: Mapper,
    soil_to_fertilizer_map: Mapper,
    fertilizer_to_water_map: Mapper,
    water_to_light_map: Mapper,
    light_to_temperature_map: Mapper,
    temperature_to_humidity_map: Mapper, 
    humidity_to_location_map: Mapper,
}

impl Board {
    fn new_from(input: &str) -> Board {
        let mut sections = input.split("\n\n");
        let seeds_in = sections.next().unwrap();
        let seed_ranges:Vec<i64> = num_vec(seeds_in);
        let mut seeds:Vec<(i64,i64)> = Vec::new();
        for i in 0..seed_ranges.len() {
            if i % 2 == 0 {
                seeds.push((seed_ranges[i],seed_ranges[i+1]));
            }
        }

        // seed-to-soil map
        let seed_to_soil_map = Mapper::new_from(sections.next().unwrap());
        // soil-to-fertilizer map
        let soil_to_fertilizer_map = Mapper::new_from(sections.next().unwrap());
        // fertilizer-to-water map
        let fertilizer_to_water_map = Mapper::new_from(sections.next().unwrap());
        // water-to-light map
        let water_to_light_map = Mapper::new_from(sections.next().unwrap());
        // light-to-temperature map
        let light_to_temperature_map = Mapper::new_from(sections.next().unwrap());
        // termperature-to-humidity map
        let temperature_to_humidity_map = Mapper::new_from(sections.next().unwrap());
        // humidity-to-location map
        let humidity_to_location_map = Mapper::new_from(sections.next().unwrap());
        // SHOULD BE NONE
        dbg!(sections.next());
        return Board{
            seeds: seeds,
            seed_to_soil_map: seed_to_soil_map,
            soil_to_fertilizer_map: soil_to_fertilizer_map,
            fertilizer_to_water_map: fertilizer_to_water_map,
            water_to_light_map: water_to_light_map,
            light_to_temperature_map: light_to_temperature_map,
            temperature_to_humidity_map: temperature_to_humidity_map,
            humidity_to_location_map: humidity_to_location_map,
        };
    }

    fn map(&self, seed: i64) -> i64 {
        let soil = self.seed_to_soil_map.map(seed);
        let fertilizer = self.soil_to_fertilizer_map.map(soil);
        let water = self.fertilizer_to_water_map.map(fertilizer);
        let light = self.water_to_light_map.map(water);
        let temperature = self.light_to_temperature_map.map(light);
        let humidity = self.temperature_to_humidity_map.map(temperature);
        let location = self.humidity_to_location_map.map(humidity);
        return location;
    }

    fn rmap(&self, location: i64) -> i64 {
        let humidity = self.humidity_to_location_map.rmap(location);
        let temperature = self.temperature_to_humidity_map.rmap(humidity);
        let light = self.light_to_temperature_map.rmap(temperature);
        let water = self.water_to_light_map.rmap(light);
        let fertilizer = self.fertilizer_to_water_map.rmap(water);
        let soil = self.soil_to_fertilizer_map.rmap(fertilizer);
        let seed = self.seed_to_soil_map.rmap(soil);
        return seed;
    }

    fn run(&self) -> i64 {
        let max = self.seeds.iter().map(|p| p.0+p.1).max();
        dbg!(max);
        for i in 0..max.unwrap() {
            let maybe = self.rmap(i);
            for (start, range) in &self.seeds {
                // dbg!(i, maybe, start, range);
                if maybe >= *start && maybe < start+range {
                    return i;
                }
            }
        }
        return -1;
    }
}

struct Mapper {
    mappings: Vec<Mapping>
}

impl Mapper {
    fn new_from(input: &str) -> Mapper {
        return Mapper {
            mappings: input.lines().skip(1).map(|l| Mapping::new_from(l)).collect(),
        };
    }

    fn map(&self, src: i64) -> i64 {
        for mapper in &self.mappings {
            let maybe = mapper.map(src);
            if maybe.is_some() {
                return maybe.unwrap();
            }
        }
        return src;
    }

    fn rmap(&self, src: i64) -> i64 {
        for mapper in &self.mappings {
            let maybe = mapper.rmap(src);
            if maybe.is_some() {
                return maybe.unwrap();
            }
        }
        return src;
    }

}

#[derive(Debug)]
struct Mapping {
    dest_range_start: i64,
    src_range_start: i64,
    range_length: i64,
}

impl Mapping {
    fn new_from(input: &str) -> Mapping {
        let nums = num_vec(input);
        if nums.len() == 3 {
            return Mapping {
                dest_range_start: nums[0],
                src_range_start: nums[1],
                range_length: nums[2],
            };
        }
        dbg!(input);
        panic!("Invalid input")
    }

    fn map(&self, src: i64) -> Option<i64> {
        let delta = src - self.src_range_start;
        if delta >= 0 && delta < self.range_length {
            return Some(self.dest_range_start + delta);
        }
        return None;
    }
    
    fn rmap(&self, src: i64) -> Option<i64> {
        let delta = src - self.dest_range_start;
        if delta >= 0 && delta < self.range_length {
            return Some(self.src_range_start + delta);
        }
        return None;
    }
}

fn num_vec(input: &str) -> Vec<i64> {
    let results: Vec<i64> = input.split(' ')
        .map(|s| s.parse::<i64>())
        .filter(|r| r.is_ok())
        .map(|r| r.unwrap())
        .collect();
    return results;
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    static TEST_INPUT: &str = "seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4";

    #[test]
    fn test_it_works() {
        assert_eq!("46", part2(TEST_INPUT))
    }

    #[test]
    fn test_seeds(){
        let b = Board::new_from(TEST_INPUT);
  //       assert_eq!(HashSet![79,80,81,82,83,84,85,86,87,88,89,90,91,92,55,56,57,58,59,60,61,62,63,64,65,66,67], b.seeds);
    }

    #[test_case(97, None)]
    #[test_case(98, Some(50))]
    #[test_case(99, Some(51))]
    #[test_case(100, None)]
    fn test_mapping(input: i64, expected: Option<i64>) {
        let mapping = Mapping::new_from("50 98 2");
        assert_eq!(expected, mapping.map(input))
    }

    #[test_case(49, None)]
    #[test_case(50, Some(98))]
    #[test_case(51, Some(99))]
    #[test_case(52, None)]
    fn test_rmapping(input: i64, expected: Option<i64>) {
        let mapping = Mapping::new_from("50 98 2");
        assert_eq!(expected, mapping.rmap(input))
    }

    #[test_case(49, 49)]
    #[test_case(50, 52)]
    #[test_case(51, 53)]
    #[test_case(96, 98)]
    #[test_case(97, 99)]
    #[test_case(98, 50)]
    #[test_case(99, 51)]
    #[test_case(100, 100)]
    fn test_mapper(input: i64, expected: i64) {
        let mapper = Mapper::new_from("seed-to-soil map:
50 98 2
52 50 48");
        assert_eq!(expected, mapper.map(input))
    }

    #[test_case(49, 49)]
    #[test_case(50, 52)]
    #[test_case(51, 53)]
    #[test_case(96, 98)]
    #[test_case(97, 99)]
    #[test_case(98, 50)]
    #[test_case(99, 51)]
    #[test_case(100, 100)]
    fn test_rmapper(seed: i64, location: i64) {
        let mapper = Mapper::new_from("seed-to-soil map:
50 98 2
52 50 48");
        assert_eq!(seed, mapper.rmap(location))
    }

    #[test_case(79, 82)]
    #[test_case(14, 43)]
    #[test_case(55, 86)]
    #[test_case(13, 35)]
    fn test_board_map(seed: i64, expected:i64) {
        let b = Board::new_from(TEST_INPUT);
        assert_eq!(expected, b.map(seed))
    }

    #[test_case(79, 82)]
    #[test_case(14, 43)]
    #[test_case(55, 86)]
    #[test_case(13, 35)]
    fn test_board_rmap(seed: i64, location:i64) {
        let b = Board::new_from(TEST_INPUT);
        assert_eq!(seed, b.rmap(location))
    }

    #[test]
    fn test_num_vec() {
        let input = " 41 48 83 86 17 ";
        assert_eq!(vec![41, 48, 83, 86, 17], num_vec(input));
    }

}