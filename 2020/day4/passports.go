package day4

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ErrParse a parsing error
var ErrParse = errors.New("Parsing Error")

type passport struct {
	byr string //(Birth Year)
	iyr string // (Issue Year)
	eyr string // (Expiration Year)
	hgt string //(Height)
	hcl string //(Hair Color)
	ecl string //(Eye Color)
	pid string //(Passport ID)
	cid string //(Country ID)
}

func (p passport) isValid() bool {
	return p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func (p passport) isMoreValid() bool {
	return isNumberInRange(p.byr, 1920, 2002) &&
		isNumberInRange(p.iyr, 2010, 2020) &&
		isNumberInRange(p.eyr, 2020, 2030) &&
		isValidHeight(p.hgt) &&
		isValidColor(p.hcl) &&
		isValidEyeColor(p.ecl) &&
		isValidPID(p.pid)
}

func isNumberInRange(value string, min, max int) bool {
	val, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return min <= val && val <= max
}

var heightInCM = regexp.MustCompile(`([0-9]+)cm`)
var heightInInches = regexp.MustCompile(`([0-9]+)in`)

func isValidHeight(value string) bool {
	matchCm := heightInCM.FindStringSubmatch(value)
	if len(matchCm) >= 2 {
		return isNumberInRange(matchCm[1], 150, 193)
	}
	matchIn := heightInInches.FindStringSubmatch(value)
	if len(matchIn) >= 2 {
		return isNumberInRange(matchIn[1], 59, 76)
	}
	return false
}

var color = regexp.MustCompile(`#[0-9a-f]{6}`)

func isValidColor(value string) bool {
	return color.MatchString(value)
}

var eyeColors = map[string]struct{}{
	"amb": {},
	"blu": {},
	"brn": {},
	"gry": {},
	"grn": {},
	"hzl": {},
	"oth": {},
}

func isValidEyeColor(value string) bool {
	_, exists := eyeColors[value]
	return exists
}

var pid = regexp.MustCompile(`\d{9}`)

func isValidPID(value string) bool {
	return pid.MatchString(value) && len(value) == 9
}

func (p *passport) setField(nameAndValue string) error {
	nvArray := strings.Split(nameAndValue, ":")
	if len(nvArray) != 2 {
		return fmt.Errorf("error parsing key/value from %v", nameAndValue)
	}
	switch nvArray[0] {
	case "byr":
		p.byr = nvArray[1]
	case "iyr":
		p.iyr = nvArray[1]
	case "eyr":
		p.eyr = nvArray[1]
	case "hgt":
		p.hgt = nvArray[1]
	case "hcl":
		p.hcl = nvArray[1]
	case "ecl":
		p.ecl = nvArray[1]
	case "pid":
		p.pid = nvArray[1]
	case "cid":
		p.cid = nvArray[1]
	default:
		return fmt.Errorf("error parsing key from %v", nameAndValue)
	}
	return nil
}

// CountValidPassports counts valid passports.
func CountValidPassports(in []string) (int, int, error) {
	passports, err := parsePassports(in)
	if err != nil {
		return 0, 0, err
	}
	var count, count2 int
	for _, p := range passports {
		if p.isValid() {
			count++
		}
		if p.isMoreValid() {
			// fmt.Printf("valid: %+v\n", p)
			count2++
		}
	}
	return count, count2, nil
}

func parsePassports(in []string) ([]passport, error) {
	var passports []passport
	cur := &passport{}
	for _, s := range in {
		line := strings.Split(s, " ")
		if len(line[0]) == 0 {
			// fmt.Printf("Done with %v (valid? %v)\n", cur, cur.isValid())
			passports = append(passports, *cur)
			cur = &passport{}
		}
		for _, kv := range line {
			if kv == "" {
				continue
			}
			if err := cur.setField(kv); err != nil {
				return []passport{}, err
			}
		}
	}
	passports = append(passports, *cur)
	return passports, nil
}
