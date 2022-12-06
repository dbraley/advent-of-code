package main

import (
	"fmt"
	"github.com/dbraley/advent-of-code/file"
	"github.com/golang-collections/collections/stack"
	"os"
	"strconv"
	"strings"
)

var stacks map[string]*stack.Stack

func main() {

	input, err := file.ReadFile("day5/input.txt")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	emptyLine := 0
	for lineNumber, line := range input {
		if line == "" {
			emptyLine = lineNumber
			fmt.Printf("Break at %v\n", lineNumber)
			break
		}
	}

	data := input[0:emptyLine]
	instructions := input[emptyLine+1:]

	stacks = makeStacks(data, emptyLine)

	followInstructions(instructions, stacks)

	fmt.Println()
	printStacks()

}

func followInstructions(instructions []string, stacks map[string]*stack.Stack) {
	for _, line := range instructions {
		if strings.Contains(line, "move") {
			lineArgs := strings.Split(line, " ")
			count, err := strconv.Atoi(lineArgs[1])
			if err != nil {
				fmt.Printf("Couldn't convert count %v to int", lineArgs[1])
				os.Exit(2)
			}
			source := lineArgs[3]
			dest := lineArgs[5]

			moving := stack.New()
			for i := 0; i < count; i++ {
				moving.Push(stacks[source].Pop())
			}
			for i := 0; i < count; i++ {
				stacks[dest].Push(moving.Pop())
			}
			//printStacks()
		}
	}
}

func makeStacks(data []string, emptyLine int) map[string]*stack.Stack {
	stacks := make(map[string]*stack.Stack)

	stackLine := data[len(data)-1]
	fmt.Println(stackLine)
	for col, r := range stackLine {
		if r != ' ' {
			stackName := string(r)
			//fmt.Printf("Stack %v at col %d\n", stackName, col)
			stacks[stackName] = stack.New()
			for i := emptyLine - 2; i >= 0; i-- {
				if len(data[i]) > col {
					v := string(data[i][col])
					if v != " " {
						//fmt.Printf("Pushing to Stack %v value %v\n", stackName, v)
						stacks[stackName].Push(v)
					}
				}
			}
		}
	}
	return stacks
}

func printStacks() (int, error) {
	return fmt.Printf("%v%v%v%v%v%v%v%v%v\n",
		stacks["1"].Peek(),
		stacks["2"].Peek(),
		stacks["3"].Peek(),
		stacks["4"].Peek(),
		stacks["5"].Peek(),
		stacks["6"].Peek(),
		stacks["7"].Peek(),
		stacks["8"].Peek(),
		stacks["9"].Peek(),
	)
}
