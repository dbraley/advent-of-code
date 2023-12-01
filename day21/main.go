package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Example:")
	ept1, ept2 := day20("example.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ept1, ept2)

	fmt.Println("Input:")
	ipt1, ipt2 := day20("input.txt")
	fmt.Printf("  Pt1: %v\n  Pt2: %v\n", ipt1, ipt2)
}

func day20(fileName string) (pt1, pt2 int) {
	input, err := file.ReadFile("day21/" + fileName)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	mathMonkeyMap := createMonkeyMap(input)
	fmt.Println("len check", len(input), len(mathMonkeyMap))

	rootMonkey := mathMonkeyMap["root"]
	pt1 = rootMonkey.Answer(mathMonkeyMap)

	// we can figure this out pretty easily if there's only one use of humn!!!
	//fmt.Println(rootMonkey.Uses("humn", mathMonkeyMap))
	// Cool, both scenarios respond 1

	rootMonkey.mathType = "=="
	pt2 = rootMonkey.FindHumanAnswer(0, mathMonkeyMap)
	return pt1, pt2
}

func createMonkeyMap(input []string) map[string]mathMonkey {
	a := regexp.MustCompile(`(\S{4}): (-?\d+)`)
	m := regexp.MustCompile(`(\S{4}): (\S{4}) ([+\-*/]) (\S{4})`)
	mathMonkeyMap := make(map[string]mathMonkey)
	for _, line := range input {
		answerMonkeyBlobs := a.FindStringSubmatch(line)
		if len(answerMonkeyBlobs) != 0 {
			answer, _ := strconv.Atoi(answerMonkeyBlobs[2])
			mathMonkeyMap[answerMonkeyBlobs[1]] = mathMonkey{
				name:        answerMonkeyBlobs[1],
				knowsAnswer: true,
				answer:      answer,
				mathType:    "",
				m1:          "",
				m2:          "",
			}

		} else {
			mathMonkeyBlobs := m.FindStringSubmatch(line)
			if len(mathMonkeyBlobs) != 0 {
				mathMonkeyMap[mathMonkeyBlobs[1]] = mathMonkey{
					name:        mathMonkeyBlobs[1],
					knowsAnswer: false,
					answer:      0,
					mathType:    mathMonkeyBlobs[3],
					m1:          mathMonkeyBlobs[2],
					m2:          mathMonkeyBlobs[4],
				}
			}
		}
	}
	//for k, v := range mathMonkeyMap {
	//	fmt.Println(k, v)
	//}
	return mathMonkeyMap
}

type mathMonkey struct {
	name        string
	knowsAnswer bool
	answer      int
	mathType    string
	m1          string
	m2          string
}

func (m mathMonkey) Answer(mm map[string]mathMonkey) int {
	if m.knowsAnswer {
		return m.answer
	} else {
		switch m.mathType {
		case "+":
			return mm[m.m1].Answer(mm) + mm[m.m2].Answer(mm)
		case "-":
			return mm[m.m1].Answer(mm) - mm[m.m2].Answer(mm)
		case "*":
			return mm[m.m1].Answer(mm) * mm[m.m2].Answer(mm)
		case "/":
			return mm[m.m1].Answer(mm) / mm[m.m2].Answer(mm)
		default:
			panic(fmt.Sprintf("Don't know how to perform %v: %v", m.mathType, m))
		}
	}
}

func (m mathMonkey) Uses(name string, mm map[string]mathMonkey) int {
	if m.name == name {
		return 1
	}
	if m.knowsAnswer {
		return 0
	}
	return mm[m.m1].Uses(name, mm) + mm[m.m2].Uses(name, mm)
}

func (m mathMonkey) FindHumanAnswer(e int, mm map[string]mathMonkey) int {
	if m.name == "humn" {
		return e
	}
	if m.knowsAnswer {
		panic("Shouldn't have gotten here!")
	}
	monkey1 := mm[m.m1]
	monkey2 := mm[m.m2]
	switch m.mathType {
	case "+":
		if monkey1.Uses("humn", mm) > 0 {
			return monkey1.FindHumanAnswer(e-monkey2.Answer(mm), mm)
		} else {
			return monkey2.FindHumanAnswer(e-monkey1.Answer(mm), mm)
		}
	case "-":
		if monkey1.Uses("humn", mm) > 0 {
			// e = humn - m2answer => humn = e + m2answer
			return monkey1.FindHumanAnswer(e+monkey2.Answer(mm), mm)
		} else {
			// e = m1answer - humn => humn = m1answer - e
			return monkey2.FindHumanAnswer(monkey1.Answer(mm)-e, mm)
		}
	case "*":
		if monkey1.Uses("humn", mm) > 0 {
			// e = humn + m2answer => humn = e / m2answer
			return monkey1.FindHumanAnswer(e/monkey2.Answer(mm), mm)
		} else {
			return monkey2.FindHumanAnswer(e/monkey1.Answer(mm), mm)
		}
	case "/":
		if monkey1.Uses("humn", mm) > 0 {
			// e = humn / m2answer => m1.answer = e*m2answer
			return monkey1.FindHumanAnswer(e*monkey2.Answer(mm), mm)
		} else {
			// e = m1answer / humn => m2.answer = m1answer/e
			return monkey2.FindHumanAnswer(monkey1.Answer(mm)/e, mm)
		}
	case "==":
		if monkey1.Uses("humn", mm) > 0 {
			return monkey1.FindHumanAnswer(monkey2.Answer(mm), mm)
		} else {
			return monkey2.FindHumanAnswer(monkey1.Answer(mm), mm)
		}
		// Doesn't happen
		return 0
	default:
		panic(fmt.Sprintf("Don't know how to perform %v: %v", m.mathType, m))
	}
}
