package day4

import (
	"errors"
	"fmt"
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
func CountValidPassports(in []string) (int, error) {
	passports, err := parsePassports(in)
	if err != nil {
		fmt.Printf("Error parsing passports %v\n", err)
		return 0, err
	}
	var count int
	for _, p := range passports {
		if p.isValid() {
			count++
		}
	}
	return count, nil
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
