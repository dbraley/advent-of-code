package main

import (
	"fmt"
)
import "math/big"

type Monkey struct {
	Items   []big.Int
	Inspect func(big.Int) big.Int
	Test    func(big.Int) int
}

func (m *Monkey) Add(item int64) {
	m.Items = append(m.Items, *big.NewInt(item))
}

func (m *Monkey) Catch(item big.Int) {
	m.Items = append(m.Items, item)
}

func main() {
	//f := "example.txt"
	f := "input.txt"

	//input, err := file.ReadFile("day11/" + f)
	//if err != nil {
	//	fmt.Printf("Error reading file %v\n", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Printf("%d\n", len(input))

	monkeyCount := make(map[int]int64)

	monkeys, relief := getMonkeys(f)
	inspectAt := []int{1, 20, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000}
	maxRound := 10000
	//maxRound := 20
	for round := 0; round < maxRound; round++ {
		if contains(inspectAt, round) {
			printMonkeyCounts(round, monkeyCount)
		}
		for n := 0; n < len(monkeys); n++ {
			monkey := monkeys[n]
			// fmt.Printf("Monkey %d:\n", n)
			for _, item := range monkey.Items {
				// fmt.Printf("\tMonkey inspects an item with a worry level of %v.\n", item.Text(10))
				monkeyCount[n]++
				newVal := monkey.Inspect(item)
				newVal = *newVal.Mod(&newVal, big.NewInt(relief))
				// fmt.Printf("\t\tMonkey gets bored with item. Worry level is now %v\n", newVal.Text(10))
				newMonkey := monkey.Test(newVal)
				// fmt.Printf("\t\tItem with worry level %v is thrown to monkey %d.\n", newVal.Text(10), newMonkey)
				monkeys[newMonkey].Catch(newVal)
			}
			monkey.Items = []big.Int{}
		}
		// fmt.Printf("After round %d, the monkeys are holding items with these worry levels:\n", round)
		for n := 0; n < len(monkeys); n++ {
			monkey := monkeys[n]
			var text []string
			for _, item := range monkey.Items {
				text = append(text, item.Text(10))
			}
			// fmt.Printf("Monkey %d: %v\n", n, strings.Join(text, ", "))
		}
		//fmt.Println()
	}

	printMonkeyCounts(maxRound, monkeyCount)
}

func printMonkeyCounts(round int, monkeyCount map[int]int64) {
	var max0 int64
	var max1 int64
	fmt.Printf("== After round %d ==\n", round)
	for i := 0; i < len(monkeyCount); i++ {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, monkeyCount[i])
		if monkeyCount[i] > max0 {
			max1 = max0
			max0 = monkeyCount[i]
		} else if monkeyCount[i] > max1 {
			max1 = monkeyCount[i]
		}
	}
	fmt.Printf("Max: %d\nNext: %d\nScore: %d", max0, max1, max0*max1)
	fmt.Println()
}

func toBigArray(ints []int64) []big.Int {
	var r []big.Int
	for _, i := range ints {
		r = append(r, *big.NewInt(i))
	}
	return r
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func makeMultiplyInspectMethod(mult int64) func(big.Int) big.Int {
	return func(x big.Int) big.Int {
		newVal := x.Mul(&x, big.NewInt(mult))
		// fmt.Printf("\t\tWorry level is multiplied by %d to %v.\n", mult, newVal)
		return *newVal
	}
}

func makeSquareInspectMethod() func(big.Int) big.Int {
	return func(x big.Int) big.Int {
		newVal := x.Mul(&x, &x)
		// fmt.Printf("\t\tWorry level is multiplied by itself to %v.\n", newVal)
		return *newVal
	}
}

func makeAddInspectMethod(add int64) func(big.Int) big.Int {
	return func(x big.Int) big.Int {
		plus := big.NewInt(add)
		newVal := x.Add(&x, plus)
		// fmt.Printf("\t\tWorry level is increases by %d to %v.\n", add, newVal.Text(10))
		return *newVal
	}
}

func makeTestMethod(mod int64, pass, fail int) func(x big.Int) int {
	return func(x big.Int) int {
		if big.NewInt(0).Mod(&x, big.NewInt(mod)).Cmp(big.NewInt(0)) == 0 {
			// fmt.Printf("\t\tCurrent worry level is divisible by %d\n", mod)
			return pass
		}
		// fmt.Printf("\t\tCurrent worry level is not divisible by %d\n", mod)
		return fail
	}
}

func getMonkeys(f string) (map[int]*Monkey, int64) {
	monkeys := make(map[int]*Monkey)
	relief := 1
	if f == "example.txt" {
		monkeys[0] = &Monkey{
			Inspect: makeMultiplyInspectMethod(19),
			Test:    makeTestMethod(23, 2, 3),
		}
		relief *= 23
		monkeys[0].Add(79)
		monkeys[0].Add(98)

		monkeys[1] = &Monkey{
			Items:   toBigArray([]int64{54, 65, 75, 74}),
			Inspect: makeAddInspectMethod(6),
			Test:    makeTestMethod(19, 2, 0),
		}
		relief *= 19

		monkeys[2] = &Monkey{
			Items:   toBigArray([]int64{79, 60, 97}),
			Inspect: makeSquareInspectMethod(),
			Test:    makeTestMethod(13, 1, 3),
		}
		relief *= 13
		monkeys[3] = &Monkey{
			Items:   toBigArray([]int64{74}),
			Inspect: makeAddInspectMethod(3),
			Test:    makeTestMethod(17, 0, 1),
		}
		relief *= 17
	} else {
		monkeys[0] = &Monkey{
			Items:   toBigArray([]int64{89, 74}),
			Inspect: makeMultiplyInspectMethod(5),
			Test:    makeTestMethod(17, 4, 7),
		}
		relief *= 17
		monkeys[1] = &Monkey{
			Items:   toBigArray([]int64{75, 69, 87, 57, 84, 90, 66, 50}),
			Inspect: makeAddInspectMethod(3),
			Test:    makeTestMethod(7, 3, 2),
		}
		relief *= 7
		monkeys[2] = &Monkey{
			Items:   toBigArray([]int64{55}),
			Inspect: makeAddInspectMethod(7),
			Test:    makeTestMethod(13, 0, 7),
		}
		relief *= 13
		monkeys[3] = &Monkey{
			Items:   toBigArray([]int64{69, 82, 69, 56, 68}),
			Inspect: makeAddInspectMethod(5),
			Test:    makeTestMethod(2, 0, 2),
		}
		relief *= 2
		monkeys[4] = &Monkey{
			Items:   toBigArray([]int64{72, 97, 50}),
			Inspect: makeAddInspectMethod(2),
			Test:    makeTestMethod(19, 6, 5),
		}
		relief *= 19
		monkeys[5] = &Monkey{
			Items:   toBigArray([]int64{90, 84, 56, 92, 91, 91}),
			Inspect: makeMultiplyInspectMethod(19),
			Test:    makeTestMethod(3, 6, 1),
		}
		relief *= 3
		monkeys[6] = &Monkey{
			Items:   toBigArray([]int64{63, 93, 55, 53}),
			Inspect: makeSquareInspectMethod(),
			Test:    makeTestMethod(5, 3, 1),
		}
		relief *= 5
		monkeys[7] = &Monkey{
			Items:   toBigArray([]int64{50, 61, 52, 58, 86, 68, 97}),
			Inspect: makeAddInspectMethod(4),
			Test:    makeTestMethod(11, 5, 4),
		}
		relief *= 11

	}
	return monkeys, int64(relief)
}
